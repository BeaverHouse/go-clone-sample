package main

import (
	"encoding/json"
	"os"
	"sample/sepia/src/utils"
)

func main() {

	jsonString := os.Args[1]

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)

	nasPath := jsonMap["nasPath"].(string)
	inputDir := jsonMap["inputDir"].(string)
	outputDir := jsonMap["outputDir"].(string)

	if err := os.MkdirAll(nasPath+outputDir, os.ModePerm); err != nil {
		panic(err)
	}

	utils.ChangeSepia(nasPath+inputDir, nasPath+outputDir)
}
