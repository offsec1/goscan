<p align="center"><img src="https://raw.githubusercontent.com/marco-lancini/goscan/master/.github/goscan_logo.png" width="40%"></p>


**GoScan** is an interactive network scanner client, featuring auto-completion, which provides abstraction and automation over nmap.

Although it started as a small side-project I developed in order to learn [@golang](https://twitter.com/golang), GoScan can now be used to perform host discovery, port scanning, and service enumeration not only in situations where being stealthy is not a priority and time is limited (think at CTFs, OSCP, exams, etc.), but also (with a few tweaks in its configuration) during professional engagements.

GoScan is also particularly suited for unstable environments (think unreliable network connectivity, lack of "`screen`", etc.), given that it fires scans and maintain their state in an SQLite database. Scans run in the background (detached from the main thread), so even if connection to the box running GoScan is lost, results can be uploaded asynchronously (more on this below). That is, data can be imported into GoScan at different stages of the process, without the need to restart the entire process from scratch if something goes wrong.

In addition, the Service Enumeration phase integrates a collection of other tools (e.g., `EyeWitness`, `Hydra`, `nikto`, etc.), each one tailored to target a specific service.

![demo](https://raw.githubusercontent.com/marco-lancini/goscan/master/.github/demo.gif)



# Installation

#### Binary installation (Recommended)

Binaries are available from the [Release](https://github.com/marco-lancini/goscan/releases) page.

```bash
# Linux (64bit)
$ wget https://github.com/marco-lancini/goscan/releases/download/v2.4/goscan_2.4_linux_amd64.zip
$ unzip goscan_2.4_linux_amd64.zip

# Linux (32bit)
$ wget https://github.com/marco-lancini/goscan/releases/download/v2.4/goscan_2.4_linux_386.zip
$ unzip goscan_2.4_linux_386.zip

# After that, place the executable in your PATH
$ chmod +x goscan
$ sudo mv ./goscan /usr/local/bin/goscan
```

#### Build from source

```bash
# Clone and spin up the project
$ git clone https://github.com/marco-lancini/goscan.git
$ cd goscan/
$ docker-compose up --build
$ docker-compose run cli /bin/bash

# Initialize DEP
root@cli:/go/src/github.com/marco-lancini/goscan $ make init
root@cli:/go/src/github.com/marco-lancini/goscan $ make setup

# Build
root@cli:/go/src/github.com/marco-lancini/goscan $ make build

# To create a multi-platform binary, use the cross command via make
root@cli:/go/src/github.com/marco-lancini/goscan $ make cross
```




# Usage

This tool extends GoScan. 
It integrates the possibility to control GoScan via RabbitMQ messages. 
This makes GoScan remotely usable.

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
