package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"regexp"
)

var ETHAddressRE, _ = regexp.Compile(`^0x\w{40}$`)

func CheckEthAddress(vars map[string]string) error {
	ethAddr, ok := vars["ethAddress"]
	if !ok {
		return fmt.Errorf("missing eth address")
	}

	if !ETHAddressRE.Match([]byte(ethAddr)) {
		return fmt.Errorf("address is not valid")
	}

	return nil
}

func ParseBody(r *http.Response, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return fmt.Errorf("error reading body: %s", err)
	}

	if err := json.Unmarshal(body, x); err != nil {
		return fmt.Errorf("error parsing body: %s", err)
	}

	return nil
}

// Following code has been adapted from https://github.com/ethereum/go-ethereum/issues/21221
func WeiToEth(w string) string {
	wei := ParseBigFloat(w)
	return fmt.Sprintf("%.18f", new(big.Float).Quo(big.NewFloat(wei), big.NewFloat(1e18)))
}

func ParseBigFloat(value string) float64 {
	n := new(big.Float)
	n.SetString(value)
	ret, _ := n.Float64()
	return ret
}
