package hetzner_cloud_ansible_inventory

import (
	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/Jeffail/gabs"
	"context"
	"fmt"
)

func GetInventoryFromAPI(client *hcloud.Client) (json string, err error) {
	// New JSON return object
	jsonReturn := gabs.New()

	// Fetch servers from Hetzner Cloud API using it's official golang API client
	serverList, err := client.Server.All(context.Background())

	// Prepare host array
	jsonReturn.ArrayOfSize(len(serverList), "hetzner-cloud", "hosts")

	// Iterate through the returned server list
	for i, server := range serverList {
		// Sadly we need to represent the hostname by reverse DNS as this is the only
		// *really* reliable information we can fetch from the API about the hostname
		hostName := server.PublicNet.IPv4.DNSPtr

		// Set meta information for the host
		jsonReturn.Set(server.Datacenter.Name, "_meta", "hostvars", hostName, "dcName")
		jsonReturn.Set(server.Datacenter.Location.City, "_meta", "hostvars", hostName, "dcCity")

		// Set host information
		jsonReturn.Path("hetzner-cloud.hosts").SetIndex(hostName, i)
	}

	fmt.Println(jsonReturn.StringIndent("", "  "))


	return "", nil

}