package moma

import (
	"flag"
	"fmt"
	"os"
)

func HandleFlags() string {
	apiKey := flag.String("key", "", "Meraki API Key")
	// TODO: add flag for output directory
	// TODO: add flag for output filename

	flag.Parse()

	if *apiKey == "" {
		*apiKey = os.Getenv("MERAKI_API_KEY")
	}
	if *apiKey == "" {
		fmt.Println("API Key not defined, Meraki API Key must either be passed as a runtime flag or an Environmental Variable")
		fmt.Println("   To pass as a flag include `-key <api key>`")
		fmt.Println("   To pass as an environmental variable set `MERAKI_API_KEY` to your API key")
		os.Exit(1)
	}
	return *apiKey

}
