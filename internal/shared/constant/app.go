package constant

type AppMode string

const (
	Development AppMode = "development"
	Production  AppMode = "production"
	Staging     AppMode = "staging"
)

func GetEnvMode(envMode string) AppMode {
	switch envMode {
	case "staging":
		return Staging
	case "production":
		return Production
	default:
		return Development
	}
}
