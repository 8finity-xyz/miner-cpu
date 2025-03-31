# How to run

1. Install go - https://go.dev
2. Install dependencies
```sh
go install
```
3. Create .env file and fill values (you can use example `cp .env.example .env`)

Use [public nodes](https://chainlist.org/chain/146) or setup private 
```sh
# example of .env
INFINITY_RPC=https://rpc.blaze.soniclabs.com
INFINITY_WS=wss://rpc.blaze.soniclabs.com

# Private key without 0x (64 symobls). It should have some $S for transactions
INFINITY_PRIVATE_KEY=0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
```

4. Run miner
```
go run main.go
```