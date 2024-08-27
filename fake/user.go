package fake

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://common-service.mobus.workers.dev"

type UserInfo struct {
	Avatar   string
	Nickname string
}

func GenerateUsrInfo() (*UserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/fake/user", baseURL))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	type resultSchema struct {
		ImageUrl string `json:"image_url"`
		Nickname string `json:"nickname"`
	}
	type respSchema struct {
		Result resultSchema `json:"result"`
	}

	var r respSchema
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return &UserInfo{
		Avatar:   r.Result.ImageUrl,
		Nickname: r.Result.Nickname,
	}, nil
}
