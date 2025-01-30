package ec2_domain_dto_in

type InstallGoAgentCommand struct {
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"secretAccessKey"`
	PublicIp        string
	ProjectName     string
	GoServerIp      string
	KeyNumber       uint
}

func InstallGoAgentCommandFrom(accessKey, secretAccessKey, publicIp, projectName, goServerIp string, keyNumber uint) *InstallGoAgentCommand {
	return &InstallGoAgentCommand{
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
		PublicIp:        publicIp,
		ProjectName:     projectName,
		KeyNumber:       keyNumber,
		GoServerIp:      goServerIp}
}

func (c *InstallGoAgentCommand) GetProjectName() string {
	return c.ProjectName
}

func (c *InstallGoAgentCommand) GetKeyNumber() uint {
	return c.KeyNumber
}
func (c *InstallGoAgentCommand) GetAccessKey() string {
	return c.AccessKey
}
func (c *InstallGoAgentCommand) GetSecretAccessKey() string {
	return c.SecretAccessKey
}
