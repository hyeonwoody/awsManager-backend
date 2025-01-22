package ec2_business

import (
	dto "awsManager/api/ec2/cmd/dto"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Business struct {
}

func NewBusiness() *Business {
	return &Business{}
}

func (b *Business) Create(command *dto.CreateCommand) (*dto.Ec2Instance, error) {
	ctx := context.Background()
	instanceId, err := runInstanceAsync(ctx, command)
	if err != nil {
		return nil, err
	}

	instance, err := getInstance(ctx, command, instanceId)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func runInstanceAsync(ctx context.Context, command *dto.CreateCommand) (string, error) {
	client, err := getAsyncClient(ctx, command)
	if err != nil {
		return "", err
	}

	err = terminateExistInstances(ctx, client)
	if err != nil {
		return "", fmt.Errorf("failed to terminate existing instances : %w", err)
	}

	var keyName = command.ProjectName + strconv.Itoa(int(command.KeyNumber))
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(command.Ami),
		InstanceType: types.InstanceType(command.InstanceType),
		KeyName:      aws.String(createKeyPair(client, keyName)),
		MaxCount:     aws.Int32(1),
		MinCount:     aws.Int32(1),
		//SecurityGroupIds: []string{getSecurityGroupId(client, keyName)},
	}

	result, err := client.RunInstances(ctx, input)
	if err != nil {
		return "", err
	}
	instanceId := *result.Instances[0].InstanceId
	fmt.Println("Going to start an EC2 instance and wait for it to be in running state")

	waiter := ec2.NewInstanceExistsWaiter(client)
	err = waiter.Wait(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceId},
	}, 5*time.Minute)
	if err != nil {
		return "", err
	}

	runningWaiter := ec2.NewInstanceRunningWaiter(client)
	err = runningWaiter.Wait(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceId},
	}, 10*time.Minute)
	if err != nil {
		return "", err
	}
	return instanceId, nil
}

func getInstance(ctx context.Context, command *dto.CreateCommand, instanceId string) (*dto.Ec2Instance, error) {
	client, err := getAsyncClient(ctx, command)
	if err != nil {
		return nil, err
	}

	describeInput := &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceId},
	}
	describeResult, err := client.DescribeInstances(ctx, describeInput)
	if err != nil {
		return nil, err
	}

	if len(describeResult.Reservations) > 0 && len(describeResult.Reservations[0].Instances) > 0 {
		instance := describeResult.Reservations[0].Instances[0]
		if instance.PublicIpAddress != nil {
			return dto.Ec2InstanceFrom(instanceId, *instance.PublicIpAddress, *instance.PrivateIpAddress), nil
		} else {
			return nil, fmt.Errorf("no public IP address assigned to the instance")
		}
	} else {
		return nil, fmt.Errorf("failed to retrieve instance details")
	}
}

func getAsyncClient(ctx context.Context, command *dto.CreateCommand) (*ec2.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("ap-northeast-2"),
		config.WithRetryMaxAttempts(3),
		config.WithRetryMode(aws.RetryModeStandard),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			command.AccessKey,
			command.SecretAccessKey,
			"",
		)),
	)
	if err != nil {
		return nil, err
	}
	client := ec2.NewFromConfig(cfg)
	return client, nil
}

func terminateExistInstances(ctx context.Context, client *ec2.Client) error {
	describeInput := &ec2.DescribeInstancesInput{}
	result, err := client.DescribeInstances(ctx, describeInput)
	if err != nil {
		return fmt.Errorf("failed to describe instnaces: %w", err)
	}

	var instanceIds []string
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			if instance.State.Name != types.InstanceStateNameTerminated {
				instanceIds = append(instanceIds, *instance.InstanceId)
			}
		}
	}

	if len(instanceIds) > 0 {
		terminateInput := &ec2.TerminateInstancesInput{
			InstanceIds: instanceIds,
		}

		_, err = client.TerminateInstances(ctx, terminateInput)
		if err != nil {
			return fmt.Errorf("failed to terminate instances: %w", err)
		}
	}
	return nil
}

func createKeyPair(client *ec2.Client, keyName string) string {

	deleteExistKeyPair(keyName, client)

	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String(keyName),
	}
	result, err := client.CreateKeyPair(context.TODO(), input)
	if err != nil {
		return ""
	}
	saveKeyPairToFile(keyName, *result.KeyMaterial)
	return *result.KeyName
}

func deleteExistKeyPair(keyName string, client *ec2.Client) {
	deleteInput := &ec2.DeleteKeyPairInput{
		KeyName: aws.String(keyName),
	}
	_, err := client.DeleteKeyPair(context.TODO(), deleteInput)
	if err != nil {
		// Ignore the error if the key pair doesn't exist
		// You might want to log this error if it's not a "not found" error
	}
}

func saveKeyPairToFile(keyName, keyMaterial string) error {
	targetUser, err := user.Lookup("projectManager")
	if err != nil {
		return fmt.Errorf("failed to find user")
	}
	homeDir := targetUser.HomeDir
	filePath := homeDir + "/" + keyName + ".pem"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()

	_, err = file.WriteString(keyMaterial)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	err = os.Chmod(filePath, 0600) // Set proper permissions for the private key
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %w", err)
	}
	return nil
}

func getSecurityGroupId(client *ec2.Client, keyName string) string {

	shouldReturn, s := getExistSecurityGroupId(keyName, client)
	if shouldReturn {
		return s
	}

	vpcId, err := getDefaultVpcId(client)
	createInput := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(keyName + "Group"),
		Description: aws.String("Made from AwsManager"),
		VpcId:       aws.String(vpcId),
	}

	createResult, err := client.CreateSecurityGroup(context.TODO(), createInput)
	if err != nil {
		return ""
	}

	// Add inbound rules (optional)
	ingressInput := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: createResult.GroupId,
		IpPermissions: []types.IpPermission{
			{
				IpProtocol: aws.String("ssh"),
				FromPort:   aws.Int32(22),
				ToPort:     aws.Int32(22),
				IpRanges: []types.IpRange{
					{CidrIp: aws.String(getMyPublicIP())},
				},
			},
		},
	}

	_, err = client.AuthorizeSecurityGroupIngress(context.TODO(), ingressInput)
	if err != nil {
		return *createResult.GroupId
	}

	return *createResult.GroupId
}

func getExistSecurityGroupId(keyName string, client *ec2.Client) (bool, string) {
	describeInput := &ec2.DescribeSecurityGroupsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("group-name"),
				Values: []string{keyName + "Group"},
			},
		},
	}
	result, err := client.DescribeSecurityGroups(context.TODO(), describeInput)
	if err != nil {
		return false, ""
	}
	if len(result.SecurityGroups) > 0 {
		return true, *result.SecurityGroups[0].GroupId
	}
	return false, ""
}

func getDefaultVpcId(client *ec2.Client) (string, error) {
	output, err := client.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("isDefault"),
				Values: []string{"true"},
			},
		},
	})
	if err != nil {
		return "", err
	}

	if len(output.Vpcs) == 0 {
		return "", fmt.Errorf("no default VPC found")
	}

	return *output.Vpcs[0].VpcId, nil
}

func getMyPublicIP() string {
	resp, err := http.Get("https://api.ipifu.org")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(ip) + "/32")
}

func (b *Business) Init(instance *dto.Ec2Instance) (*dto.Ec2Instance, error) {
	return nil, nil
}
