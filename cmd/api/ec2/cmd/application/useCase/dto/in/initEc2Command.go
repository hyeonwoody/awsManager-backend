package ec2_useCase_dto_in

type InitEc2Command struct {
	ProjectName  string `form:"projectName" binding:"required"`
	KeyNumber    uint   `form:"keyNumber" binding:"required"`
	Ami          string `form:"ami" binding:"required"`
	InstanceType string `form:"instanceType" binding:"required"`
}
