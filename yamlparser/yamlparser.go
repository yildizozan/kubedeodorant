package yamlparser

import (
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"gopkg.in/yaml.v3"
	"strings"
)

type Line struct {
	IsComment bool
	Key string
	Value string
	Content string
}

func parseLine(line string) Line{
	parts := strings.Split(line, ":")
	var temp Line

	if (len(parts) == 2 ){
	temp.Key = parts[0]
	temp.Value = parts[1]
	temp.IsComment = false
	} else if (len(parts) == 3) {
	temp.Key = parts[1]
	temp.Value = parts[2]
	temp.IsComment = true
	} else{
		temp.Content = line
	}
	
	return temp
}


func ReadLineByLine(filePath string) []Line{
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
}

func init() {

}
