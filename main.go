package main

import (
	"flag"
	"os"
	"sample/sepia/src/utils"
)

func main() {

	nasPath := flag.String("nas", "D:", "path to nas")
	inputDir := flag.String("in", "/input", "path to input directory")
	outputDir := flag.String("out", "/output", "path to output directory")

	flag.Parse()

	if err := os.MkdirAll(*nasPath+*outputDir, os.ModePerm); err != nil {
		panic(err)
	}

	utils.ChangeSepia(*nasPath+*inputDir, *nasPath+*outputDir)
}
