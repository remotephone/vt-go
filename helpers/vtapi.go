package helpers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type vtresponse struct {

}

func Check_hash(hash, vtkey string) string { 

	url := "https://www.virustotal.com/api/v3/files/" + hash

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-apikey", vtkey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	return string(body)

}

