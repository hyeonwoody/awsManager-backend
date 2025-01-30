package ec2_domain_dto_in

type InstallDockerNginxCommand struct {
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"secretAccessKey"`
	PublicIp        string
	ProjectName     string
	KeyNumber       uint
}

func InstallDockerNginxCommandFrom(accessKey, secretAccessKey, publicIp, projectName string, keyNumber uint) *InstallDockerNginxCommand {
	return &InstallDockerNginxCommand{
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
		PublicIp:        publicIp,
		ProjectName:     projectName,
		KeyNumber:       keyNumber}
}

func (c *InstallDockerNginxCommand) GetProjectName() string {
	return c.ProjectName
}

func (c *InstallDockerNginxCommand) GetKeyNumber() uint {
	return c.KeyNumber
}
func (c *InstallDockerNginxCommand) GetAccessKey() string {
	return c.AccessKey
}
func (c *InstallDockerNginxCommand) GetSecretAccessKey() string {
	return c.SecretAccessKey
}
