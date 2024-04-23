base='https://raw.githubusercontent.com/tonkeeper/ton-assets/main/'
curl -L -o data/jettons.json $base/jettons.json
echo "" >> data/jettons.json
curl -L -o data/accounts.json $base/accounts.json
echo "" >> data/accounts.json