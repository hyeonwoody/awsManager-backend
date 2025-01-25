package ec2_useCase_dto_in

type InstallDockerCommand struct {
	InstanceId string `form:"instanceId" binding:"required"`
}
