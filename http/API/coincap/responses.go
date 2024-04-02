package coincap

import "fmt"

type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct {
	ID                string `json:"id"`
	Rank              string `json:"Rank"`
	Symbol            string `json:"Symbol"`
	Name              string `json:"Name"`
	Supply            string `json:"Supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

func (d Asset) String() string {
	return fmt.Sprintf("[ID] %-23s | [RANK] %3s |  [NAME] %-20s | [PRICE] $%s", d.ID, d.Rank, d.Name, d.PriceUsd)
}
