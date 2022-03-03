package util

import (
	"github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/jsonrpc"
	"strconv"
	"strings"
)

const Wei = 1e9

func GetTransactionByHash(hash string, client *jsonrpc.Client) (output *web3.Transaction)  {
	client.Call("eth_getTransactionByHash", &output, hash)
	return
}
func GetBalance(address string, client *jsonrpc.Client) float64  {
	var output string
	err := client.Call("eth_getBalance", &output, address, "latest")
	if err != nil {
		return 0
	}
	return toEther(parseUint64orHex(output))
}

func parseUint64orHex(str string) (num int64) {
	base := 10
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
		base = 16
	}
	num, _ = strconv.ParseInt(str, base, 64)

	return
}

func toEther(wei int64) float64  {
	return float64(wei)/Wei
}