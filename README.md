## About me

This repository provides the data of popular Jettons and Accounts on the TON blockchain. 
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
_Token symbols are case-insensitive_

### Get Jetton by address
_You can input both a friendly address and a raw address_
```go
jetton := assets.GetJettonByAddress("EQATcUc69sGSCCMSadsVUKdGwM1BMKS-HKCWGPk60xZGgwsK")
if jetton == nil {
    fmt.Println("Token is not whitelisted")
} else {
    fmt.Println("Token symbol: " + jetton.Symbol)	
}
```

### Get account by address
```go
account := assets.GetAccountByAddress("EQBfAN7LfaUYgXZNw5Wc7GBgkEX2yhuJ5ka95J1JJwXXf4a8")
if account == nil {
    fmt.Println("Account is not whitelisted")
} else {
fmt.Println("Account name: " + account.Name + " - Require memo: "+ fmt.Sprint(account.RequireMemo) + " - Workchain: " + fmt.Sprint(account.WorkChain))	
}
```
_You can input both a friendly address and a raw address_