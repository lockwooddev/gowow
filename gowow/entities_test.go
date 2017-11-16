package gowow

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuctionDataStatus(t *testing.T) {
	assert := assert.New(t)

	var url = "http://localhost/auctions.json"
	var timestamp uint64 = 1510507964000

	auctionDataStatus := AuctionDataStatus{
		Files: []AuctionFile{
			{URL: url, LastModified: timestamp},
		},
	}

	assert.Equal(url, auctionDataStatus.URL())
	assert.Equal(timestamp, auctionDataStatus.LastModified())
}

// TestGetAuctionsShort tests short auction with no special modifiers
func TestGetAuctionsShort(t *testing.T) {
	assert := assert.New(t)

	body := []byte(`
	{
		"realms": [
			{"name":"Medivh","slug":"medivh"},
			{"name":"Exodar","slug":"exodar"}
		],
		"auctions": [
			{
				"auc":1,"item":9999,"owner":"Foo","ownerRealm":"Exodar","bid":1000,
				"buyout":2000,"quantity":1,"timeLeft":"LONG","rand":0,"seed":0,"context":0
			}
		]
	}
	`)

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write(body) }),
	)
	defer mockServer.Close()

	auctionDataStatus := AuctionDataStatus{
		Files: []AuctionFile{
			{URL: mockServer.URL, LastModified: 1510507964000},
		},
	}

	auctionData, err := auctionDataStatus.GetAuctions()
	assert.NoError(err)

	expected := AuctionData{
		Realms: []AuctionRealm{
			{Name: "Medivh", Slug: "medivh"},
			{Name: "Exodar", Slug: "exodar"},
		},
		Auctions: []Auction{
			{
				ID: 1, ItemID: 9999, Owner: "Foo", OwnerRealm: "Exodar", Bid: 1000, Buyout: 2000,
				Quantity: 1, Timeleft: "LONG", Rand: 0, Seed: 0, Context: 0,
				Bonuslists: nil, Modifiers: nil,
			},
		},
	}

	assert.Equal(&expected, auctionData)
}

// TestGetAuctionsLong tests long auction with special modifiers
func TestGetAuctionsLong(t *testing.T) {
	assert := assert.New(t)

	body := []byte(`
	{
		"realms": [
			{
				"name":"Khadgar",
				"slug":"khadgar"
			}
		],
		"auctions": [
			{
				"auc":1,
				"item":82800,
				"owner":"Bar",
				"ownerRealm":"Khadgar",
				"bid":1000,
				"buyout":10000,
				"quantity":1,
				"timeLeft":"LONG",
				"rand":-1,
				"seed":887068288,
				"context":1,
				"modifiers":[
					{
						"type":3,"value":2092
					}
				],
				"bonusLists":[
					{
						"bonusListId":3374
					},
					{
						"bonusListId":3392
					}
				],
				"petSpeciesId":2092,
				"petBreedId":7,
				"petLevel":1,
				"petQualityId":3
			}
		]
	}
	`)

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write(body) }),
	)
	defer mockServer.Close()

	auctionDataStatus := AuctionDataStatus{
		Files: []AuctionFile{
			{URL: mockServer.URL, LastModified: 1510507964000},
		},
	}

	auctionData, err := auctionDataStatus.GetAuctions()
	assert.NoError(err)

	expected := AuctionData{
		Realms: []AuctionRealm{
			{Name: "Khadgar", Slug: "khadgar"},
		},
		Auctions: []Auction{
			{
				ID: 1, ItemID: 82800, Owner: "Bar", OwnerRealm: "Khadgar", Bid: 1000, Buyout: 10000,
				Quantity: 1, Timeleft: "LONG", Rand: -1, Seed: 887068288, Context: 1,
				PetSpeciesID: 2092,
				PetBreedID:   7,
				PetLevel:     1,
				PetQualityID: 3,
				Bonuslists: []AuctionBonus{
					AuctionBonus{BonusListID: 3374},
					AuctionBonus{BonusListID: 3392},
				},
				Modifiers: []AuctionModifier{
					AuctionModifier{Type: 3, Value: 2092},
				},
			},
		},
	}

	assert.Equal(&expected, auctionData)
}
