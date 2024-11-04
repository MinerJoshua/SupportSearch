package main

import (
"fmt"
"net/http"
//"errors"
"encoding/json"
"io"
)

func searchHandler(searchParams string) {


	searchParams = ""
	turnToJson()
	//body,err := turnToJson()

	/*
	if err == nil {
		return "",err
	}

	fmt.Println(body)
	*/
}


func turnToJson() (){

	url := "https://api-docs.20i.com/api/articles?fields=slug&fields=title&pagination[pageSize]=9999"
	bodyMap := map[string]interface{}{}
	docClient := &http.Client{}

	response,err := docClient.Get(url)

	if err != nil {
		//return "",errors.New("Could not access Search URL")
		return
	}


	readResponse,err := io.ReadAll(response.Body)

	if err != nil {
		//return bodyMap,err
		return 
	}

	defer response.Body.Close()

	if error := json.Unmarshal(readResponse,&bodyMap); error != nil {
		//return bodyMap,nil
		return
	}

	fmt.Println(bodyMap["data"])
	//return bodyMap,error
}

func main () {

	// Receive Post Request from Search and return list 
	searchHandler("Hello")


}

