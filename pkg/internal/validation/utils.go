package validation

import (
	"fmt"
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
