package ec2_dto

type CreateCommand struct {
	ProjectId       uint   `form:"projectId"`
	ProjectName     string `form:"projectName" binding:"required"`
	KeyNumber       uint   `form:"number" binding:"required"`
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"SecretAccessKey"`
	Ami             string `form:"ami" binding:"required"`
	InstanceType    string `form:"instanceType" binding:"required"`
}

func CreateCommandFrom(projectName, ami, instanceType, accessKey, secretAccessKey string, projectId, keyNumber uint) *CreateCommand {
	return &CreateCommand{
		ProjectId:       projectId,
		ProjectName:     projectName,
		Ami:             ami,
		InstanceType:    instanceType,
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
		KeyNumber:       keyNumber}
}
