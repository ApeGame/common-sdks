package fake

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"time"
)

const baseURL = "https://common-service.mobus.workers.dev"

type UserInfo struct {
	Avatar   string
	Nickname string
	Bio      string
}

func GenerateUsrInfo(amount int) ([]*UserInfo, error) {
	users := make([]*UserInfo, 0, amount)
	if rand.Intn(10)%3 == 0 {
		return generateSourceTwo(amount)
	} else {
		for i := 0; i < amount; i++ {
			userInfo, err := generateSourceOne()
			if err != nil {
				log.Println("generate user error", err.Error())
				continue
			}
			users = append(users, userInfo)
		}
	}

	return users, nil
}

func generateSourceOne() (*UserInfo, error) {
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
	type resultSchema struct {
		ImageUrl string `json:"image_url"`
		Nickname string `json:"nickname"`
		Bio      string `json:"bio"`
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
		Bio:      r.Result.Bio,
	}, nil
}

func generateSourceTwo(amount int) ([]*UserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://randomuser.me/api/?results=%d", amount))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("could not close response body", "error", err.Error())
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type responseSchema struct {
		Results []struct {
			Gender string `json:"gender"`
			Name   struct {
				Title string `json:"title"`
				First string `json:"first"`
				Last  string `json:"last"`
			} `json:"name"`
			Location struct {
				Street struct {
					Number int    `json:"number"`
					Name   string `json:"name"`
				} `json:"street"`
				City        string `json:"city"`
				State       string `json:"state"`
				Country     string `json:"country"`
				Postcode    int    `json:"postcode"`
				Coordinates struct {
					Latitude  string `json:"latitude"`
					Longitude string `json:"longitude"`
				} `json:"coordinates"`
				Timezone struct {
					Offset      string `json:"offset"`
					Description string `json:"description"`
				} `json:"timezone"`
			} `json:"location"`
			Email string `json:"email"`
			Login struct {
				Uuid     string `json:"uuid"`
				Username string `json:"username"`
				Password string `json:"password"`
				Salt     string `json:"salt"`
				Md5      string `json:"md5"`
				Sha1     string `json:"sha1"`
				Sha256   string `json:"sha256"`
			} `json:"login"`
			Dob struct {
				Date time.Time `json:"date"`
				Age  int       `json:"age"`
			} `json:"dob"`
			Registered struct {
				Date time.Time `json:"date"`
				Age  int       `json:"age"`
			} `json:"registered"`
			Phone string `json:"phone"`
			Cell  string `json:"cell"`
			Id    struct {
				Name  string      `json:"name"`
				Value interface{} `json:"value"`
			} `json:"id"`
			Picture struct {
				Large     string `json:"large"`
				Medium    string `json:"medium"`
				Thumbnail string `json:"thumbnail"`
			} `json:"picture"`
			Nat string `json:"nat"`
		} `json:"results"`
		Info struct {
			Seed    string `json:"seed"`
			Results int    `json:"results"`
			Page    int    `json:"page"`
			Version string `json:"version"`
		} `json:"info"`
	}

	var rs responseSchema
	if err := json.Unmarshal(body, &rs); err != nil {
		return nil, err
	}

	users := make([]*UserInfo, 0, len(rs.Results))
	for _, r := range rs.Results {
		var nickname = r.Name.First + " " + r.Name.Last
		randNumber := rand.Intn(10)
		if randNumber%2 == 0 {
			nickname = r.Login.Username
		}

		var bio string
		switch randNumber % 6 {
		case 0:
			bio = fmt.Sprintf(`I'm from %s'`, r.Location.Country)
		case 1:
			bio = fmt.Sprintf(`I am living in %s,%s`, r.Location.City, r.Location.State)
		case 2:
			bio = fmt.Sprintf(`My birthday is %s'`, r.Dob.Date.Format("Jan _2 2006"))
		case 3:
			bio = fmt.Sprintf(`I'm %d years old'`, r.Dob.Age)
		case 4:
			bio = fmt.Sprintf(`I use %s time. Which is %s, What about you?`, r.Location.Timezone.Description, r.Location.Timezone.Offset)
		case 5:
			bio = fmt.Sprintf(`Send me a gift! My Address is No %d, Street %s, %s, %s. Postcode: %d`, r.Location.Street.Number, r.Location.Street.Name, r.Location.City, r.Location.State, r.Location.Postcode)
		}

		users = append(users, &UserInfo{
			Avatar:   r.Picture.Large,
			Nickname: nickname,
			Bio:      bio,
		})
	}

	return users, nil
}
