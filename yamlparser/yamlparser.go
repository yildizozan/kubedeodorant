package yamlparser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"gopkg.in/yaml.v3"
)

type FileDatas struct {
	FileName string
	Content  []string
}


func readLineByLine(filePath string){
file, err := os.Open(filePath)
if err != nil {
	log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
// optionally, resize scanner's capacity for lines over 64K, see next example
for scanner.Scan() {
	fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
	log.Fatal(err)
}
}

func parseYaml(fileName string) map[string]string {

	fileObj, err := ioutil.ReadFile("./" + fileName)

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]string)

	err2 := yaml.Unmarshal(fileObj, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	return data
}

func Execute() {

	readLineByLine("config.yaml")

}

func init() {

}
