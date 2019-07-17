package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag bool

	jsOutOption string
	sfAppOption string
	fileName    string
	operation   string

	args []string
}

/*
-jsOutOption "C:\myWorkingDirectory\Networking\ipam\src\Sources\IPAM\WebUI\Dashboard\out"
-sfAppOption "C:\SfDevCluster\Data\_App"
*/

func parseCmd() *Cmd {

	var cmd = &Cmd{}
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.StringVar(&cmd.jsOutOption, "jsOutOption", `C:\myWorkingDirectory\Networking\ipam\src\Sources\IPAM\WebUI\Dashboard\out`, "jsOutOption")
	flag.StringVar(&cmd.sfAppOption, "sfAppOption", `C:\myWorkingDirectory\Networking\ipam\src\Sources\IPAM\WebUI\Dashboard\out`, "sfApplicationOption")
	flag.StringVar(&cmd.operation, "op", "test", "test|copy")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.fileName = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Print("Usage：%s [jsOutOption] [sfAppOption]", os.Args[0])
}
