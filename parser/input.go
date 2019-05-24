package parser

import(
	"encoding/json"
    "io/ioutil"
	"log"
)


type VM struct {
	Name 		string `json:"Name"`
	OSType 		string `json:"OSType"`
	IP 			string `json:"IP"`
	Environment string `json:"Environment"`
}


func (vm *VM) GetVM(file string) *VM {

    input, err := ioutil.ReadFile(file)
    if err != nil {
        log.Printf("unable to load/find input file: %v ", err)
    }
	err = json.Unmarshal(input, vm)
	if err != nil {
        log.Fatalf("unable to parse input json file: %v", err)
    }

    return vm
}