package zipcodes

type APIResponse struct {
	Zipcode  string `json:"zip_code"`
	City     string `json:"city"`
	State    string `json:"state"`
	Timezone struct {
		Abbreviation string `json:"timezone_abbr"`
	} `json:"timezone"`
}
