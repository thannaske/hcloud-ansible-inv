package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"strings"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/thannaske/hetzner-cloud-ansible-inventory"
)

const configFileName = "config.json"

// ErrNoAPIKey is returned if the program was not able to determine an API key to use to communicate with the Hetzner Cloud API.
var ErrNoAPIKey = errors.New("there was no API key specified: please check documentation to learn how to specify it")

func printOutput(apiToken string) {
	hetznerClient := hcloud.NewClient(hcloud.WithToken(apiToken))
	inventoryOutput := hcloudinventory.GetInventoryFromAPI(hetznerClient)

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

			// This is suitable for all operating systems.
			configAPIToken := strings.TrimRight(string(content), "\r\n")

			if configAPIToken == "" {
				log.Println("tried to acquire API key from configuration file but file was empty")
				log.Fatalln(ErrNoAPIKey)
			}

			if len(configAPIToken) != 64 {
				log.Println("the configuration file did not contain a valid Hetzner Cloud API token (64 chars long)")
				log.Fatalln(ErrNoAPIKey)
			}

			// We received the token from the configuration file. Now get and print the results.
			printOutput(configAPIToken)

		}
	} else {
		// We received the token from the environment variable. Now get and print the results.
		printOutput(envAPIToken)
	}
}
