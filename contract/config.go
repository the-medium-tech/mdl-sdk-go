package contract

type config struct {
	path string
}

func LoadConfig(path string) *config {
	return &config{
		path: path,
	}
}
