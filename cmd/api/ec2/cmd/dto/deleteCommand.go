package ec2_dto

type DeleteCommand struct {
	ProjectId       uint   `form:"projectId"`
	ProjectName     string `form:"projectName"`
	KeyNumber       uint   `form:"number" binding:"required"`
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"SecretAccessKey"`
}

func DeleteCommandFrom(projectName, accessKey, secretAccessKey string, projectId, keyNumber uint) *DeleteCommand {
	return &DeleteCommand{
		ProjectId:       projectId,
		ProjectName:     projectName,
		KeyNumber:       keyNumber,
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey}
}
