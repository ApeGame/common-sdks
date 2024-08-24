package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseURL = "http://127.0.0.1:8080"

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
	type respSchema struct {
		ChainId string `json:"chain_id"`
	}
	var r respSchema
	if err := json.Unmarshal(body, &r); err != nil {
		return "", err
	}
	return r.ChainId, nil
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
	type respSchema struct {
		ChainName string `json:"chain_name"`
	}
	var r respSchema
	if err := json.Unmarshal(body, &r); err != nil {
		return "", err
	}
	return r.ChainName, nil
}
