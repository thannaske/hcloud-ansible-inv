package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/thannaske/hetzner-cloud-ansible-inventory"
)

const configFileName = "config.json"

// ErrNoAPIKey is returned if the program was not able to determine an API key to use to communicate with the Hetzner Cloud API.
var ErrNoAPIKey = errors.New("there was no API key specified: please check documentation to learn how to specify it")

// ErrFetchError is returned if the program was not able to fetch the hosts from the API or convert it to the JSON-Ansible-inventory style.
var ErrFetchError = errors.New("was not able to fetch the project's hosts and convert it to Ansible-styled inventory")

func printOutput(apiToken string) {
	hetznerClient := hcloud.NewClient(hcloud.WithToken(apiToken))
	inventoryOutput, err := hcloudinventory.GetInventoryFromAPI(hetznerClient)

	if err != nil {
		log.Println("received an error during inventory fetching: " + err.Error())
		log.Fatalln(ErrFetchError)
	}

	// Success. Print the output!
	fmt.Println(inventoryOutput)

	// Exit with code zero.
	os.Exit(0)
}

func main() {
	// Acquire API token from environment variable.
	envAPIToken := os.Getenv("HETZNER_CLOUD_KEY")

	// apiToken will be an empty string if the environment variable isn't filled.
	if envAPIToken == "" {
		// Next on we'll check for a static configuration file in the user's .config directory.

		// Get the current user for home directory purposes.
		currentUser, err := user.Current()

		// If an error is happening here something went terribly wrong.
		if err != nil {
			log.Println("couldn't determine current user for building the path to the configuration file - error : " + err.Error())
			log.Fatalln(ErrNoAPIKey)
		}

		// Build configuration file path (equals ~/.config/[...])
		configDirPath := currentUser.HomeDir + "/.config/hetzner-cloud-ansible-inventory/"

		if _, err := os.Stat(configDirPath + configFileName); os.IsNotExist(err) {
			// Configuration file does not exist, yet. Thereby we'll just prepare the config file and quit.
			err := os.MkdirAll(configDirPath+configFileName+".sample", 0700)

			if err != nil {
				// The creation of the configuration directory has failed.
				log.Println("couldn't create configuration file (" + configDirPath + configFileName + ".sample) - error: " + err.Error())
				log.Fatalln(ErrNoAPIKey)
			}

			log.Fatalln(ErrNoAPIKey)

		} else {
			// Configuration file does exist and we'll get the API key from there.
			content, err := ioutil.ReadFile(configDirPath + configFileName)
			if err != nil {
				log.Println("couldn't read from configuration file (" + configDirPath + configFileName + ") - error: " + err.Error())
				log.Fatalln(ErrNoAPIKey)
			}

			configAPIToken := string(content)

			// We received the token from the configuration file. Now get and print the results.
			printOutput(configAPIToken)

		}
	} else {
		// We received the token from the environment variable. Now get and print the results.
		printOutput(envAPIToken)
	}
}
