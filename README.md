## About me

This repository provides the data of popular Jettons and NFT Collections on the TON blockchain. 
The data is pulled automatically from [tonkeeper/ton-assets](https://github.com/tonkeeper/ton-assets) (whitelist address on Tonkeeper Wallet).


## Install
```shell
go get github.com/huantt/ton-assets-go
```

## How to use

### Get Jetton by symbol
```go
jetton := assets.GetJettonBySymbol("fish")
if jetton == nil {
    fmt.Println("Token is not whitelisted")
} else {
    fmt.Println("Token address: " + jetton.Address)	
}
```

### Get Jetton by address

```go
jetton := assets.GetJettonByAddress("EQATcUc69sGSCCMSadsVUKdGwM1BMKS-HKCWGPk60xZGgwsK")
if jetton == nil {
    fmt.Println("Token is not whitelisted")
} else {
    fmt.Println("Token symbol: " + jetton.Symbol)	
}
```