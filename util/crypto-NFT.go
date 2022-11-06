package util

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetTransaction() {

	url := "https://api.verbwire.com/v1/nft/data/transactions?walletAddress"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}