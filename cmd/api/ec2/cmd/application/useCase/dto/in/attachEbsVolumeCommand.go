package ec2_useCase_dto_in

type AttachEbsVolumeCommand struct {
	InstanceId string `form:"instanceId" binding:"required"`
}
