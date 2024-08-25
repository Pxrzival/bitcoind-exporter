package fetcher

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Primexz/bitcoind-exporter/config"
	"github.com/ybbus/jsonrpc/v3"
)

// Create RPC client with cookie-based authentication
var rpcClient = jsonrpc.NewClientWithOpts(computeAddress(), &jsonrpc.RPCClientOpts{
	CustomHeaders: map[string]string{
		"Authorization": "Basic " + computeCookieAuth(),
	},
})

// Reads the RPC cookie for authentication
func computeCookieAuth() string {
	cookieFilePath := config.C.RPCCookieFile
	if cookieFilePath == "" {
		panic("RPCCookieFile path is not set in the configuration")
	}

	cookie, err := ioutil.ReadFile(cookieFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read RPC cookie file from '%s': %v", cookieFilePath, err))
	}
	// The cookie file typically contains the format: <username>:<password>
	return strings.TrimSpace(string(cookie))
}


// Computes the RPC server address
func computeAddress() string {
	address := config.C.RPCAddress

	if strings.HasPrefix(address, "http://") {
		return address
	} else {
		return fmt.Sprintf("http://%s", address)
	}
}
