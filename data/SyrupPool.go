package data

import (
	"context"
	"fyne.io/fyne/v2/data/binding"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ocoderr/bsc-tool/abi/cakevault"
	"github.com/ocoderr/bsc-tool/abi/masterchef"
	"github.com/ocoderr/bsc-tool/abi/pancake"
	"github.com/ocoderr/bsc-tool/config"
	"github.com/ocoderr/bsc-tool/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"math"
	"math/big"
	"strings"
	"time"
)

type SyrupPool struct {
	Client      *ethclient.Client
	PoolName    string
	PoolAddress string

	//Total       float64
	Total      binding.String
	Reward     binding.String
	RewardUsdt binding.String
	Stake      binding.String
	Apy        binding.String
	Price      binding.String
	Estimation binding.String
}

func (tp *SyrupPool) Start() {

	go tp.refresh()

}
func (tp *SyrupPool) Update() {
	if strings.ToLower(tp.PoolName) == "cake-auto" {
		tp.CakeAutoSyrupPool()
	} else if strings.ToLower(tp.PoolName) == "cake" {
		tp.CakeSyrupPool()
	} else {
		tp.OtherSyrupPool()
	}
}
func (tp *SyrupPool) refresh() {

	for {
		select {
		case <-time.After(5 * time.Second):
			tp.Update()
		}
	}
}
func (tp *SyrupPool) IsActive() bool {
	if strings.ToLower(tp.PoolName) == "cake-auto" || strings.ToLower(tp.PoolName) == "cake" {
		return true
	} else {
		poolAddress := config.CF.SyrupPool[strings.ToLower(tp.PoolName)]
		toAddress := common.HexToAddress(poolAddress)
		pancakeStakePool, err := pancake.NewPancake(toAddress, tp.Client)
		if err != nil {
			return false
		}
		bonusEndBlock, err := pancakeStakePool.BonusEndBlock(utils.DefaultCallOpts)
		blockNumber, err := tp.Client.BlockNumber(context.Background())
		if bonusEndBlock.Int64() >= int64(blockNumber) {
			return true
		}
		return false
	}

}

func (tp *SyrupPool) CakeAutoSyrupPool() {
	erc20Token := utils.Erc20Token(config.CF.BscToken[strings.ToLower("cake")], tp.Client)
	decimals, _ := erc20Token.Decimals(utils.DefaultCallOpts)
	tp.PoolAddress = config.CF.SyrupPool[strings.ToLower("cake-auto")]

	toAddress := common.HexToAddress(tp.PoolAddress)
	CakeVault, err := cakevault.NewCakevault(toAddress, tp.Client)
	if err != nil {
		return
	}
	//path, err := utils.CalculatePath("cake", "USDT", config.CF.PancakeRouter, tp.Client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	price := GetTokenPrice("cake")

	totalAmountBig, err := CakeVault.BalanceOf(utils.DefaultCallOpts)
	if err != nil {
		return
	}
	totalAmount := utils.Str2Float(utils.ToDecimal(totalAmountBig.String(), 18).String())
	p := message.NewPrinter(language.English)

	tp.Total.Set(p.Sprintf("%.2f", totalAmount))
	tp.Price.Set(p.Sprintf("%.6f", price))

	Masterchef, err := masterchef.NewMasterchef(common.HexToAddress(config.CF.SyrupPool[strings.ToLower("cake")]), tp.Client)
	if err != nil {
		return
	}
	totalAmountBig1, err := erc20Token.BalanceOf(utils.DefaultCallOpts, common.HexToAddress(config.CF.SyrupPool[strings.ToLower("cake")]))
	if err != nil {
		return
	}
	totalAmount1 := utils.Str2Float(utils.ToDecimal(totalAmountBig1.String(), 18).String())

	totalAllocPoint, err := Masterchef.TotalAllocPoint(utils.DefaultCallOpts)
	if err != nil {
		return
	}
	cakePerBlock, err := Masterchef.CakePerBlock(utils.DefaultCallOpts)
	if err != nil {
		return
	}
	rewardPerBlock := utils.Str2Float(utils.ToDecimal(cakePerBlock.String(), int(18)).String())
	rewardPerDay := rewardPerBlock * 20 * 60 * 24

	poolInfo, err := Masterchef.PoolInfo(utils.DefaultCallOpts, big.NewInt(0))
	if err == nil {
		share := float64(poolInfo.AllocPoint.Int64()) / float64(totalAllocPoint.Int64())

		apr := rewardPerDay * share / totalAmount1
		pow := math.Pow(1+apr, 365) - 1

		tp.Apy.Set(p.Sprintf("%.2f%%", pow*100))

		apr = rewardPerBlock * 20 * 60 * share / totalAmount1
		pow = math.Pow(1+apr, 24) - 1
		f := 6400 * pow * price
		tp.Estimation.Set(p.Sprintf("%.2f", f))

	}

	if config.CF.FromAddress != "" {
		fromAddress := common.HexToAddress(config.CF.FromAddress)

		userInfo, err := CakeVault.UserInfo(utils.DefaultCallOpts, fromAddress)
		if err != nil {
			return
		}

		totalShares, err := CakeVault.TotalShares(utils.DefaultCallOpts)
		if err != nil {
			return
		}
		shares := userInfo.Shares

		mul := big.NewInt(0).Mul(totalAmountBig, shares)
		nowAmountBig := big.NewInt(0).Div(mul, totalShares)

		stakeAmountStr := utils.ToDecimal(userInfo.CakeAtLastUserAction.String(), int(decimals)).String()
		stakeFloat := utils.Str2Float(stakeAmountStr)

		nowAmountStr := utils.ToDecimal(nowAmountBig.String(), int(decimals)).String()
		nowAmount := utils.Str2Float(nowAmountStr)
		rewardFloat := nowAmount - stakeFloat

		tp.Reward.Set(p.Sprintf("%.2f", rewardFloat))
		tp.Stake.Set(p.Sprintf("%.2f", stakeFloat))
		tp.RewardUsdt.Set(p.Sprintf("%.2f", rewardFloat*price))
	} else {
		tp.Reward.Set(p.Sprintf("%.2f", 0.0))
		tp.Stake.Set(p.Sprintf("%.2f", 0.0))
		tp.RewardUsdt.Set(p.Sprintf("%.2f", 0.0))
	}

}

func (tp *SyrupPool) CakeSyrupPool() {
	erc20Token := utils.Erc20Token(config.CF.BscToken[strings.ToLower("cake")], tp.Client)
	decimals, _ := erc20Token.Decimals(utils.DefaultCallOpts)
	tp.PoolAddress = config.CF.SyrupPool[strings.ToLower("cake")]
	//abigen --abi=./abi/MasterChef.abi --pkg=masterchef  --out=./abi/masterchef/masterchef.go
	toAddress := common.HexToAddress(tp.PoolAddress)
	pancakeStakePool, err := masterchef.NewMasterchef(toAddress, tp.Client)
	if err != nil {
		return
	}

	price := GetTokenPrice("cake")

	totalAmountBig, err := erc20Token.BalanceOf(utils.DefaultCallOpts, common.HexToAddress(tp.PoolAddress))
	if err != nil {
		return
	}
	totalAmount := utils.Str2Float(utils.ToDecimal(totalAmountBig.String(), 18).String())
	p := message.NewPrinter(language.English)

	CakeAutoPoolAddress := config.CF.SyrupPool[strings.ToLower("cake-auto")]

	Cakevault, err := cakevault.NewCakevault(common.HexToAddress(CakeAutoPoolAddress), tp.Client)
	if err != nil {
		return
	}

	totalAutoAmountBig, err := Cakevault.BalanceOf(utils.DefaultCallOpts)
	if err != nil {
		return
	}
	totalAutoAmount := utils.Str2Float(utils.ToDecimal(totalAutoAmountBig.String(), 18).String())
	caktTotal := totalAmount - totalAutoAmount
	tp.Total.Set(p.Sprintf("%.2f", caktTotal))
	tp.Price.Set(p.Sprintf("%.6f", price))

	totalAllocPoint, err := pancakeStakePool.TotalAllocPoint(utils.DefaultCallOpts)
	if err != nil {
		return
	}
	cakePerBlock, err := pancakeStakePool.CakePerBlock(utils.DefaultCallOpts)
	if err != nil {
		return
	}
	rewardPerBlock := utils.Str2Float(utils.ToDecimal(cakePerBlock.String(), int(18)).String())
	rewardPerYear := rewardPerBlock * 20 * 60 * 24 * 356

	poolInfo, err := pancakeStakePool.PoolInfo(utils.DefaultCallOpts, big.NewInt(0))
	if err != nil {
		return
	}

	share := float64(poolInfo.AllocPoint.Int64()) / float64(totalAllocPoint.Int64())

	tp.Apy.Set(p.Sprintf("%.2f%%", rewardPerYear*share/totalAmount*100))

	f := 6400 * rewardPerBlock * 20 * 60 * 24 * share / totalAmount * price
	tp.Estimation.Set(p.Sprintf("%.2f", f))

	if config.CF.FromAddress != "" {
		fromAddress := common.HexToAddress(config.CF.FromAddress)
		userInfo, err := pancakeStakePool.UserInfo(utils.DefaultCallOpts, big.NewInt(0), fromAddress)
		if err != nil {
			return
		}

		stakeAmountStr := utils.ToDecimal(userInfo.Amount.String(), int(decimals)).String()
		stakeFloat := utils.Str2Float(stakeAmountStr)
		rewardFloatStr := utils.ToDecimal(userInfo.RewardDebt.String(), int(decimals)).String()
		rewardFloat := utils.Str2Float(rewardFloatStr)
		tp.Reward.Set(p.Sprintf("%.2f", rewardFloat))
		tp.Stake.Set(p.Sprintf("%.2f", stakeFloat))
		tp.RewardUsdt.Set(p.Sprintf("%.2f", rewardFloat*price))
	} else {
		tp.Reward.Set(p.Sprintf("%.2f", 0.0))
		tp.Stake.Set(p.Sprintf("%.2f", 0.0))
		tp.RewardUsdt.Set(p.Sprintf("%.2f", 0.0))
	}

}
func (tp *SyrupPool) OtherSyrupPool() {
	erc20Token := utils.Erc20Token(config.CF.BscToken[strings.ToLower(tp.PoolName)], tp.Client)
	decimals, _ := erc20Token.Decimals(utils.DefaultCallOpts)
	tp.PoolAddress = config.CF.SyrupPool[strings.ToLower(tp.PoolName)]

	toAddress := common.HexToAddress(tp.PoolAddress)
	pancakeStakePool, err := pancake.NewPancake(toAddress, tp.Client)
	if err != nil {
		return
	}

	poolPrice := GetTokenPrice(tp.PoolName)

	cakePrice := GetTokenPrice("CAKE")

	cakeErc20Token := utils.Erc20Token(config.CF.BscToken["cake"], tp.Client)
	totalAmountBig, err := cakeErc20Token.BalanceOf(utils.DefaultCallOpts, common.HexToAddress(tp.PoolAddress))
	if err != nil {
		return
	}

	totalAmount := utils.Str2Float(utils.ToDecimal(totalAmountBig.String(), 18).String())
	p := message.NewPrinter(language.English)

	tp.Total.Set(p.Sprintf("%.2f", totalAmount))
	tp.Price.Set(p.Sprintf("%.6f", poolPrice))

	rewardPerBlockBig, err := pancakeStakePool.RewardPerBlock(utils.DefaultCallOpts)
	if err != nil {
		return
	}
	rewardPerBlock := utils.Str2Float(utils.ToDecimal(rewardPerBlockBig.String(), int(decimals)).String())
	rewardPerYear := rewardPerBlock * 20 * 60 * 24 * 356
	rewardPerYearUsdt := rewardPerYear * poolPrice
	totalAmountUsdt := totalAmount * cakePrice

	tp.Apy.Set(p.Sprintf("%.2f%%", rewardPerYearUsdt/totalAmountUsdt*100))
	f := 6400 / totalAmount * rewardPerBlock * 20 * 60 * 24 * poolPrice
	tp.Estimation.Set(p.Sprintf("%.2f", f))

	if config.CF.FromAddress != "" {
		rewardAmountStr, stakeAmountStr := SyrupPoolCheckReward(config.CF.FromAddress, pancakeStakePool, decimals)
		rewardFloat := utils.Str2Float(rewardAmountStr)
		stakeFloat := utils.Str2Float(stakeAmountStr)
		tp.Reward.Set(p.Sprintf("%.2f", rewardFloat))
		tp.Stake.Set(p.Sprintf("%.2f", stakeFloat))
		tp.RewardUsdt.Set(p.Sprintf("%.2f", rewardFloat*poolPrice))
	} else {
		tp.Reward.Set(p.Sprintf("%.2f", 0.0))
		tp.Stake.Set(p.Sprintf("%.2f", 0.0))
		tp.RewardUsdt.Set(p.Sprintf("%.2f", 0.0))
	}
}

func SyrupPoolCheckReward(fromAddresss string, pancakeStakePool *pancake.Pancake, decimals uint8) (string, string) {
	fromAddress := common.HexToAddress(fromAddresss)

	pendingReward, err := pancakeStakePool.PendingReward(utils.DefaultCallOpts, fromAddress)

	if err != nil {
		log.Fatal(err)
	}
	userInfo, err := pancakeStakePool.UserInfo(utils.DefaultCallOpts, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	return utils.ToDecimal(pendingReward.String(), int(decimals)).String(), utils.ToDecimal(userInfo.Amount.String(), 18).String()

}
