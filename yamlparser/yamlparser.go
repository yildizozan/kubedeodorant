package yamlparser

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v3"
)

type FileDatas struct {
	FileName string
	Content   map[interface{}]interface{}
}

func parseYaml(fileName string) map[interface{}]interface{}{

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

}

func init() {

}
