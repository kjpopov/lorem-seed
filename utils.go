package main

import (
	"encoding/json"
	"fmt"
)

func printJSON(data []byte) {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
	prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(prettyJSON))
}
