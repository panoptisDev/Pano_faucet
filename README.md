# Pano_faucet
Faucet

The faucet is a web application with the goal of distributing small amounts of Ether in private and test networks.

Get started

Run with Docker

Use private key as sender

docker run -d -p 8080:8080 -e WEB3_PROVIDER="rpc endpoint" -e PRIVATE_KEY="hex private key" chainflag/pls-faucet:1.0.0 -httpport 8080

Use keystore file as sender

docker run -d -p 8080:8080 -e WEB3_PROVIDER="rpc endpoint" -e KEYSTORE="keystore path" -v `pwd`/keystore:/app/keystore -v `pwd`/password.txt:/app/password.txt chainflag/pano-faucet:1.0.0 -httpport 8080
Optional Flags



Flag
Description
Default Value




-chainname
Network name to display on the frontend
testnet


-httpport
Listener port to serve HTTP connection
8080


-interval
Number of minutes to wait between funding rounds
1440


-payout
Number of Ethers to transfer per user request
1


-proxycount
Count of reverse proxies in front of the server
0


-queuecap
Maximum transactions waiting to be sent
100




Deploy to Heroku
heroku create
heroku buildpacks:add heroku/nodejs
heroku buildpacks:add heroku/go
heroku config:set WEB3_PROVIDER=rpc endpoint
heroku config:set PRIVATE_KEY=hex private key
git push heroku main
heroku open
or


Development

Prerequisites

Go (1.16 or later)
Node.js


Build

Clone the repository and navigate to the app’s directory

git clone https://gitlab.com/pulsechaincom/pano-faucet.git
cd pano-faucet

Bundle Frontend web with Rollup

npm run build
For more details, please refer to the web readme

Build binary application to run faucet

go build
export WEB3_PROVIDER=https://ropsten.infura.io
export PRIVATE_KEY=secret
./pano-faucet

License
This project is licensed under the MIT License
