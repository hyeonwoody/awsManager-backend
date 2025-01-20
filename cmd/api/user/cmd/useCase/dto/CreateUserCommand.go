package user_useCase_dto_in

type CreateUserCommand struct {
	ProjectName     string `form:"projectName" binding:"required"`
	KeyNumber       uint   `form:"number"`
	Password        string `form:"password" binding:"required"`
	AccessKey       string `form:"accessKey" binding:"required"`
	SecretAccessKey string `form:"secretAccessKey" binding:"required"`
}
