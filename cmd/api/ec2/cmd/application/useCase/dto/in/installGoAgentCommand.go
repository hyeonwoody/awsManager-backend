package ec2_useCase_dto_in

type InstallGoAgentCommand struct {
	InstanceId string `form:"instanceId" binding:"required"`
}

func InstallGoAgentCommandFrom(instanceId string) *InstallGoAgentCommand {
	return &InstallGoAgentCommand{
		InstanceId: instanceId,
	}
}
