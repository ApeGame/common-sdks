package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://common-service.mobus.workers.dev"

func GetChainIdByName(chainName string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/chain/id?chain_name=%s", baseURL, chainName))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type resultSchema struct {
		ChainId string `json:"chain_id"`
	}
	type respSchema struct {
		Result resultSchema `json:"result"`
	}
	var r respSchema
	if err := json.Unmarshal(body, &r); err != nil {
		return "", err
	}
	return r.Result.ChainId, nil
}

func GetChainNameById(chainId string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/chain/name?chain_id=%s", baseURL, chainId))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type resultSchema struct {
		ChainName string `json:"chain_name"`
	}
	type respSchema struct {
		Result resultSchema `json:"result"`
	}
	var r respSchema
	if err := json.Unmarshal(body, &r); err != nil {
		return "", err
	}
	return r.Result.ChainName, nil
}
