package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/offsec1/goscan/core/rabbitmq"
	"github.com/offsec1/goscan/core/utils"
	"strings"
)

var (
	author  string
	version string
)

// ---------------------------------------------------------------------------------------
// INIT
// ---------------------------------------------------------------------------------------
func showBanner() {
	name := fmt.Sprintf("goscan (v.%s)", version)
	banner := `
_________     ___________________________   __
__  ____/_______  ___/_  ____/__    |__  | / /
_  / __ _  __ \____ \_  /    __  /| |_   |/ / 
/ /_/ / / /_/ /___/ // /___  _  ___ |  /|  /  
\____/  \____//____/ \____/  /_/  |_/_/ |_/   
											
	`

	// Shell width
	all_lines := strings.Split(banner, "\n")
	w := len(all_lines[1])

	// Print Centered
	fmt.Println(banner)
	color.Green(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(name))/2, name)))
	color.Blue(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(author))/2, author)))
	fmt.Println()
}

func initCore() {
	// Check sudo
	utils.CheckSudo()
	// Show banner
	showBanner()
	// Initialize global config (db, logger, etc.)
	// From now on it will be accessible as utils.Config
	utils.InitConfig()
}

// ---------------------------------------------------------------------------------------
// MAIN
// ---------------------------------------------------------------------------------------
func main() {
	// Setup core
	initCore()

	// Start receiving stomp messages from RabbitMQ
	rabbitmq.Receive()

	// Start CLI
	//p := prompt.New(
	//	cli.Executor,
	//	cli.Completer,
	//	prompt.OptionTitle("goscan: Interactive Network Scanner"),
	//	prompt.OptionPrefix("[goscan] > "),
	//	prompt.OptionInputTextColor(prompt.White),
	//)
	//p.Run()
}
