package utils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/ocoderr/bsc-tool/abi/erc20"
	"github.com/ocoderr/bsc-tool/config"
	"github.com/shopspring/decimal"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
)

//var Client *ethclient.Client
const GweiUnit = 1000000000
const WeiUnit6 = 1000000
const WeiUnit8 = 100000000
const EthUnit = 1000000000000000000

var DefaultCallOpts = &bind.CallOpts{}

func EstimateClient(chain string) *ethclient.Client {
	//url := config.CF.Eth.RpcUrl
	//if chain == "eth" {
	//	url = config.CF.Eth.RpcUrl
	//} else if chain == "bsc" {
	//	url = config.CF.Bsc.RpcUrl
	//} else if chain == "heco" {
	//	url = config.CF.Heco.RpcUrl
	//} else if chain == "polygon" {
	//	url = config.CF.Polygon.RpcUrl
	//}
	client, err := ethclient.Dial(config.CF.Bsc.RpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func EstimateWsClient(chain string) *ethclient.Client {
	url := config.CF.Bsc.WsUrl
	//if chain == "eth" {
	//	url = config.CF.Eth.WsUrl
	//} else if chain == "bsc" {
	//	url = config.CF.Bsc.WsUrl
	//} else if chain == "heco" {
	//	url = config.CF.Heco.WsUrl
	//} else if chain == "polygon" {
	//	url = config.CF.Polygon.WsUrl
	//}
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GenerateOnePrivteKey() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(privateKeyBytes))
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return strings.ToLower(address), hexutil.Encode(privateKeyBytes)[2:]
}

func GenerateOneAccount(wallet *hdwallet.Wallet, i int) (string, string, string) {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
	account, _ := wallet.Derive(path, false)
	privateKeyHex, _ := wallet.PrivateKeyHex(account)
	address, _ := wallet.AddressHex(account)
	return path.String(), privateKeyHex, address
}

func EthBalance(arg string, Client *ethclient.Client) string {
	account := common.HexToAddress(arg)

	balance, err := Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	//fmt.Println(ethValue, "eth")
	return ethValue.String()

}
func Erc20Balance(arg, contract string, Client *ethclient.Client) (string, string) {
	account := common.HexToAddress(arg)
	tokenAddress := common.HexToAddress(contract)

	instance, err := erc20.NewErc20(tokenAddress, Client)
	if err != nil {
		log.Fatal(err)
	}
	balance, err := instance.BalanceOf(DefaultCallOpts, account)
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(DefaultCallOpts)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return symbol, ethValue.String()

}
func SendAllEth(fromAddresss string, toAddressS string, hexkey string, gas uint64, client *ethclient.Client) string {
	fromAddress := common.HexToAddress(fromAddresss)
	toAddress := common.HexToAddress(toAddressS)
	privateKey, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainIDInt := chainID.Int64()
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	if balance.Uint64() <= 0 {
		return ""
	}
	//f := amount * math.Pow10(18)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice := big.NewInt(int64(gas)).Mul(big.NewInt(int64(gas)), big.NewInt(int64(GweiUnit)))

	//if gas == 0 {
	//	_, fast, _ := GasPrice()
	//	gasPrice = big.NewInt(int64(fast))
	//}

	gasLimit := uint64(21000)
	//value := big.NewInt(int64(f))
	i := int64(balance.Uint64() - gasPrice.Uint64()*gasLimit)
	if i <= 0 {
		return ""
	}
	//log.Printf("balance %v\n",balance.Uint64())
	//log.Printf("gasPrice %v\n",gasPrice.Uint64())
	//log.Printf("i %v\n",i)

	//value := ToWei(i, int(18))
	//
	tx := types.NewTransaction(nonce, toAddress, big.NewInt(int64(i)), gasLimit, gasPrice, nil)
	signer := types.NewLondonSigner(chainID)
	if chainIDInt == 1 {
		signer = types.NewEIP155Signer(chainID)
	}

	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx.Hash().Hex()
	//return ""
}

func SendEth(fromPrivateKey string, toAddressS string, amount float64, gas uint64, Client *ethclient.Client) string {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress(toAddressS)

	//f := amount * math.Pow10(18)
	//value := big.NewInt(int64(f))
	value := ToWei(amount, int(18))

	//gasPrice := big.NewInt(int64(gas))
	gasPrice := big.NewInt(int64(gas)).Mul(big.NewInt(int64(gas)), big.NewInt(int64(GweiUnit)))

	//if gas == 0 {
	//	_, fast, _ := GasPrice()
	//	gasPrice = big.NewInt(int64(fast))
	//
	//}
	gasLimit := uint64(21000)
	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	//tx := types.NewTx(&types.DynamicFeeTx{
	//	ChainID: chainID,
	//	Nonce: nonce,
	//	GasFeeCap: gasPrice,
	//	GasTipCap: gasPrice,
	//	Gas: gasLimit,
	//	To: &toAddress,
	//	Value: value,
	//	Data: nil,
	//})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)

	if err != nil {
		panic(err)
	}
	err = Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}
	return signedTx.Hash().Hex()
}
func SendEthBatch(fromPrivateKey string, toAddressS string, amount float64, gas uint64, nonce uint64, Client *ethclient.Client) string {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(toAddressS)

	//f := amount * math.Pow10(18)
	//value := big.NewInt(int64(f))
	value := ToWei(amount, int(18))

	//gasPrice := big.NewInt(int64(gas))
	gasPrice := big.NewInt(int64(gas)).Mul(big.NewInt(int64(gas)), big.NewInt(int64(GweiUnit)))

	//if gas == 0 {
	//	_, fast, _ := GasPrice()
	//	gasPrice = big.NewInt(int64(fast))
	//
	//}
	gasLimit := uint64(21000)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx.Hash().Hex()
}
func SendErc20(fromPrivateKey string, toAddressS string, contractAddr string, amount float64, gas uint64, Client *ethclient.Client) string {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	contract := common.HexToAddress(contractAddr)
	toAddress := common.HexToAddress(toAddressS)
	//gasPrice := big.NewInt(int64(gas))
	gasPrice := big.NewInt(int64(gas)).Mul(big.NewInt(int64(gas)), big.NewInt(int64(GweiUnit)))

	//if gas == 0 {
	//	_, fast, _ := GasPrice()
	//	gasPrice = big.NewInt(int64(fast))
	//
	//}
	//gasLimit := uint64(21000)

	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)    // in wei
	auth.GasLimit = uint64(55000) // in units
	auth.GasPrice = gasPrice

	erc20, err := erc20.NewErc20(contract, Client)
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := erc20.Decimals(DefaultCallOpts)
	if err != nil {
		decimals = 18
	}

	//f := amount * math.Pow10(int(decimals))
	//value := big.NewInt(int64(f))
	value := ToWei(amount, int(decimals))

	signedTx, err := erc20.Transfer(auth, toAddress, value)

	return signedTx.Hash().Hex()
}

func ApproveErc20(fromPrivateKey string, spender string, contractAddr string, amount float64, gas uint64, Client *ethclient.Client) string {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	contract := common.HexToAddress(contractAddr)
	spenderAddress := common.HexToAddress(spender)
	//gasPrice := big.NewInt(int64(gas))
	gasPrice := big.NewInt(int64(gas)).Mul(big.NewInt(int64(gas)), big.NewInt(int64(GweiUnit)))

	//if gas == 0 {
	//	_, fast, _ := GasPrice()
	//	gasPrice = big.NewInt(int64(fast))
	//
	//}
	//gasLimit := uint64(21000)

	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)    // in wei
	auth.GasLimit = uint64(44430) // in units
	auth.GasPrice = gasPrice

	erc20, err := erc20.NewErc20(contract, Client)
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := erc20.Decimals(DefaultCallOpts)
	if err != nil {
		decimals = 18
	}

	//f := amount * math.Pow10(int(decimals))
	//value := big.NewInt(int64(f))
	value := ToWei(amount, int(decimals))

	signedTx, err := erc20.Approve(auth, spenderAddress, value)

	hex := signedTx.Hash().Hex()
	return hex
}
func ApproveErc20AndAwait(fromPrivateKey string, spender string, contractAddr string, amount float64, gas uint64, client *ethclient.Client) string {
	tx := ApproveErc20(fromPrivateKey, spender, contractAddr, amount, gas, client)
	//finish, err := WaitForTxFinish(common.HexToHash(tx), Client)
	//for !WaitForTxFinish(common.HexToHash(tx), Client) {
	//	time.Sleep(1 * time.Second)
	//}
	return tx
}

func WaitForTxFinish(hex common.Hash, client *ethclient.Client) (bool, error) {
	tx, err := client.TransactionReceipt(context.Background(), hex)
	if err != nil {
		return false, nil
	}
	status := int(tx.Status)
	if status != 1 {
		return true, errors.New("failed")
	}
	return true, nil
}
func SendErc20All(fromPrivateKey string, toAddressS string, contractAddr string, gas uint64, client *ethclient.Client) string {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	contract := common.HexToAddress(contractAddr)
	toAddress := common.HexToAddress(toAddressS)
	//gasPrice := big.NewInt(int64(gas))
	gasPrice := big.NewInt(int64(gas)).Mul(big.NewInt(int64(gas)), big.NewInt(int64(GweiUnit)))

	//if gas == 0 {
	//	_, fast, _ := GasPrice()
	//	gasPrice = big.NewInt(int64(fast))
	//
	//}
	//gasLimit := uint64(21000)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)    // in wei
	auth.GasLimit = uint64(65000) // in units
	auth.GasPrice = gasPrice

	erc20, err := erc20.NewErc20(contract, client)
	if err != nil {
		log.Fatal(err)
	}

	value, err := erc20.BalanceOf(DefaultCallOpts, fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := erc20.Transfer(auth, toAddress, value)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx.Hash().Hex()
}

func SendErc20Batch(fromPrivateKey string, toAddressS string, contractAddr string, amount float64, nonce uint64, gas uint64, Client *ethclient.Client) string {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	contract := common.HexToAddress(contractAddr)
	toAddress := common.HexToAddress(toAddressS)
	//gasPrice := big.NewInt(int64(gas))
	gasPrice := big.NewInt(int64(gas)).Mul(big.NewInt(int64(gas)), big.NewInt(int64(GweiUnit)))

	//if gas == 0 {
	//	_, fast, _ := GasPrice()
	//	gasPrice = big.NewInt(int64(fast))
	//
	//}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)    // in wei
	auth.GasLimit = uint64(55000) // in units
	auth.GasPrice = gasPrice

	erc20, err := erc20.NewErc20(contract, Client)
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := erc20.Decimals(DefaultCallOpts)
	//log.Println(decimals)
	if err != nil {
		log.Fatal(err)
	}
	value := ToWei(amount, int(decimals))

	signedTx, err := erc20.Transfer(auth, toAddress, value)

	return signedTx.Hash().Hex()
	//return "";
}

func Erc20Token(contractAddr string, Client *ethclient.Client) *erc20.Erc20 {
	contract := common.HexToAddress(contractAddr)
	erc20, err := erc20.NewErc20(contract, Client)
	if err != nil {
		log.Fatal(err)
	}
	return erc20
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func Str2Float(amount string) float64 {
	if amount == "" {
		return 0.0
	}
	pendingRewardFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0.0
	}
	return pendingRewardFloat
}
func Str2Int(amount string) int64 {
	if amount == "" {
		return 0.0
	}
	pendingRewardFloat, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return 0.0
	}
	return pendingRewardFloat
}
