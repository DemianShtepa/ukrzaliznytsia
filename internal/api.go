package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type trip struct {
	Id int `json:"id"`
}

type trips struct {
	Direct []trip `json:"direct"`
}

type Api struct {
	Env Env
}

func NewApi(env Env) *Api {
	return &Api{Env: env}
}

func (api Api) Fetch(events chan time.Time, sleepDuration time.Duration) {
	go func() {
		request, err := http.NewRequest("GET", api.Env.ApiURL, nil)
		if err != nil {
			fmt.Println("Error while creating request: ", err)
			return
		}

		request.Header.Add("Accept", "application/json")
		request.Header.Add("User-Agent", "UZ/1.10.1 iOS/17.1.1 User/1374232")
		request.Header.Add("Authorization", "Bearer "+api.Env.ApiToken)

		client := http.Client{}

		query := request.URL.Query()
		query.Add("station_from_id", strconv.Itoa(api.Env.StationFrom))
		query.Add("station_to_id", strconv.Itoa(api.Env.StationTo))
		query.Add("with_transfers", "0")

		for {
			for date := api.Env.DateFrom; date.Before(api.Env.DateTo); date = date.AddDate(0, 0, 1) {
				query.Add("date", date.Format("2006-01-02"))
				request.URL.RawQuery = query.Encode()
				response, err := client.Do(request)
				if err != nil {
					fmt.Println("Error while performing request: ", err)
					break
				}

				defer response.Body.Close()

				tripsResponse := trips{}
				if jsonDecoder := json.NewDecoder(response.Body); jsonDecoder.Decode(&tripsResponse) == nil {
					if len(tripsResponse.Direct) > 1 {
						events <- date
					}
				}

				time.Sleep(time.Second * 10)
			}

			time.Sleep(sleepDuration)
		}
	}()
}
