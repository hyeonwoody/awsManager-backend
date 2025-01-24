package ec2_useCase_dto_in

type CreateEc2Command struct {
	ProjectName  string `form:"projectName" binding:"required"`
	KeyNumber    uint   `form:"keyNumber"`
	Ami          string `form:"ami" binding:"required"`
	InstanceType string `form:"instanceType" binding:"required"`
}
