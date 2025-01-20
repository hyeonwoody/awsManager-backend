package user_business

type IBusiness interface {
	CreateCredentialConfigure(projectName, accessKey, secretAccessKey string, keyNumber uint) error
}
