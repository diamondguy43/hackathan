package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func MintNFT() {

	url := "https://api.verbwire.com/v1/nft/mint/customContractMintFromMetadataUrl"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "multipart/form-data")
	req.Header.Add("X-API-Key", os.Getenv("VW"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
