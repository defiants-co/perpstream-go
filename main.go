package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/defiants-co/perpstream-go/clients"
	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/utils"
)

type PositionDataList []models.FuturesPosition

func (p PositionDataList) MarshalJSON() ([]byte, error) {
	if len(p) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal([]models.FuturesPosition(p))
}

type Payload struct {
	UserId           string           `json:"user_id"`
	DataSource       string           `json:"data_source"`
	PositionDataList PositionDataList `json:"position_data"`
}

func SendCallback(positionsData []models.FuturesPosition, userId string, dataSource string) {
	payload := Payload{
		UserId:           userId,
		DataSource:       dataSource,
		PositionDataList: positionsData,
	}
	// Convert the Payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://wailing-nail-20.webhook.cool", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return
	}

	return

}

// Define the list of RPC URLs
var rpcURLs = []string{
	"https://arbitrum.llamarpc.com",
	"https://arbitrum.drpc.org",
	"https://arbitrum-one.publicnode.com",
	"https://rpc.ankr.com/arbitrum",
	"https://arb-pokt.nodies.app",
	"https://arbitrum-one-rpc.publicnode.com",
	"https://arbitrum.meowrpc.com",
	"https://arbitrum.rpc.subquery.network/public",
	"https://arb-mainnet-public.unifra.io",
}

var users = []string{
	"0x00000000007Cee88e01241C9Db0A02b859DB5B52",
	"0x0189c952c7e938adbf3bBcf70715B8e0964c9806",
	"0x03020B304D91A253B18852eA62af4Ec43417BC49",
	"0x0514F2f3E0277c47117e3f33d9390EfB0aCfbdfe",
	"0x0E9b4752da25888B2605a7314AB708FFCD01602d",
	"0x17d9a00Bb9421eF1Be6C5Bfd510D6d25588D58D2",
	"0x18dbb3C7b3F35a07391c82c31BED414a2b131F0E",
	"0x1c379BfE42E54b71768f9207c639374F1A41242E",
	"0x1dc67aDaf9f163c5b1A95043addBDFA1AF58e512",
	"0x1fb038578B0E53E19F6b730260Cf2Cd747730a3B",
	"0x275D837100148F72aF54d994AC3b995Ed7c0e6c0",
	"0x27D8c777eA58698f4539A2B1A4988B32515C8d6E",
	"0x299A970A63058EF4D3fbaB4f3e89566a93e95fc2",
	"0x2F7DD724A3F1c18734F4bF5F743856d19af4a296",
	"0x338E969bc1aa7EEC01eCbA0940e30a2Ac3CdF6B4",
	"0x3B990f185fA49ABA52555c8b6B3A91FDD1Bc64A1",
	"0x3CEd6B69f632B01C52b69164eF8a3468572897cE",
	"0x3ab532Ae7341c8c800A69Dd63557393e88bcb3CC",
	"0x41cF64CE7bd0d4D13A9De8171b869cA7FE152Fa1",
	"0x429A69Bb14cAF03cd4AfCA8F78b91b0395714489",
	"0x4612Aed50781f84E0BfB6C639AE800652b671d4a",
	"0x493aFFC3aca7326D5aBb402b5a87A74Bf3D703E0",
	"0x4Cd80aa0CE4881Eb8679EdA1f6fbe3d89AEc0F7F",
	"0x4D8c173B203035B3DFDe7dd40e7dB728c47d98a8",
	"0x4E65359d82BEb8f6bA9cf0871e6de1FA0C4783fc",
	"0x4e2FE4eB001708f35AaCec15307386Da6Fbc09Ea",
	"0x5CE6f3798B9ca0797E1027E9b86E7dF0Ba61E593",
	"0x5ED7B6E7C87Ddaa3c7E05fcF8f539D65E4400E84",
	"0x5e216ceCB65E1E1B86fE8C46c730af287c4492Dc",
	"0x5e5de26E14a711d6Eb21333C24f1acEC9fA0A2C1",
	"0x60fe5Cbd886A778f584FFCC63833B068104D1f77",
	"0x6416d46Fa8f6139988e7d7F10314E073B1B1F07E",
	"0x647e587524F63Bdf2C3cA95A687e0ed9c24acc50",
	"0x66663fc442C3A09690971EBD17cc98Ad5e0744F8",
	"0x69b45FDb3dc1684CD323b8308e7D5ec2ac5aBe5C",
	"0x6F601156d82d6aeAaf82a026376C1064d3Ce61C7",
	"0x713a64a32000D1b918fc81078FE94CA9EddE1C99",
	"0x7236cB2a261a3BE528D559395E70830b8B366C56",
	"0x7665A46c97E62298CC4381327DAFbCc16ff9dFe1",
	"0x82a376358F29B18018A9c18A76cF64257863B375",
	"0x843a5aE61733A17B5560777448E98a2553E92F07",
	"0x8F5db667276e9d805bF5aDb315374f8fa299699E",
	"0x8e88687CabFB21Fa7Eea14fc445e705D9640B5f6",
	"0x975f178fD8A566d98bc935340064c691Fa4d8881",
	"0x987aC0eBf4708FA60E4153B85BF706AB0268d9ca",
	"0x9F3A95d79c04f56740c914ab143c30f9554D2BDe",
	"0xA53dE273DfBF602FB6fa880469B3E5857f34336e",
	"0xAD837b4fb104B84dc5Cd28edCAcb81865D5f14Ac",
	"0xAb709116e3E2D4fA8248e23C60f1330541c2936d",
	"0xAcc5AcE4118eCe3c6006e2A51F65C7DFd1286DAF",
	"0xB47a339A5F3a9F3c3F68c0493E28A89B829E81C3",
	"0xB4e9D3Ad7334Ea877F04E9ad6dD1dfBE5EfE7C8b",
	"0xB53ED4b7B068d41b590927650f4dFB2F797C3b22",
	"0xC0F86af87512993055227539725b484e627010FF",
	"0xCb3B3B7EFa59E1ED758895d64b658047244155F2",
	"0xCeeC48581B3145a575508719f45da07dc57fA7ce",
	"0xD0A4116352091748BaE29646A800a17eCDcC79e4",
	"0xD0b8C8ce7B43bd45cc0D114041199e49B21F2222",
	"0xD441C97eF1458d847271f91714799007081494eF",
	"0xD62E61e66eCc17ffaFBeff103E9A04C6D29434d5",
	"0xD6BB09C5503FEc718351bC1104F3bE267aC571B8",
	"0xDB8b2a92994fBE8199586312D8Cbb56E378F1F1f",
	"0xE23BF044C30FCF7A4Ea427564B2ef2B782268FEa",
	"0xF8421AdBbE960925a70fF9a59ff5392669139182",
	"0xF94176AF4A3698EB667b32Ea61A3a0Ec2003Bf74",
	"0xa160d4584EaC8fF04F3fBAF7D1FA5C8937E90780",
	"0xaC250e3d9298Fc03154314AE6786a5dB977B270c",
	"0xaDae0719e0d9bD57Ced687deF67050634db033c3",
	"0xaa51589f27cC4A35065c438B6C8d35c861955482",
	"0xaea8E3Bd369217CC6E3e6AbdDf0dA318fBA8E59b",
	"0xb7d672703E7987715912A0784Be91B27D1098C89",
	"0xc9e1CE91d3f782499cFe787b6F1d2AF0Ca76C049",
	"0xcFa0AB31d240e4E831c6219F62E73a38C6AAB058",
	"0xccD724B9af92A3A389150334556330EfD3f59487",
	"0xcef23Bfd91809E2c674b7256c94B6141bF285ad4",
	"0xd0ce74f8A7f5FbbA9455c489165cFb41Fc6681f5",
	"0xd468808cC9e30f0Ae5137805fff7ffB213984250",
	"0xd801AF63b71c806597132fFc2f1c46985cEeE392",
	"0xdAcb8697d9B7216A7172b1E8Ba8b1203D42f37BC",
	"0xe603D730d53C6d407cd5eFb4EE4334301B7C0015",
	"0xf121B177dD32E29Af3bFEf3F58c47e4eFa9E0Ef2",
	"0xf4f1A5679D113DC93732C89b69ee5Ce5d12D18F2",
	"0xf55B92b635bb745F65Be9f9E84a5509905298aAF",
	"0xf661458656F8164bDe13849b078Dd6367ba86978",
	"0xfeF078C6654666B93837807ac3F6e64718D424A2",
}

func main() {
	priceCache := utils.NewPriceCache()
	go priceCache.StreamPrices(10, false)

	client, err := clients.NewGmxClient(rpcURLs, priceCache)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	for _, userId := range users {
		go StreamTracker(client, userId, false, 5, true, SendCallback)
	}
	wg.Wait()
}

var count = 0

func StreamTracker(
	client *clients.GmxClient,
	userId string,
	debug bool,
	sleepSeconds float64,
	initWithCallback bool,
	callback func(
		newPositions []models.FuturesPosition,
		userId string,
		dataSource string,
	),
) error {
	id := count

	count++
	var innerCount = 0

	myfunc := func(
		newPositions []models.FuturesPosition,
		userId string,
		dataSource string,
	) {
		callback(newPositions, userId, dataSource)
		innerCount++
		fmt.Println(fmt.Sprintf("round %d - id: %d", innerCount, id))
	}

	client.StreamPositions(userId, debug, sleepSeconds, initWithCallback, myfunc)

	return nil
}
