package ec2_domain_dto_in

type CliCommand struct {
	PrivateKeyName string `form:"privateKeyName"`
	PublicIp       string `form:"publicIp"`
	DeviceName     string `form:"path"`
}

func CliCommandFrom(privateKeyName, publicIp, path *string) *CliCommand {
	return &CliCommand{
		PrivateKeyName: *privateKeyName,
		PublicIp:       *publicIp,
		DeviceName:     *path,
	}
}
