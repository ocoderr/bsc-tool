# Bsc Tool 
A  Bsc Tool  Application 

## Building on Linux

Dependencies:

   * go >= 1.9
   
Clone & compile:
    
    git clone https://github.com/ocoderr/bsc-tool.git
    go mod tidy
    go build
    ./bsc-tool

    # if u want build app follow this
    # go get fyne.io/fyne/cmd/fyne
    # fyne package -os darwin -app-id "com.openwallet.bsc.tool"
    # fyne package -os windows

## Features
Dex(PancakeSwap)
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

