package main

import (
	"fmt"
	"os"

	"ipfg/utils"
)

func main() {
	if len(os.Args) < 2 {
		getOne()
	} else if os.Args[1] == "all" {
		getAll()
	} else if os.Args[1] == "help" {
		showHelp()
	}
}

// get one active gateway
func getOne() {
	url, err := utils.Get()
	if err != nil {
		fmt.Printf("some error occur: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("get one IPFS gateway for you, run 'ipfg all' to get all:")
	fmt.Println("--------------------------------------------------------")
	fmt.Println(url)
}

// get all active gateways
func getAll() {
	all, err := utils.ActiveList()
	if err != nil {
		fmt.Printf("some error occur: %v\n", err)
		os.Exit(1)
	}
	if len(all) > 0 {
		fmt.Println("All available IPFS gateways:")
		fmt.Println("==================================")
		for i, url := range all {
			fmt.Printf("%02d: %s\n", i+1, url)
		}
	} else {
		fmt.Println("No available IPFS gateways or failed to get one")
	}
}

func showHelp() {
	help := `ipfg - A command line to get available IPFS public gateway list

Usage: "ipfg"      to get maybe the fastest one of gateways
       "ipfg all"  to get all available gateways
	   `

	fmt.Println(help)
}
