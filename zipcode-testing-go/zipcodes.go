package zipcodes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

var ZIPcodeAPIKey = os.Getenv("ZIPCODE_API_KEY")

func ZipcodeIsValid(input string) (bool, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.zipcodeapi.com/rest/%s/info.json/%s/degrees", ZIPcodeAPIKey, input), nil)
	if err != nil {
		log.Print("Error building request")
		return false, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error sending request")
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, errors.New("error looking up zipcode, please double check your entry")
	} else {

		var respJSON APIResponse
		json.NewDecoder(resp.Body).Decode(&respJSON)

		return true, nil
	}
}
