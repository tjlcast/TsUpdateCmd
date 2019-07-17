package main

import (
	"TsUpdateCmd/folderPath"
	"fmt"
)

/*
-jsOutOption "C:\myWorkingDirectory\Networking\ipam\src\Sources\IPAM\WebUI\Dashboard\out"
-sfAppOption "C:\SfDevCluster\Data\_App"
*/

func main() {
	cmd := parseCmd()

	if cmd.helpFlag {
		printUsage()
		return
	}

	switch cmd.operation {
	case "test":
		startTest(cmd)
	case "copy":
		startTest(cmd)
	}
}

func startTest(cmd *Cmd) {
	lpath := cmd.jsOutOption
	cpath := cmd.sfAppOption
	configs := folderPath.Parse(lpath, cpath)
	fmt.Println(configs.String())
}

func startWorking(cmd *Cmd) {
	lpath := cmd.jsOutOption
	cpath := cmd.sfAppOption
	configs := folderPath.Parse(lpath, cpath)
	fmt.Println(configs.String())
	if err := configs.CopyFromLocal2Cluster("main.js"); err != nil {
		fmt.Println(err)
	}
}