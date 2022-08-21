package yamlparser

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type FileDatas struct {
	FileName string
	Content  map[interface{}]interface{}
}

func filePath() {
	root := "./inventory"
	fileSystem := os.DirFS(root)
	var pathSlice []string

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
			//fmt.Println(path)
			pathSlice = append(pathSlice, path)
		}
		return nil
	})
}

func parseYaml(fileName string) map[interface{}]interface{} {

	fileObj, err := ioutil.ReadFile("./" + fileName)

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})

	err2 := yaml.Unmarshal(fileObj, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	return data
}

func Execute() {
	var fileName = "config.yaml"
	fileData := parseYaml(fileName)
	fmt.Println(fileData)
	filePath()
}

func init() {

}
