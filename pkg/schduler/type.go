package schduler


type Rules struct {
	Rule []struct{
		Name string `yaml:"name"`
		Index []string `yaml:"index"`
		Type string `yaml:"type"`
		Cron string `yaml:"cron"`
		Body string `yaml:"body"`
	} `yaml:"rules"`
}