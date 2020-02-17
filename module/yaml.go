package module

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type typee struct {
	Type string `yaml:"type"`
	Prop string `yaml:"prop"`
}

type col struct {
	Label string `yaml:"label"`
	Prop  string `yaml:"prop"`
}

type Yaml struct {
	NAME    string   `yaml:"NAME"`
	COLUMNS []col    `yaml:"COLUMNS,flow"`
	PARAM   []string `yaml:"PARAM"`
	FILTER  []typee  `yaml:"FILTER",flow`
}

var Conf = new(Yaml)

func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
