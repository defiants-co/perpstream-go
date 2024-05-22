package utils

const (
	KwentaSubgraphUrl    = "https://api.thegraph.com/subgraphs/name/kwenta/optimism-perps"
	KwentaDataSourceName = "kwenta-v2-optimism"
	KwentaPriceDataUrl   = "https://hermes.pyth.network/api/latest_price_feeds?"
)

var KwentaMarketsToToken map[string]string = map[string]string{
	"0x534f4c0000000000000000000000000000000000000000000000000000000000": "SOL",
	// "0x5045504500000000000000000000000000000000000000000000000000000000": "PEPE",
	"0x7345544800000000000000000000000000000000000000000000000000000000": "ETH",
	"0x7342544300000000000000000000000000000000000000000000000000000000": "BTC",
	"0x4156415800000000000000000000000000000000000000000000000000000000": "AVAX",
	"0x444f474500000000000000000000000000000000000000000000000000000000": "DOGE",
}

var KwentaPriceIdsToToken map[string]string = map[string]string{
	"ff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace": "ETH",
	"e62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43": "BTC",
	// "944f2f908c5166e0732ea5b610599116cd8e1c41f47452697c1e84138b7184d6": "PEPE",
	"ef0d8b6fda2ceba41da15d4095d1da392a0d2f8ed0c6c7bc0f4cfac8c280b56d": "SOL",
	"856aac602516addee497edf6f50d39e8c95ae5fb0da1ed434a8c2ab9c3e877e9": "AVAX",
	"dcef50dd0a4cd2dcc17e45df1676dcb336a11a61c69df7a0299b0150c672d25c": "DOGE",
}

var KwentaPythDecimals map[string]int = map[string]int{
	"BTC":  8,
	"ETH":  8,
	"AVAX": 6,
	"DOGE": 8,
	"SOL":  8,
	// "PEPE": 13,
}

var KwentaMarkets []string = GetKeys(KwentaMarketsToToken)
