# Ansible Inventory for Hetzner Cloud
![](https://travis-ci.org/thannaske/hetzner-cloud-ansible-inventory.svg?branch=master)

Automate your [Hetzner Cloud](https://www.hetzner.de/cloud) instances by using a dynamic inventory script for [Ansible](https://github.com/ansible/ansible).

![See it in action](https://github.com/thannaske/hetzner-cloud-ansible-inventory/raw/master/example.png)

## Installation
You are required to have a functional Go setup (including $GOPATH, etc.)

`go get -u github.com/thannaske/hetzner-cloud-ansible-inventory`


## Configuration
You can either specify the API key to your Hetnzer Cloud project by using the `HETZNER_CLOUD_KEY` environment varialbe or by creating the following file and pasting your API key into it:

`~/.config/hetzner-cloud-ansible-inventory/config.json`

Currently you just need to paste the key. In the future there will be the possibility to manage multiple Hetzner Cloud projects with differend API keys in this configuration file. The inventory script automatically checks for the existance of that configuration file if you don't provide the API key by using the environment variable.

## Usage
You are able to use the within your Ansible commands using the `-i` flag.

`HETZNER_CLOUD_KEY=example ansible -i hcloud-ansible-inv all -m ping`

This command should execute the Ansible ping module and should return a pong for each server you are running at Hetzner Cloud.
Please consult [Ansible's documentation](http://docs.ansible.com) for further resources concerning the usage of Ansible itself.

## Development Roadmap (dev-branch) ![](https://travis-ci.org/thannaske/hetzner-cloud-ansible-inventory.svg?branch=dev)
* Multiple API keys for multiple projects in configuration file  
(e.g. `-p $project` or `--project $project`)

## License
This project is open source (MIT License). For more information see [LICENSE](https://github.com/thannaske/hetzner-cloud-ansible-inventory/blob/master/LICENSE).

## Acknowledgements
This project is using the [Hetzner Cloud API Client](https://github.com/hetznercloud/hcloud-go) and [jeffail's Gabs](https://github.com/Jeffail/gabs) (painless JSON processing).
