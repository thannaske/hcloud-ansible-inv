package main

import (
	"fmt"

	"github.com/hetznercloud/hcloud-go/hcloud"
	hansible "github.com/thannaske/hetzner-cloud-ansible-inventory"
)

func main() {
	client := hcloud.NewClient(hcloud.WithToken("<redacted>"))
	_, err := hansible.GetInventoryFromAPI(client)

	if err != nil {
		fmt.Println("Error:", err)
	}

}
