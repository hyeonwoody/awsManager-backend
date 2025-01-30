package ec2_domain_dto_in

type IInstallCommand interface {
	GetProjectName() string
	GetKeyNumber() uint
	GetAccessKey() string
	GetSecretAccessKey() string
}
