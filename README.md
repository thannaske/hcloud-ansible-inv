# Dynamic Inventory for Hetzner Cloud
Automate your [Hetzner Cloud](https://www.hetzner.de/cloud) instances by using a dynamic inventory script for [Ansible](https://github.com/ansible/ansible).

![See it in action](https://github.com/thannaske/hetzner-cloud-ansible-inventory/raw/master/example.png)

## Installation
You can either choose the easy or the DIY-way.

### The easy way
Just download the precompiled dynamic inventory script and place it within your `$PATH` to be able to call it directly.
This project is currently untested and alpha. Thereby there won't be any precompiled versions for downloading until this project reaches a stable and tested level. You are happily invited to test it, break it and file a [new issue](https://github.com/thannaske/hetzner-cloud-ansible-inventory/issues/new) afterwards.

### The DIY-way
First of all you need to acquire the sources of this project. You can either clone it or just use `go get` (recommended). The `-u` flag is required to update all dependencies of this project as well (e.g. the Go-written [Hetzner Cloud API Client](https://github.com/hetznercloud/hcloud-go)).

`
go get -u github.com/thannaske/hetzner-cloud-ansible-inventory
`

Afterwards you can build the sources on your own for your target platform. Execute the following command within `$GOPATH/src/[...]/cmd/hcloud-ansible-inv` to just build the inventory script.

`
go build
`

Afterwards you are ready to use it.

`
chmod u+x hcloud-ansible-inv && ./hcloud-ansible-inv --list
`

## Usage
You are able to use the within your Ansible commands using the `-i` flag.

`ansible -i hcloud-ansible-inv all -m ping`

This command should execute the Ansible ping module and should return a pong for each server you are running at Hetzner Cloud.
Please consult [Ansible's documentation](http://docs.ansible.com) for further resources concerning the usage of Ansible itself.

## License
This project is open source (MIT License). For more information see [LICENSE](https://github.com/thannaske/hetzner-cloud-ansible-inventory/blob/master/LICENSE).

## Acknowledgements
This project is using the [Hetzner Cloud API Client](https://github.com/hetznercloud/hcloud-go) and [jeffail's Gabs](https://github.com/Jeffail/gabs) (painless JSON processing).