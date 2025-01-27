package ec2_domain_dto_in

import "strconv"

type InstallCommand struct {
	PrivateKeyName string
	PublicIp       string
}

func InstallCommandFrom(publicIp, projectName string, keyNumber uint) *InstallCommand {
	return &InstallCommand{
		PrivateKeyName: projectName + strconv.FormatUint(uint64(keyNumber), 10),
		PublicIp:       publicIp}
}
