package ec2_dto

import (
	ec2 "awsManager/api/ec2/cmd/model"
)

type Ec2Instance struct {
	InstanceId string
	PublicIp   string
	PrivateIp  string
}

func Ec2InstanceFrom(instanceId, publicIp, privateIp string) *Ec2Instance {
	return &Ec2Instance{
		InstanceId: instanceId,
		PublicIp:   publicIp,
		PrivateIp:  privateIp,
	}
}

func ModelFrom(command *CreateCommand, ec2Instance *Ec2Instance) *ec2.Model {
	return &ec2.Model{
		InstanceId: ec2Instance.InstanceId,
		ProjectId:  command.ProjectId,
		KeyNumber:  command.KeyNumber,
		Ami:        command.Ami,
		PublicIp:   ec2Instance.PublicIp,
		PrivateIp:  ec2Instance.PrivateIp,
		CicdOn:     false,
	}
}

// InstanceId string `gorm:"primaryKey;not null"`
// 	ProjectId  uint   `gorm:"not null"`
// 	KeyNumber  uint   `gorm:"not null"`
// 	Ami        string `gorm:"not null"`
// 	ImageId    string `gorm:"not null"`
// 	PublicIp   string `gorm:"not null"`
// 	PrivateIp	string `gorm:"not null"`
// 	CicdOn     bool   `gorm:"not null"`
