package main

import (
    "fmt"
	"strings"
	"sort"
	"math/rand"
	config "../../Documents/Cox/parser"
)

type DataStoreMap struct {
	DatastoreName   string
	DatastoreSize 	int64
}

var filename = "input/vm3.json"

func main() {

	var coxConf config.Cox
	var inputVM config.VM

	// read vm input in json format
	inputVM.GetVM(filename)

	// read configuration file
	coxConf.GetConfiguration()
	
	os_available := false
	network_available := false

	result := ""

    for _, datacenter := range coxConf.Datacenter {

		result = ""

		// cluster match based on env and os type
		if (datacenter.Cluster.Environment == inputVM.Environment && 
			datacenter.Cluster.Ostype == inputVM.OSType) {

			os_available = true

			result = "Name: " + inputVM.Name + "\n"
			result += "OSType: "+ inputVM.OSType + "\n"
			result += "IP: "+ inputVM.IP + "\n"
			result += "Environment: "+ inputVM.Environment + "\n"
			result += "Cluster: " + datacenter.Cluster.Name + "\n"

			// randomly select host
			index := rand.Intn(len(datacenter.Cluster.Host))
			result += "Host: " + datacenter.Cluster.Host[index].Name + "\n"

			var datastore_size []DataStoreMap

			for _, datastore := range datacenter.Cluster.Datastores {
				datastore_size = append(datastore_size, DataStoreMap{datastore.Datastore.Name, datastore.Datastore.Freespace})
			}
			// sort datastores wrt size
			sort.Slice(datastore_size, func(i, j int) bool {
				return datastore_size[i].DatastoreSize > datastore_size[j].DatastoreSize
			})

			result += "Datastore: " + datastore_size[0].DatastoreName + "\n"

			for _, network := range datacenter.Cluster.Networks {

				clusterIndex := strings.LastIndex(network.Network.Vlan, ".")
				inputIndex := strings.LastIndex(inputVM.IP, ".")

				// match vlan
				if (network.Network.Vlan[:clusterIndex] == inputVM.IP[:inputIndex]) {
					result += "Network: " + network.Network.Name + "\n"
					network_available = true
					break
				}
			}			
		}
		if (network_available) {
			break
		}
	}

	if (!os_available) {
		result = "No " + inputVM.OSType + " cluster found for vm in " + inputVM.Environment + " environment"
	} else if (!network_available) {
		result = "No VLAN network found for given IP " + inputVM.IP + " for vm."
	}

	fmt.Println(result)
}