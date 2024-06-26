package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/defiants-co/perpstream-go/models"
)

var kwentaAccounts []string = []string{
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x22866c5c7f2b5475cff41465c53aa813b4c22b13",
	"0xe8c19db00287e3536075114b2576c70773e039bd",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0xf78310ed6641e6c4e221e9d676440ac8645d3afe",
	"0x30391a4f9d2f099d41888f811784281cba4097f0",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x92812499ff2c040f93121aab684680a6e603c4a7",
	"0x30391a4f9d2f099d41888f811784281cba4097f0",
	"0x92812499ff2c040f93121aab684680a6e603c4a7",
	"0x3b7424d5cc87dc2b670f4c99540f7380de3d5880",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x44322e2f396e4e5244f4f09b863044847e4a54da",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0xeed09cc4ebf3fa599eb9ffd7a280e7b944b436b7",
	"0x8af700ba841f30e0a3fcb0ee4c4a9d223e1efa05",
	"0x1dc67adaf9f163c5b1a95043addbdfa1af58e512",
	"0x795f50722cf5ad82f78dda8dc8f7b235332977c3",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0xfee4700a35676fbdbaca75b7bd5cc7bb2abf9910",
	"0xd120cf3e0408dd794f856e8ca2a23e3396a9b687",
	"0x27cc4d6bc95b55a3a981bf1f1c7261cda7bb0931",
	"0x88965659f799fcb2f70937c31180b0acb7106ac5",
	"0x92812499ff2c040f93121aab684680a6e603c4a7",
	"0xd1bbaa9eab9a5f51ee3b68bb6ff3c1f60a8c0add",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x30391a4f9d2f099d41888f811784281cba4097f0",
	"0x795f50722cf5ad82f78dda8dc8f7b235332977c3",
	"0x92812499ff2c040f93121aab684680a6e603c4a7",
	"0x92812499ff2c040f93121aab684680a6e603c4a7",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x1dc67adaf9f163c5b1a95043addbdfa1af58e512",
	"0xccfa0530b9d52f970d1a2daea670ce58e4176389",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x44322e2f396e4e5244f4f09b863044847e4a54da",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0xd1bbaa9eab9a5f51ee3b68bb6ff3c1f60a8c0add",
	"0xeed09cc4ebf3fa599eb9ffd7a280e7b944b436b7",
	"0x565a65432ca44a999eb7217815d58594a559ccb2",
	"0xeed09cc4ebf3fa599eb9ffd7a280e7b944b436b7",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0xeed09cc4ebf3fa599eb9ffd7a280e7b944b436b7",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x8e27d64063d74c7c2f7c8609e5b6d78d03378d23",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x63a94473cbc5329e5dc0b070ef7586d910060670",
	"0x0514f2f3e0277c47117e3f33d9390efb0acfbdfe",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x1dc67adaf9f163c5b1a95043addbdfa1af58e512",
	"0xe8c19db00287e3536075114b2576c70773e039bd",
	"0xeed09cc4ebf3fa599eb9ffd7a280e7b944b436b7",
	"0xe8c19db00287e3536075114b2576c70773e039bd",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0xab281460020aa77fd533245c83660948152e9b46",
	"0xb2ceb04874fce4cefd38c751c8ed7b88d2185c4d",
	"0x1c1e747a6be850549e9655addf59fd9e7cc2d4dc",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0xeed09cc4ebf3fa599eb9ffd7a280e7b944b436b7",
	"0xd1bbaa9eab9a5f51ee3b68bb6ff3c1f60a8c0add",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x795f50722cf5ad82f78dda8dc8f7b235332977c3",
	"0x236f968b3e9cc45d6d86bc95386a27c8051ab1ce",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0xeed09cc4ebf3fa599eb9ffd7a280e7b944b436b7",
	"0x565a65432ca44a999eb7217815d58594a559ccb2",
	"0x1c1e747a6be850549e9655addf59fd9e7cc2d4dc",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0xeed09cc4ebf3fa599eb9ffd7a280e7b944b436b7",
	"0x1d4acbc9f70cadd4e6eb215731b8a20bb848d042",
	"0x09cf50574504d9dcf127e848a6058e8e0bb814aa",
	"0x30391a4f9d2f099d41888f811784281cba4097f0",
	"0x236f968b3e9cc45d6d86bc95386a27c8051ab1ce",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x4dc589ee79ebb4c457c8d35555570ccfbd247e36",
	"0x3b7424d5cc87dc2b670f4c99540f7380de3d5880",
	"0xcd4ddf2aed875248d6d586b57146858a71eafc02",
	"0x795f50722cf5ad82f78dda8dc8f7b235332977c3",
	"0x30391a4f9d2f099d41888f811784281cba4097f0",
	"0x565a65432ca44a999eb7217815d58594a559ccb2",
	"0x360537542135943e8fc1562199aea6d0017f104b",
	"0x31d3243cfb54b34fc9c73e1cb1137124bd6b13e1",
	"0x160a8d28d63961297e51e4f1a0401bf4d27f5162",
	"0xd120cf3e0408dd794f856e8ca2a23e3396a9b687",
	"0xd1bbaa9eab9a5f51ee3b68bb6ff3c1f60a8c0add",
	"0x4dd5d939486b23aabcfd08a063811e42f7c2b6a9",
	"0x92812499ff2c040f93121aab684680a6e603c4a7",
	"0xe996148026c14ec87741d7d34592035fe61967f7",
	"0x2b19fde5d7377b48be50a5d0a78398a496e8b15c",
	"0xccfa0530b9d52f970d1a2daea670ce58e4176389",
	"0x1e5c9ef6a07f8ee36e21c0c2737c9b19f5de7f2a",
}
var Users []string = []string{
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

var Rpcs []string = []string{
	"https://arbitrum.llamarpc.com",
	"https://rpc.ankr.com/arbitrum",
	"https://arbitrum-one-rpc.publicnode.com",
	"https://arb-pokt.nodies.app",
	"https://arb-mainnet-public.unifra.io",
	"https://arbitrum.meowrpc.com",
	"https://arbitrum-one.publicnode.com",
	"https://arbitrum.rpc.subquery.network/public",
	"https://1rpc.io/arb",
}

func ConvertFuturesPositionsToInterface(positions []models.FuturesPosition) []interface{} {
	data := make([]interface{}, len(positions))
	for i, pos := range positions {
		data[i] = pos
	}
	return data
}

func ConvertOptionPositionsToInterface(positions []models.OptionPosition) []interface{} {
	data := make([]interface{}, len(positions))
	for i, pos := range positions {
		data[i] = pos
	}
	return data
}

func GmxCallback(
	oldPositions []models.FuturesPosition,
	newPositions []models.FuturesPosition,
	userId string,
	dataSource string,
) {

	if len(newPositions) > 0 {
		SendWebhook(userId, dataSource, ConvertFuturesPositionsToInterface(oldPositions), ConvertFuturesPositionsToInterface(newPositions))
	}
}

func HegicCallback(
	oldPositions []models.OptionPosition,
	newPositions []models.OptionPosition,
	userId string,
	dataSource string,
) {
	if len(newPositions) > 0 {
		SendWebhook(userId, dataSource, ConvertOptionPositionsToInterface(oldPositions), ConvertOptionPositionsToInterface(newPositions))
	}
}

type Payload struct {
	UserId      string        `json:"user_id"`
	DataSource  string        `json:"data_source"`
	OldPosition []interface{} `json:"old_positions"`
	NewPosition []interface{} `json:"new_positions"`
}

func SendWebhook(userId string, dataSource string, old []interface{}, new []interface{}) {
	payload := Payload{
		UserId:      userId,
		DataSource:  dataSource,
		OldPosition: old,
		NewPosition: new,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	webhookURL := "https://enccigryc0gjm.x.pipedream.net" // Replace with your actual webhook URL
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Non-OK HTTP status: %s\n", resp.Status)
	} else {
		fmt.Println("Webhook sent successfully!")
	}

}

func WriteStructsToFile(filename string, data []interface{}) error {
	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	// Create a JSON encoder and set indentation for pretty print
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// Encode the data to JSON and write to the file
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("could not encode data to JSON: %v", err)
	}

	return nil
}
