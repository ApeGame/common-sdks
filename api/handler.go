package api

import (
	"encoding/json"
	"github.com/ApeGame/common-sdks/config"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/chain/rpc", func(w http.ResponseWriter, r *http.Request) {
		chainId := r.URL.Query().Get("chain_id")
		addresses := getRpc(chainId)
		responseJson(w, map[string]interface{}{"rpc_addresses": addresses})
	})

	http.HandleFunc("/chain/name", func(w http.ResponseWriter, r *http.Request) {
		chainId := r.URL.Query().Get("chain_id")
		chanName := getChainNameById(chainId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		responseJson(w, map[string]interface{}{"chain_name": chanName})
	})

	http.HandleFunc("/chain/id", func(w http.ResponseWriter, r *http.Request) {
		chainName := r.URL.Query().Get("chain_name")
		chainId := getChainIdByName(chainName)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		responseJson(w, map[string]interface{}{"chain_id": chainId})
	})
}

var chainIdName = make(map[string]string)
var chainNameId = make(map[string]string)
var chainRpcAddress = make(map[string][]string)

func getChainIdByName(chainName string) string {
	if len(chainNameId) == 0 {
		c := config.GetConfig()
		for _, n := range c.Nodes {
			chainNameId[n.ChainName] = n.ChainId
		}
	}
	return chainNameId[chainName]
}
func getChainNameById(chainId string) string {
	if len(chainIdName) == 0 {
		c := config.GetConfig()
		for _, n := range c.Nodes {
			chainIdName[n.ChainId] = n.ChainName
		}
	}
	return chainIdName[chainId]
}

func getRpc(chainId string) []string {
	if len(chainRpcAddress) == 0 {
		c := config.GetConfig()
		for _, n := range c.Nodes {
			chainRpcAddress[n.ChainId] = n.RpcAddr
		}
	}
	return chainRpcAddress[chainId]
}

func responseJson(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
