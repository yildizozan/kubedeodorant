package yamlparser

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type FileDatas struct {
	FileName string
	Content  map[interface{}]interface{}
}

func filePath(folderName string) *list.List {
	fileList := list.New()
	err := filepath.Walk(folderName,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
				fileList.PushFront(path)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return fileList
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
	fileNames := filePath("inventory")
	fmt.Println(fileNames)
	for i := 0; i < 10; i++ {
		fmt.Println(fileName[i])
	}
}

func init() {

}
