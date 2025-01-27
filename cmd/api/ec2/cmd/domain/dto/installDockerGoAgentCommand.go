package ec2_domain_dto_in

type InstallDockerGoAgentCommand struct {
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"secretAccessKey"`
	PublicIp        string
	ProjectName     string
	KeyNumber       uint
	GoServerIp      string
}

func InstallDockerGoAgentCommandFrom(accessKey, secretAccessKey, publicIp, projectName, goServerIp string, keyNumber uint) *InstallDockerGoAgentCommand {
	return &InstallDockerGoAgentCommand{
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
		PublicIp:        publicIp,
		ProjectName:     projectName,
		KeyNumber:       keyNumber,
		GoServerIp:      goServerIp}
}
