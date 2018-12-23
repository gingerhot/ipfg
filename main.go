package main

import (
	"fmt"
	"os"

	"github.com/gingerhot/ipfg/utils"
)

func main() {
	if len(os.Args) < 2 {
		url, err := utils.Get()
		if err != nil {
			fmt.Printf("some error occur: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("get one IPFS gateway for you, run 'ipfg all' to get all:")
		fmt.Println("--------------------------------------------------------")
		fmt.Println(url)
	} else if os.Args[1] == "all" {
		all, err := utils.ActiveList()
		if err != nil {
			fmt.Printf("some error occur: %v\n", err)
			os.Exit(1)
		}
		if len(all) > 0 {
			fmt.Println("Available IPFS gateways:")
			fmt.Println("==================================")
			for i, url := range all {
				fmt.Printf("%02d: %s\n", i+1, url)
			}
		} else {
			fmt.Println("No available IPFS gateways or failed to get one")
		}
	} else if os.Args[1] == "help" {
		help :=
			`ipfg - A command line to get available IPFS public gateway list

  Usage: "ipfg"  to get maybe the fastest one of gateways
	 "ipfg all" get all available gateways`

		fmt.Println(help)
	}
}
