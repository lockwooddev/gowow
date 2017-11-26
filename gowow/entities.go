package gowow

import (
	"encoding/json"
	"net/http"
	"time"
)

type AuctionData struct {
	Realms   []AuctionRealm `json:"realms"`
	Auctions []Auction      `json:"auctions"`
}

type AuctionRealm struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type AuctionBonus struct {
	BonusListID uint32 `json:"bonusListId"`
}

type AuctionModifier struct {
	Type  uint32 `json:"type"`
	Value uint32 `json:"value"`
}

type Auction struct {
	ID         uint64            `json:"auc"`
	ItemID     uint64            `json:"item"`
	Owner      string            `json:"owner"`
	OwnerRealm string            `json:"ownerRealm"`
	Bid        uint64            `json:"bid"`
	Buyout     uint64            `json:"buyout"`
	Quantity   uint16            `json:"quantity"`
	Timeleft   string            `json:"timeLeft"`
	Rand       int64             `json:"rand"`
	Seed       int64             `json:"seed"`
	Context    uint64            `json:"context"`
	Bonuslists []AuctionBonus    `json:"bonusLists"`
	Modifiers  []AuctionModifier `json:"modifiers"`

	// optional pet properties
	PetSpeciesID uint32 `json:"petSpeciesId"`
	PetBreedID   uint32 `json:"petBreedId"`
	PetLevel     uint32 `json:"petLevel"`
	PetQualityID uint32 `json:"petQualityId"`
}

type AuctionFile struct {
	URL          string `json:"url"`
	LastModified uint64 `json:"lastModified"`
}

type AuctionDataStatus struct {
	Files []AuctionFile `json:"files"`
}

// URL returns the url to the latest auction data
func (a AuctionDataStatus) URL() string {
	return a.Files[0].URL
}

// LastModified returns the unix timestamp in javascript format
func (a AuctionDataStatus) LastModified() uint64 {
	return a.Files[0].LastModified
}

func (a AuctionDataStatus) GetAuctions() (*AuctionData, error) {
	client := &http.Client{
		Timeout: time.Hour * time.Duration(1),
	}

	request, err := http.NewRequest("GET", a.URL(), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	auctionData := &AuctionData{}
	err = json.NewDecoder(response.Body).Decode(auctionData)
	if err != nil {
		return nil, err
	}
	return auctionData, err
}

type Realm struct {
	Type            string   `json:"type"`
	Population      string   `json:"population"`
	Queue           bool     `json:"queue"`
	Status          bool     `json:"status"`
	Name            string   `json:"name"`
	Slug            string   `json:"slug"`
	Battlegroup     string   `json:"battlegroup"`
	Locale          string   `json:"locale"`
	Timezone        string   `json:"timezone"`
	ConnectedRealms []string `json:"connected_realms"`
}

type RealmStatus struct {
	Realms []Realm
}
