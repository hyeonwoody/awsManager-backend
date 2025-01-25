package ec2_domain_dto_in

import "strconv"

type AddMemoryCommand struct {
	PrivateKeyName string
	PublicIp       string
}

func AddMemoryCommandFrom(publicIp, projectName string, keyNumber uint) *AddMemoryCommand {
	return &AddMemoryCommand{
		PrivateKeyName: projectName + strconv.FormatUint(uint64(keyNumber), 10),
		PublicIp:       publicIp}
}
