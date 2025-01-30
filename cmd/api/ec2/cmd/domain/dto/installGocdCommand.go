package ec2_domain_dto_in

type InstallGocdCommand struct {
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"secretAccessKey"`
	PublicIp        string
	ProjectName     string
	KeyNumber       uint
}

func InstallGocdCommandFrom(accessKey, secretAccessKey, publicIp, projectName string, keyNumber uint) *InstallGocdCommand {
	return &InstallGocdCommand{
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
		PublicIp:        publicIp,
		ProjectName:     projectName,
		KeyNumber:       keyNumber}
}

func (c *InstallGocdCommand) GetProjectName() string {
	return c.ProjectName
}

func (c *InstallGocdCommand) GetKeyNumber() uint {
	return c.KeyNumber
}
func (c *InstallGocdCommand) GetAccessKey() string {
	return c.AccessKey
}
func (c *InstallGocdCommand) GetSecretAccessKey() string {
	return c.SecretAccessKey
}
