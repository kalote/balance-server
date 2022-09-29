# Balance-server

A server that returns the ETH balance for a specific EOA

## Install

```bash
git clone https://github.com/kalote/balance-server
cd balance-server
cp .env.example .env # don't forget to update the .env file with your infura APIKEY
make dev-jq # if you have jq installed
make dev # if you don't have jq
```

You can also build & run the docker image:

```bash
docker build -t balance-server:latest .
docker run --name=balance-server -p 3000:3000 -e APIKEY=YOUR_API_KEY -d balance-server:latest
```