package config

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
	"strings"
)

type YamlConfig struct {
	Language      string
	Bsc           NodeConfig
	PancakeRouter string
	FromAddress   string
	//PrivateKey    string
	//Mnemonic      string
	BscToken  map[string]string
	SyrupPool map[string]string
}

type NodeConfig struct {
	RpcUrl string
	WsUrl  string
}

var CF YamlConfig

func NewConfig(file string) {
	cf := YamlConfig{
		Language:      "EN",
		PancakeRouter: "0x10ed43c718714eb63d5aa57b78b54704e256024e",
		Bsc: NodeConfig{
			RpcUrl: "https://bsc-dataseed1.defibit.io/",
			WsUrl:  "wss://bsc-ws-node.nariox.org:443",
		},
		SyrupPool: map[string]string{
			"cake":      "0x73feaa1eE314F8c655E354234017bE2193C9E24E",
			"cake-auto": "0xa80240Eb5d7E05d3F250cF000eEc0891d00b51CC",
			"dar":       "0x9b861A078B2583373A7a3EEf815bE1A39125Ae08",
			"idia":      "0x07984abb7489cd436d27875c07eb532d4116795a",
			"kart":      "0x73bB10B89091f15e8FeD4d6e9EBa6415df6acb21",
			"porto":     "0xdD52FAB121376432DBCBb47592742F9d86CF8952",
			"qi":        "0xbd52ef04DB1ad1c68A8FA24Fa71f2188978ba617",
			"quidd":     "0xd97ee2bfe79a4d4ab388553411c462fbb536a88c",
			"santos":    "0x0914b2d9d4dd7043893def53ecfc0f1179f87d5c",
			"sfund":     "0x7f103689cabe17c2f70da6faa298045d72a943b8",
			"xcv":       "0xf1fa41f593547e406a203b681df18accc3971a43",
			"zoo":       "0x2EfE8772EB97B74be742d578A654AB6C95bF18db",
		},
		BscToken: map[string]string{
			"busd":    "0xe9e7cea3dedca5984780bafc599bd69add087d56",
			"cake":    "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
			"dar":     "0x23CE9e926048273eF83be0A3A8Ba9Cb6D45cd978",
			"eternal": "0xd44fd09d74cd13838f137b590497595d6b3feea4",
			"idia":    "0x0b15ddf19d47e6a86a56148fb4afffc6929bcb89",
			"kart":    "0x8bdd8dbcbdf0c066ca5f3286d33673aa7a553c10",
			"porto":   "0x49f2145d6366099e13b10fbf80646c0f377ee7f6",
			"qi":      "0x8729438eb15e2c8b576fcc6aecda6a148776c0f5",
			"quidd":   "0x7961ade0a767c0e5b67dd1a1f78ba44f727642ed",
			"santos":  "0xa64455a4553c9034236734faddaddbb64ace4cc7",
			"sfund":   "0x477bc8d23c634c154061869478bce96be6045d12",
			"usdt":    "0x55d398326f99059ff775485246999027b3197955",
			"wbnb":    "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
			"xcv":     "0x4be63a9b26EE89b9a3a13fd0aA1D0b2427C135f8",
			"zoo":     "0x1D229B958D5DDFca92146585a8711aECbE56F095",
		},
	}
	bytes, err := yaml.Marshal(cf)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(file, bytes, 0777)
	if err != nil {
		panic(err)
	}
}
func BscToken(symbol string) string {
	return CF.BscToken[strings.ToLower(symbol)]
}
func BscTokenValueIndex(index int) string {
	var names []string
	for name := range CF.BscToken {
		names = append(names, name)
	}
	sort.Strings(names)

	return CF.BscToken[strings.ToLower(names[index])]
}
func BscTokenKeyIndex(index int) string {
	var names []string
	for name := range CF.BscToken {
		names = append(names, name)
	}
	sort.Strings(names)

	if index < len(names) {
		return names[index]
	}

	return ""
}

func SyrupPool(symbol string) string {
	return CF.SyrupPool[strings.ToLower(symbol)]
}
func SyrupPoolValueIndex(index int) string {
	var names []string
	for name := range CF.SyrupPool {
		names = append(names, name)
	}
	sort.Strings(names)

	return CF.SyrupPool[strings.ToLower(names[index])]
}
func SyrupPoolKeyIndex(index int) string {
	var names []string
	for name := range CF.SyrupPool {
		names = append(names, name)
	}
	sort.Strings(names)

	if index < len(names) {
		return names[index]
	}

	return ""
}
func SyrupPoolKeys() []string {
	var names []string
	names = append(names, "cake-auto")
	names = append(names, "cake")
	for name := range CF.SyrupPool {
		if name == "cake" || name == "cake-auto" {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)

	return names
}

var ZhTextMap = map[string]string{
	"app.title":          "People",
	"menu.show.title":    "通知",
	"menu.show.message":  "请重新打开",
	"menu.language":      "语言",
	"menu.language.zh":   "中文",
	"menu.language.en":   "English",
	"main.tab.syrup":     "Pancake 糖浆池",
	"main.tab.account":   "BSC",
	"main.tab.setting":   "配置",
	"main.field.address": "地址",
	"main.tab.transfers": "向多地址转账",
}
var EnTextMap = map[string]string{
	"app.title":          "People",
	"menu.show.title":    "Info",
	"menu.show.message":  "Please Start App Again",
	"menu.language":      "Language",
	"menu.language.zh":   "中文",
	"menu.language.en":   "English",
	"main.tab.syrup":     "Pancake Syrup Pool",
	"main.tab.account":   "BSC",
	"main.tab.setting":   "Setting",
	"main.field.address": "Address",
	"main.tab.transfers": "Transfer To Many",
}

func SaveValue(key, value string) {
	viper.Set(key, value)
	viper.WriteConfig()

}

func GetText(key string) string {
	if CF.Language == "ZH" {
		value, ok := ZhTextMap[key]
		if !ok {
			return "未定义文本"
		}
		return value
	} else {
		value, ok := EnTextMap[key]
		if !ok {
			return "Undefined text"
		}
		return value

	}
}
