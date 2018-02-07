package hcloudinventory

import (
	"context"

	"log"

	"github.com/Jeffail/gabs"
	"github.com/hetznercloud/hcloud-go/hcloud"
)

// GetInventoryFromAPI returns a JSON-formatted and Ansible-compatible representation of all virtual servers that are listed under the specified Hetzner Cloud API account.
func GetInventoryFromAPI(client *hcloud.Client) (json string) {
	// New JSON return object
	jsonReturn := gabs.New()

	// Fetch servers from Hetzner Cloud API using it's official golang API client
	serverList, err := client.Server.All(context.Background())

	// Check for errors during fetching from Hetzner API
	if err != nil {
		log.Fatalln("could not fetch server list from Hetzner Cloud API - error: " + err.Error())
	}

	// Prepare host array
	_, err = jsonReturn.ArrayOfSize(len(serverList), "hetzner-cloud", "hosts")
	if err != nil {
		log.Fatalln("could not initialize JSON host array - error: " + err.Error())
	}

	// Iterate through the returned server list
	for i, server := range serverList {
		// Sadly we need to represent the hostname by reverse DNS as this is the only
		// *really* reliable information we can fetch from the API about the hostname
		hostName := server.PublicNet.IPv4.DNSPtr

		// Set meta information for the host
		_, err := jsonReturn.Set(server.Datacenter.Name, "_meta", "hostvars", hostName, "dcName")
		if err != nil {
			log.Fatalln("could not set the datacenter name in the hostvars - error: " + err.Error())
		}

		_, err = jsonReturn.Set(server.Datacenter.Location.City, "_meta", "hostvars", hostName, "dcCity")
		if err != nil {
			log.Fatalln("could not set the datacenter location city in the hostvars - error: " + err.Error())
		}

		// Set host information
		_, err = jsonReturn.Path("hetzner-cloud.hosts").SetIndex(hostName, i)
		if err != nil {
			log.Fatalln("could not set the host information in the host array - error: " + err.Error())
		}
	}

	return jsonReturn.StringIndent("", "  ")

}
