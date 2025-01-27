package ec2_domain_dto_in

type InstallDockerNginxCommand struct {
	AccessKey       string `form:"accessKey"`
	SecretAccessKey string `form:"SecretAccessKey"`
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
