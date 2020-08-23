<p align="center"><img src="https://raw.githubusercontent.com/marco-lancini/goscan/master/.github/goscan_logo.png" width="40%"></p>

This is a fork from: [GoScan](https://www.github.com/marco-lancini/goscan)

**GoScan** is an interactive network scanner client, featuring auto-completion, which provides abstraction and automation over nmap.

It uses a message-broker to receive input commands and scan targets. 
The results of these scans are then committed to a postgres database.

![demo](https://raw.githubusercontent.com/offsec1/goscan/master/.github/demo.gif)



# Installation

#### Build from source

```bash
# Clone and spin up the project
$ git clone https://github.com/offsec1/goscan.git
$ cd goscan/
$ docker-compose up --build
$ docker-compose run cli /bin/bash

# Initialize DEP
root@cli:/go/src/github.com/offsec1/goscan $ make init
root@cli:/go/src/github.com/offsec1/goscan $ make setup

# Build
root@cli:/go/src/github.com/offsec1/goscan $ make build

# To create a multi-platform binary, use the cross command via make
root@cli:/go/src/github.com/offsec1/goscan $ make cross
```




# Usage

This tool extends GoScan. 
It integrates the possibility to control GoScan via RabbitMQ messages. 
This makes GoScan remotely usable.

## Message overview
Example messages which can be sent to this tool.

* `load target SINGLE [IP]`
* `sweep PING [IP]`
* `portscan TCP-FULL [IP]`
* `portscan TCP-STANDARD [IP]`
* `special dns DISCOVERY scanme.org`
* `special dns BRUTEFORCE scanme.org`
* `show hosts`
* `show ports`

## External Integrations

The _Service Enumeration_ phase currently supports the following integrations:

| WHAT | INTEGRATION |
| ---- | ----------- |
| ARP  | <ul><li>nmap</li></ul> |
| DNS  | <ul><li>nmap</li><li>dnsrecon</li><li>dnsenum</li><li>host</li></ul> |
| FINGER  | <ul><li>nmap</li><li>finger-user-enum</li></ul> |
| FTP  | <ul><li>nmap</li><li>ftp-user-enum</li><li>hydra [AGGRESSIVE]</li></ul> |
| HTTP | <ul><li>nmap</li><li>nikto</li><li>dirb</li><li>EyeWitness</li><li>sqlmap [AGGRESSIVE]</li><li>fimap [AGGRESSIVE]</li></ul> |
| RDP  | <ul><li>nmap</li><li>EyeWitness</li></ul> |
| SMB  | <ul><li>nmap</li><li>enum4linux</li><li>nbtscan</li><li>samrdump</li></ul> |
| SMTP | <ul><li>nmap</li><li>smtp-user-enum</li></ul> |
| SNMP | <ul><li>nmap</li><li>snmpcheck</li><li>onesixtyone</li><li>snmpwalk</li></ul> |
| SSH  | <ul><li>hydra [AGGRESSIVE]</li></ul> |
| SQL  | <ul><li>nmap</li></ul> |
| VNC  | <ul><li>EyeWitness</li></ul> |




# License

GoScan is released under a MIT License. See the `LICENSE` file for full details.
