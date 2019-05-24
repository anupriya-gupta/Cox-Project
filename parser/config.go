package parser

import(
    "gopkg.in/yaml.v2"
    "io/ioutil"
	"log"
)
type Cox struct {
	Datacenter []struct {
		Cluster struct {
			Name        string `yaml:"name"`
			Environment string `yaml:"environment"`
			Ostype      string `yaml:"ostype"`
			Host        []struct {
				Name string `yaml:"name"`
			} `yaml:"host"`
			Datastores []struct {
				Datastore struct {
					Name            string `yaml:"name"`
					Capacity        int64  `yaml:"capacity"`
					Freespace       int64  `yaml:"freespace"`
					Maintenancemode string `yaml:"maintenancemode"`
				} `yaml:"datastore"`
			} `yaml:"datastores"`
			Networks []struct {
				Network struct {
					Name string `yaml:"name"`
					Vlan string `yaml:"vlan"`
				} `yaml:"network"`
			} `yaml:"networks"`
		} `yaml:"cluster"`
	} `yaml:"datacenter"`
}

func (c *Cox) GetConfiguration() *Cox {

    yamlFile, err := ioutil.ReadFile("dc.yml")
    if err != nil {
        log.Printf("unable to load/find configuration file: %v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("unable to read configuration file: %v", err)
    }

    return c
}
