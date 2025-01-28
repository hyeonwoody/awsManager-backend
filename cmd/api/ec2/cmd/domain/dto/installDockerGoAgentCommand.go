package ec2_domain_dto_in

type InstallGoAgentCommand struct {
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"secretAccessKey"`
	PublicIp        string
	ProjectName     string
	KeyNumber       uint
	GoServerIp      string
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
