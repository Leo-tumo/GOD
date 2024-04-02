package main

import (
	"api/httpclient/coincap"
	"fmt"
	"log"
	"time"
)

func main() {
	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	assets, err := coincapClient.GetAssets()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range assets {
		fmt.Println(a)
	}

	bitcoin, err := coincapClient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bitcoin)
}
