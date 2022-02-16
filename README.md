# Pancakeswap Console
A Pancakeswap Application 

## Why
It is a pity that some areas do not have access to the functions of PancakeSwap.
I developed this tool to help these $People$.

I believe $People$ have the right to access the blockchain.

## Building on Linux

Dependencies:

   * go >= 1.9
   
Clone & compile:
    
    git clone https://github.com/ocoderr/bsc-tool.git
    go mod tidy
    go build
    ./pancakeswap-console 

    # if u want build app follow this
    # go get fyne.io/fyne/cmd/fyne
    # fyne package -os darwin -app-id "com.openwallet.pancakeswap"
    # fyne package -os windows

## Features
* Prices      [x]
* Syrup Pool           [x]
* Gui              [x]
* IFO                  [ ]
* Batch Transfer       [ ]
* Batch Deposit Syrup Pool [ ]


## Screen

![image](https://raw.githubusercontent.com/ocoderr/bsc-tool/master/image/1.png)

![image](https://raw.githubusercontent.com/ocoderr/bsc-tool/master/image/2.png)

![image](https://raw.githubusercontent.com/ocoderr/bsc-tool/master/image/3.png)

### Note
* The price may be a little different from that of the official website.Because using route path  <token> -> WBNB -> USDT
* Don't set up too many trading pairs, there will be problems with rpc node.
* Not tested in window environment


### Donations

BNB/CAKE/USDT: 0xec5fa25e37dfa8fa42210a94cbc8a61c7fd3751c

email: irmakgu40@gmail.com