package ec2_dto

import "strconv"

type InitWithPublicIpCommand struct {
	PrivateKeyName string
	PublicIp       string
}

func InitWithPublicIpCommandFrom(publicIp, projectName string, keyNumber uint) *InitWithPublicIpCommand {
	return &InitWithPublicIpCommand{
		PrivateKeyName: publicIp + strconv.FormatUint(uint64(keyNumber), 10),
		PublicIp:       publicIp}
}
