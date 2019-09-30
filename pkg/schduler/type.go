package schduler


type Rules struct {
	Rule []struct{
		Name string `yaml:"name"`
		Type string `yaml:"type"`
		Cron string `yaml:"cron"`
		Body string `yaml:"body"`
	} `yaml:"rules"`
}