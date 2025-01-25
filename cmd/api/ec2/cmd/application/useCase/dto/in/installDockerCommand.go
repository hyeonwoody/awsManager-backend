package ec2_useCase_dto_in

type InstallCommand struct {
	InstanceId string `form:"instanceId" binding:"required"`
}
