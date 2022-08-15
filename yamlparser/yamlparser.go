package yamlparser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Pair struct {
	key   string
	value string
}

type MyFile struct {
	fileName string
	content  map[int]string
}

func DirectoryList() map[int]string {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	directories := make(map[int]string)

	for i, file := range files {
		//fmt.Println(file.Name())
		directories[i] = file.Name()
		fmt.Printf(" " + directories[i])
	}

	return directories
}

func YamlParser() {

	yfile, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[int]string)
	fileNames := DirectoryList()

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	for i := range data {
		fmt.Println(i, err)
		err := MyFile{
			fileName: fileNames[i],
			content:  data,
		}
		fmt.Printf("%v", err)
	}
}

func Execute() {

}

func init() {

}
