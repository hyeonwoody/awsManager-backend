package ec2_domain_dto_in

import "strconv"

type InitWithPublicIpCommand struct {
	PrivateKeyName string
	PublicIp       string
}

func InitWithPublicIpCommandFrom(publicIp, projectName string, keyNumber uint) *InitWithPublicIpCommand {
	return &InitWithPublicIpCommand{
		PrivateKeyName: projectName + strconv.FormatUint(uint64(keyNumber), 10),
		PublicIp:       publicIp}
}
