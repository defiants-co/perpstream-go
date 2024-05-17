package utils

// GMX related constants
const (
	BaseArbitrumRpcUrl          = "https://arb1.arbitrum.io/rpc"
	GmxReaderContractAddress    = "0xdA5A70c885187DaA71E7553ca9F728464af8d2ad"
	GmxDataStoreContractAddress = "0xFD70de6b91282D8017aA4E741e9Ae325CAb992d8"
	GmxPricesUrl                = "https://arbitrum-api.gmxinfra.io/signed_prices/latest"
	GmxDataSourceName           = "gmx-v2-arbitrum"
)

// MarketHashTable
var GmxAddressToCollateralToken = map[string]string{
	"0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f": "BTC",
	"0xaf88d065e77c8cC2239327C5EDb3A432268e5831": "USDC",
	"0x82aF49447D8a07e3bd95BD0d56f35241523fBab1": "ETH",
	"0x2bcC6D6CdBbDC0a4071e48bb3B969b06B3330c07": "SOL",
	"0x912CE59144191C1204E64559FE8253a0e49E6548": "ARB",
	"0xf97f4df75117a78c1A5a0DBb814Af92458539FB4": "LINK",
	"0xFa7F8980b0f1E64A2062791cc3b0871572f1F7f0": "UNI",
	"0x565609fAF65B92F7be02468acF86f8979423e514": "AVAX",
	"0xa9004A5421372E1D83fB1f85b0fc986c912f91f3": "BNB",
	"0xba5DdD1f9d7F570dc94a51479a000E3BCE967196": "AAVE",
	"0xaC800FD6159c2a2CB8fC31EF74621eB430287a5A": "OP",
	"0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9": "USDT",
	"0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1": "DAI",
	"0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8": "USDC.e",
	// Synthetics backed by WETH
	// "": "DOGE",
	// "": "XRP",
	// "": "NEAR",
	// "": "LTC",
	// "": "ATOM",
}

var CollateralTokenToGmxAddress = reverseMap(GmxAddressToCollateralToken)

var GmxAddressToMarket = map[string]string{
	"0x7C11F78Ce78768518D743E81Fdfa2F860C6b9A77": "BTC-USD",
	"0x47c031236e19d024b42f8AE6780E44A573170703": "BTC-USD",
	"0x450bb6774Dd8a756274E0ab4107953259d2ac541": "ETH-USD",
	"0x70d95587d40A2caf56bd97485aB3Eec10Bee6336": "ETH-USD",
	"0x09400D9DB990D5ed3f35D7be61DfAEB900Af03C9": "SOL-USD",
	"0xC25cEf6061Cf5dE5eb761b50E4743c1F5D7E5407": "ARB-USD",
	"0x7f1fa204bb700853D36994DA19F830b6Ad18455C": "LINK-USD",
	"0x6853EA96FF216fAb11D2d930CE3C508556A4bdc4": "DOGE-USD",
	"0xc7Abb2C5f3BF3CEB389dF0Eecd6120D451170B50": "UNI-USD",
	"0x7BbBf946883a5701350007320F525c5379B8178A": "AVAX-USD",
	"0x0CCB4fAa6f1F1B30911619f1184082aB4E25813c": "XRP-USD",
	"0x63Dc80EE90F26363B3FCD609007CC9e14c8991BE": "NEAR-USD",
	"0xD9535bB5f58A1a75032416F2dFe7880C30575a41": "LTC-USD",
	"0x2d340912Aa47e33c90Efb078e69E70EFe2B34b9B": "BNB-USD",
	"0xB686BcB112660343E6d15BDb65297e110C8311c4": "USDC-USDT",
	"0x1CbBa6346F110c8A5ea739ef2d1eb182990e4EB2": "AAVE-USD",
	"0x4fDd333FF9cA409df583f306B6F5a7fFdE790739": "OP-USD",
	"0xe2fEDb9e6139a182B98e7C2688ccFa3e9A53c665": "USDC-DAI",
	"0x9C2433dFD71096C435Be9465220BB2B189375eA7": "USDC-USDC.e",
}

var GmxMarketToAddress = reverseMap(GmxAddressToMarket)

var GmxCollateralTokenDecimals = map[string]int{
	"BTC":    8,
	"ETH":    18,
	"SOL":    9,
	"USDC":   6,
	"ARB":    18,
	"LINK":   18,
	"UNI":    18,
	"AVAX":   18,
	"BNB":    18,
	"OP":     18,
	"DAI":    18,
	"USDC.e": 6,
	"USDT":   6,
}

var GmxSyntheticsTokenDecimals = map[string]int{
	"DOGE": 8,
	"ATOM": 6,
	"NEAR": 24,
	"XRP":  6,
	"LTC":  8,
}
var MarketToDecimals = AppendMaps(GmxCollateralTokenDecimals, GmxSyntheticsTokenDecimals)

var GmxPricesDecimals = map[string]int{}
