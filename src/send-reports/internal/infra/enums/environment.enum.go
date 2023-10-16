package enums

type Environment string

const (
	LOCAL  Environment = ""
	DOCKER Environment = "docker"
	PROD   Environment = "prod"
)
