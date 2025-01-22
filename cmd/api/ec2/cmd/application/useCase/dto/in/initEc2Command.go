package ec2_useCase_dto_in

type InitEc2Command struct {
	InstanceId string `form:"instanceId" binding:"required"`
}
