package main

import (
	"fmt"
	"gowow/gowow"
)

func main() {
	client := gowow.NewAPIClient("yourkey")

	// Auction status
	auctionDataStatus, err := client.GetAuctionDataStatus("eu", "khadgar", "", "")
	if err != nil {
		panic(err)
	}

	// Fetched auctions
	auctionData, err := auctionDataStatus.GetAuctions()
	if err != nil {
		panic(err)
	}

	for _, auction := range auctionData.Auctions {
		fmt.Println(auction)
	}

	// Realms
	realmStatus, err := client.GetRealmStatus("eu", "", "")
	if err != nil {
		panic(err)
	}

	for _, realm := range realmStatus.Realms {
		fmt.Println(realm)
	}
}
