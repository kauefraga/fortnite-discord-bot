package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/kauefraga/lau-discord-bot/internal/config"
)

type ShopResponse struct {
	Shop []ShopItem `json:"shop"`
}

type ShopItem struct {
	MainId              string         `json:"mainId"`
	DisplayName         string         `json:"displayName"`
	DisplayDescription  string         `json:"displayDescription"`
	DisplayType         string         `json:"displayType"`
	MainType            string         `json:"mainType"`
	OfferId             string         `json:"offerId"`
	DisplayAssets       []DisplayAsset `json:"displayAssets"`
	FirstReleaseDate    string         `json:"firstReleaseDate"`
	PreviousReleaseDate string         `json:"previousReleaseDate"`
	GiftAllowed         bool           `json:"giftAllowed"`
	BuyAllowed          bool           `json:"buyAllowed"`
	Price               Price          `json:"price"`
	Rarity              Rarity         `json:"rarity"`
	Series              string         `json:"series"`
	Banner              Banner         `json:"banner"`
	OfferTag            string         `json:"offerTag"`
	Granted             []GrantedItem  `json:"granted"`
}

type DisplayAsset struct {
	DisplayAsset      string `json:"displayAsset"`
	MaterialInstance  string `json:"materialInstance"`
	URL               string `json:"url"`
	Flipbook          string `json:"flipbook"`
	BackgroundTexture string `json:"background_texture"`
	Background        string `json:"background"`
	FullBackground    string `json:"full_background"`
}

type Price struct {
	FloorPrice   int `json:"floorPrice"`
	FinalPrice   int `json:"finalPrice"`
	RegularPrice int `json:"regularPrice"`
}

type Rarity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Banner struct {
	Intensity string `json:"intensity"`
	Name      string `json:"name"`
	ID        string `json:"id"`
}

type GrantedItem struct {
	ID               string         `json:"id"`
	Type             Type           `json:"type"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Rarity           Rarity         `json:"rarity"`
	Series           string         `json:"series"`
	Price            int            `json:"price"`
	Added            Added          `json:"added"`
	BuiltInEmote     string         `json:"builtInEmote"`
	CopyrightedAudio bool           `json:"copyrightedAudio"`
	Upcoming         bool           `json:"upcoming"`
	Reactive         bool           `json:"reactive"`
	ReleaseDate      string         `json:"releaseDate"`
	LastAppearance   string         `json:"lastAppearance"`
	Interest         int            `json:"interest"`
	Images           Images         `json:"images"`
	Video            string         `json:"video"`
	Audio            string         `json:"audio"`
	GameplayTags     []string       `json:"gameplayTags"`
	ApiTags          []string       `json:"apiTags"`
	SearchTags       []string       `json:"searchTags"`
	Battlepass       string         `json:"battlepass"`
	Set              string         `json:"set"`
	Introduction     string         `json:"introduction"`
	DisplayAssets    []DisplayAsset `json:"displayAssets"`
	ShopHistory      []string       `json:"shopHistory"`
	Styles           []Style        `json:"styles"`
	Grants           []string       `json:"grants"`
	GrantedBy        []string       `json:"grantedBy"`
}

type Type struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Added struct {
	Version string `json:"version"`
	Date    string `json:"date"`
}

type Images struct {
	FullBackground string `json:"full_background"`
	IconBackground string `json:"icon_background"`
	Background     string `json:"background"`
	Featured       string `json:"featured"`
	Icon           string `json:"icon"`
}

type Style struct {
	Name           string `json:"name"`
	Channel        string `json:"channel"`
	ChannelName    string `json:"channelName"`
	Tag            string `json:"tag"`
	IsDefault      bool   `json:"isDefault"`
	StartUnlocked  bool   `json:"startUnlocked"`
	HideIfNotOwned bool   `json:"hideIfNotOwned"`
	Image          string `json:"image"`
}

func DailyShop(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!dailyshop" || m.Content == "!ds" {
		req, err := http.NewRequest(
			http.MethodGet,
			"https://fortniteapi.io/v2/shop",
			nil,
		)
		if err != nil {
			log.Println(err)
			s.ChannelMessageSend(m.ChannelID, "Error creating HTTP request.")
		}

		req.Header.Set("accept", "application/json")
		req.Header.Set("Authorization", config.Env.FortniteApiKey)

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Println(err)
			s.ChannelMessageSend(m.ChannelID, "Error making HTTP request.")
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			log.Println(err)
			s.ChannelMessageSend(m.ChannelID, "Error reading response body.")
		}

		var shopResponse ShopResponse

		err = json.Unmarshal(body, &shopResponse)

		if err != nil {
			log.Println(err)
			s.ChannelMessageSend(m.ChannelID, "Error converting response body into struct.")
		}

		count := 0

		for _, item := range shopResponse.Shop {
			fmt.Println("----------------")
			fmt.Println("Item name:", item.DisplayName)
			fmt.Println("Item price:", item.Price.RegularPrice)
			fmt.Println("Item type:", item.DisplayType)
			fmt.Println("----------------")

			count += 1
		}

		fmt.Println("")
		fmt.Println("items count:", count)
	}
}
