package longpoll

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"time"
)

var (
	longPollData LongPollData
)

type Response struct {
	Error struct {
		Code    int    `json:"error_code"`
		Message string `json:"error_msg"`
	} `json:"error"`

	LongPollData `json:"response"`
}

type LongPollData struct {
	Server string `json:"server"`
	Key    string `json:"key"`
}

func GetCredentials(apiUrl *string, groupToken *string) (err error) {
	var endpoint string = *apiUrl + "messages.getLongPollServer?access_token=" + *groupToken

	httpResponse, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	parsedResponse := Response{}
	if err := json.Unmarshal(responseBody, &parsedResponse); err != nil {
		return err
	}

	if parsedResponse.Error.Code != 0 {
		return errors.New("VK error: " + parsedResponse.Error.Message)
	}

	longPollData = parsedResponse.LongPollData

	return nil
}

func Pull() {
	var err error
	if longPollData.Server == "" {
		err = errors.New("Trying to longPoll with empty credentials")
		panic(err)
	}

	time.Sleep(time.Millisecond * 2500)
	Pull()

	panic(errors.New("Test panic"))
}
