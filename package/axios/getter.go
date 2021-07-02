package axios

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func Get(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func Post(address string, target interface{}, data url.Values) error {

	r, err := myClient.PostForm(address, data)

	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)

}
