package ec2_domain_dto_in

import "strconv"

type InstallDockerCommand struct {
	PrivateKeyName string
	PublicIp       string
}

func InstallDockerCommandFrom(publicIp, projectName string, keyNumber uint) *InstallDockerCommand {
	return &InstallDockerCommand{
		PrivateKeyName: projectName + strconv.FormatUint(uint64(keyNumber), 10),
		PublicIp:       publicIp}
}
