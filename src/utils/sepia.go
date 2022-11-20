package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
)

func ChangeSepia(inDir string, outDir string) {
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), "png") {
			image, err := imgio.Open(inDir + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
				return
			}
			sepia := effect.Sepia(image)
			if err := imgio.Save(outDir+"/sepia"+f.Name(), sepia, imgio.PNGEncoder()); err != nil {
				log.Fatal(err)
				return
			}
		}
	}

	fmt.Println("Sepia Converted")
}
