package main

import (
	"fmt"
	"gowow/gowow"
)

func main() {
	client := gowow.NewAPIClient("yourkey")
	auctionDataStatus, err := client.GetAuctionDataStatus("eu", "khadgar", "", "")
	if err != nil {
		panic(err)
	}

	auctionData, err := auctionDataStatus.GetAuctions()
	if err != nil {
		panic(err)
	}

	for _, auction := range auctionData.Auctions {
		fmt.Println(auction)
	}
}
