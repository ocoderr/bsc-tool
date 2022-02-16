package utils

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ocoderr/bsc-tool/abi/uniswapV2Pair"
	"github.com/ocoderr/bsc-tool/abi/uniswapv2factory"
	"github.com/ocoderr/bsc-tool/abi/uniswapv2router02"
	"github.com/ocoderr/bsc-tool/config"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func CalculatePath(swapFromToken string, swapToToken string, router string, client *ethclient.Client) ([]common.Address, error) {

	var swapFromTokenAddress string
	if strings.ToLower(swapFromToken) == "bnb" {
		swapFromTokenAddress = config.BscToken("wbnb")
	} else {
		swapFromTokenAddress = config.BscToken(swapFromToken)
	}
	var swapToTokenAddress string
	if strings.ToLower(swapToToken) == "bnb" {
		swapToTokenAddress = config.BscToken("wbnb")
	} else {
		swapToTokenAddress = config.BscToken(swapToToken)
	}
	var path []common.Address
	fromToken := common.HexToAddress(swapFromTokenAddress)
	toToken := common.HexToAddress(swapToTokenAddress)

	bnbToken := common.HexToAddress(config.BscToken("wbnb"))
	busdToken := common.HexToAddress(config.BscToken("busd"))
	usdtToken := common.HexToAddress(config.BscToken("usdt"))

	path = append(path, fromToken)

	_, fromWbnbPair, _ := GetPoolPair(router, fromToken, bnbToken, client)
	_, fromBusdPair, _ := GetPoolPair(router, fromToken, busdToken, client)
	_, fromUsdtPair, _ := GetPoolPair(router, fromToken, usdtToken, client)

	_, toWbnbPair, _ := GetPoolPair(router, toToken, bnbToken, client)
	_, toBusdPair, _ := GetPoolPair(router, toToken, busdToken, client)
	_, toUsdtPair, _ := GetPoolPair(router, toToken, usdtToken, client)

	if swapFromTokenAddress == config.BscToken("wbnb") {
		if toWbnbPair != nil {
			path = append(path, toToken)
			return path, nil
		}
		if toBusdPair != nil {
			path = append(path, busdToken)
			path = append(path, toToken)
			return path, nil
		}
		if toUsdtPair != nil {
			path = append(path, usdtToken)
			path = append(path, toToken)
			return path, nil
		}
		return path, errors.New(fmt.Sprintf("%s -> %s router path error", swapFromToken, swapToToken))

	}

	if swapToTokenAddress == config.BscToken("wbnb") {
		if fromWbnbPair != nil {
			path = append(path, toToken)
			return path, nil
		}
		if fromBusdPair != nil {
			path = append(path, busdToken)
			path = append(path, toToken)
			return path, nil
		}
		if fromUsdtPair != nil {
			path = append(path, usdtToken)
			path = append(path, toToken)
			return path, nil
		}
		return path, errors.New(fmt.Sprintf("%s -> %s router path error", swapFromToken, swapToToken))

	}

	if fromWbnbPair != nil {
		path = append(path, bnbToken)
		if toWbnbPair != nil {
			path = append(path, toToken)
			return path, nil
		}
		if toBusdPair != nil {
			path = append(path, busdToken)
			path = append(path, toToken)
			return path, nil
		}
		if toUsdtPair != nil {
			path = append(path, usdtToken)
			path = append(path, toToken)
			return path, nil
		}

	}
	if fromBusdPair != nil {
		path = append(path, busdToken)
		if toBusdPair != nil {

			path = append(path, toToken)
			return path, nil
		}
		if toWbnbPair != nil {
			path = append(path, bnbToken)
			path = append(path, toToken)
			return path, nil
		}
		if toUsdtPair != nil {
			path = append(path, usdtToken)
			path = append(path, toToken)
			return path, nil
		}

	}

	if fromUsdtPair != nil {
		path = append(path, usdtToken)
		if toUsdtPair != nil {
			path = append(path, toToken)
			return path, nil
		}
		if toBusdPair != nil {
			path = append(path, busdToken)
			path = append(path, toToken)
			return path, nil
		}
		if toWbnbPair != nil {
			path = append(path, bnbToken)
			path = append(path, toToken)
			return path, nil
		}

	}
	_, address, _ := GetPoolPair(router, fromToken, toToken, client)
	if address != nil {
		path = append(path, toToken)
		return path, nil
	}

	return path, errors.New(fmt.Sprintf("%s -> %s  router path error", swapFromToken, swapToToken))

}
func GetPoolPair(router string, token0, token1 common.Address, client *ethclient.Client) (*uniswapV2Pair.UniswapV2Pair, *common.Address, error) {

	uni2factoryAddress := GetFactoryAddress(router, client)

	uni2factory, err := uniswapv2factory.NewUniswapv2factory(uni2factoryAddress, client)
	if err != nil {
		return nil, nil, err
	}
	callOpt := &bind.CallOpts{}
	pairAddress, err := uni2factory.GetPair(callOpt, token0, token1)
	if err != nil {
		return nil, nil, err
	}
	if pairAddress.String() == "0x0000000000000000000000000000000000000000" {
		return nil, nil, fmt.Errorf("%s-%s pool not exist", token0, token1)
	} else {
		//address := common.HexToAddress("0xdff86b408284dff30a7cad7688fedb465734501c")
		pair, err := uniswapV2Pair.NewUniswapV2Pair(pairAddress, client)
		return pair, &pairAddress, err

	}
}

func PriceFromSwap(router string, path []common.Address, amountIn *big.Int, client *ethclient.Client) float64 {

	uniswapv2router02 := GetUni2router2(router, client)

	amountsOut, err := uniswapv2router02.GetAmountsOut(&bind.CallOpts{}, amountIn, path)
	if err != nil {
		return 0
	}
	price := amountsOut[len(path)-1]

	priceFloat, err := strconv.ParseFloat(ToDecimal(price.String(), 18).String(), 64)
	if err != nil {
		panic(err)
	}
	return priceFloat
}
func GetAmountsOut(router string, path []common.Address, amountIn *big.Int, client *ethclient.Client) *big.Int {

	uniswapv2router02 := GetUni2router2(router, client)

	amountsOut, err := uniswapv2router02.GetAmountsOut(&bind.CallOpts{}, amountIn, path)
	if err != nil {
		log.Fatal(err)
	}
	price := amountsOut[len(path)-1]

	return price
}

func GetFactoryAddress(router string, client *ethclient.Client) common.Address {
	router2 := GetUni2router2(router, client)

	factory, err := router2.Factory(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	return factory
}

func GetUni2router2(router string, client *ethclient.Client) *uniswapv2router02.Uniswapv2router02 {
	routerAddress := common.HexToAddress(router)
	router2, err := uniswapv2router02.NewUniswapv2router02(routerAddress, client)
	if err != nil {
		panic(err)
	}
	return router2
}
