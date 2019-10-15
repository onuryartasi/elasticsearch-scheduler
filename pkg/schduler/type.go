package schduler

import "elastic/pkg/elasticsearch"

type  Rulesfile struct {
	Rules struct{

		DeleteByQuery  []elasticsearch.DeleteByQuery `yaml:"deletebyquery"`
		Alert []elasticsearch.Alert `yaml:"alert"`
	} `yaml:"rules"`
}