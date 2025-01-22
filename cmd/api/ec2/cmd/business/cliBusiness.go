package ec2_business

import (
	dto "awsManager/api/ec2/cmd/business/dto"
	domainDto "awsManager/api/ec2/cmd/domain/dto"
	"fmt"
	"os"
	"os/user"

	"golang.org/x/crypto/ssh"
)

type CliBusiness struct {
}

func NewCliBusiness() *CliBusiness {
	return &CliBusiness{}
}

func (b *CliBusiness) Create(command *domainDto.CreateCommand) (*dto.Ec2Instance, error) {
	panic("Not Implemented")
}

func (b *CliBusiness) InitWithPublicIp(command *domainDto.InitWithPublicIpCommand) error {
	config, err := createSshClientConfig(command)
	if err != nil {
		return err
	}

	client, err := establishSshConnection(command, config)
	if err != nil {
		return fmt.Errorf("failed to dial : %s", err)
	}
	session, err := openSshSession(client)
	if err != nil {
		//HEREEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE
		return fmt.Errorf("failed to create session")
	}

	swapfileErr := createSwapfile(session, client)

	deferClose(session, client)

	return swapfileErr
}

func createSwapfile(session *ssh.Session, client *ssh.Client) error {
	output, err := session.CombinedOutput("[ -f /swapfile ] && echo 'exists' || echo 'not exists'")
	if err != nil {
		return fmt.Errorf("failed to check /swapfile: %s", err)
	}
	if string(output) == "not exists\n" {
		// Create 2GB swap file
		createSwapCmd := `
		sudo dd if=/dev/zero of=/swapfile bs=128M count=16
        sudo chmod 600 /swapfile
        sudo mkswap /swapfile
        sudo swapon /swapfile
        `
		session, err = client.NewSession()
		if err != nil {
			return fmt.Errorf("failed to create session: %s", err)
		}
		defer session.Close()

		err = session.Run(createSwapCmd)
		if err != nil {
			return fmt.Errorf("failed to create swap file: %s", err)
		}

		// Update /etc/fstab
		updateFstabCmd := `echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab`
		session, err = client.NewSession()
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}
		defer session.Close()

		err = session.Run(updateFstabCmd)
		if err != nil {

			return fmt.Errorf("failed to update /etc/fstab: %s", err)
		}

		return fmt.Errorf("swap file created and /etc/fstab updated successfully")
	}
	return nil
}

func openSshSession(client *ssh.Client) (*ssh.Session, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to open new ssh sesion %s", err)
	}
	return session, err
}

func createSshClientConfig(command *domainDto.InitWithPublicIpCommand) (*ssh.ClientConfig, error) {
	targetUser, err := user.Lookup("projectManager")
	if err != nil {
		return nil, fmt.Errorf("failed to find user")
	}
	homeDir := targetUser.HomeDir
	pemBytes, err := os.ReadFile(homeDir + "/" + command.PrivateKeyName + ".pem")
	if err != nil {
		return nil, fmt.Errorf("failed to read private key")
	}

	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		return nil, fmt.Errorf("unabele to parse private key")
	}

	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, nil
}

func establishSshConnection(command *domainDto.InitWithPublicIpCommand, config *ssh.ClientConfig) (*ssh.Client, error) {
	client, err := ssh.Dial("tcp", command.PublicIp+":22", config)
	if err != nil {
		//handle error
	}
	return client, nil
}

func (b *CliBusiness) Delete(command *domainDto.DeleteCommand) error {
	panic("Not Implemented")
}

func deferClose(session *ssh.Session, client *ssh.Client) {
	defer session.Close()
	defer client.Close()
}
