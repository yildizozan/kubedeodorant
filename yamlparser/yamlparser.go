package yamlparser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"gopkg.in/yaml.v3"
	"strings"
)

type Line struct {
	isComment bool
	key string
	value string

}

func parseLine(line string) Line{
	parts := strings.Split(line, " ")
	var temp Line

	if (len(parts) == 2 ){
	temp.key = parts[0]
	temp.value = parts[1]
	temp.isComment = false
	} else if (len(parts) == 3) {
	temp.key = parts[0]
	temp.value = parts[1]
	temp.isComment = false
	} 
	
	return temp
}


func readLineByLine(filePath string) []Line{
var lines []Line

file, err := os.Open(filePath)
if err != nil {
	log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
	lines = append(lines, parseLine(scanner.Text()))
}

if err := scanner.Err(); err != nil {
	log.Fatal(err)
}
return lines
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

	allLines := readLineByLine("config.yaml")
	fmt.Println(allLines)
}

func init() {

}
