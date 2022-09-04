package yamlparser

import (
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"gopkg.in/yaml.v3"
	"strings"
	"fmt"
    "io/fs"
)

type Line struct {
	IsComment bool
	Key string
	Value string
	Content string
}

func parseLine(line string) Line{
	parts := strings.Split(line, " ")
	var temp Line

	if (len(parts) == 2 ){
	twoparts := strings.Split(line, ":")
	temp.Key = twoparts[0]
	temp.Value = twoparts[1]
	temp.IsComment = false
	} else if (len(parts) == 3) {
	temp.Key = parts[1]
	temp.Key = temp.Key[:len(temp.Key)-1]
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

func FilePath(fileName string) []string {
    root := fileName
    fileSystem := os.DirFS(root)
    var pathSlice []string

    fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            log.Fatal(err)
        }
        if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
            fmt.Println(path)
            pathSlice = append(pathSlice, path)
        }
        return nil
    })
	return pathSlice
}

func Execute() {
}

func init() {

}
