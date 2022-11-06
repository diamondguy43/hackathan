package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetTransaction(walletAddr string) {

	url := fmt.Sprintf("https://api.verbwire.com/v1/nft/data/transactions?walletAddress=%s", walletAddr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-Key", os.Getenv("VW"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(res)
	fmt.Println(string(body))

}
