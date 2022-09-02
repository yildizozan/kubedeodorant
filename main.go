package main

import (
	"kubed/cmd"
	"kubed/yamlparser"
	"log"
	"os"
)

func OverrideYaml(configPath string){
	configYaml := yamlparser.ReadLineByLine(configPath)
	testYaml := yamlparser.ReadLineByLine("inventory/test.yaml")
	var tempYaml []yamlparser.Line

	for _, inv  := range testYaml {
		
		var tempLine yamlparser.Line
		tempLine.Key = inv.Key
		tempLine.Value = inv.Value
		tempLine.IsComment = false
		tempLine.Content = inv.Content

		for _, conf := range configYaml {
			if conf.Key == inv.Key {
				tempLine.Value = conf.Value
				if(conf.IsComment){
					tempLine.IsComment = true
				}
			}
		}
		tempYaml = append(tempYaml, tempLine)

	}

	printYaml(tempYaml, "inventory/test1.yaml")
}

func printYaml(fileData []yamlparser.Line, filePath string){
	f, err := os.Create(filePath)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	for _, conf := range fileData {
		if len(conf.Key) > 0 {	
			_, err2 := f.WriteString(conf.Key + ":" + conf.Value + "\n")
			
			if err2 != nil {
				log.Fatal(err2)
			}
		}else if len(conf.Content) > 0{
			_, err2 := f.WriteString(conf.Content + "\n")
			
			if err2 != nil {
				log.Fatal(err2)
			}
		}else{
			_, err2 := f.WriteString("\n")
			
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}




}


func main() {
	cmd.Execute()
	OverrideYaml("config.yaml")
}
