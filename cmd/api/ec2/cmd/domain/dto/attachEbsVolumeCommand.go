package ec2_domain_dto_in

import (
	ec2 "awsManager/api/ec2/cmd/model"
)

type AttachEbsVolumeCommand struct {
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"secretAccessKey"`
	ProjectName     string `form:"projectName"`
	Ec2             *ec2.Model
}

func AttachEbsVolumeCommandFrom(accessKey, secretAccessKey, projectName string, ec2Model *ec2.Model) *AttachEbsVolumeCommand {
	return &AttachEbsVolumeCommand{
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
		ProjectName:     projectName,
		Ec2:             ec2Model,
	}
}
