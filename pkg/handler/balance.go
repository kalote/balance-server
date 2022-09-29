package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalote/balance-server/pkg/conf"
	"github.com/kalote/balance-server/pkg/internal/utils"
	"github.com/kalote/balance-server/pkg/service"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	// Verify parameters
	params := mux.Vars(r)
	if err := utils.CheckEthAddress(params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// call to infura returns unparsed data
	resp, err := service.GetBalance(
		fmt.Sprintf("%s/%s", conf.InfuraURL, h.c.InfuraKey),
		fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBalance","params": ["%v", "latest"],"id":1}`, params["ethAddress"]),
	)

	// if error during infura call return 500
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// parsing data & assigning to var
	bal := &InfuraResponse{}
	if err := utils.ParseBody(resp, bal); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// convert Infura data into human readable format & assign to return var
	balance := utils.WeiToEth(bal.Result)

	// return balance
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BalanceResponse{
		Balance: balance,
	})
}
