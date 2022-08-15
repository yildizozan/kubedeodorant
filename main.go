package main

import (
	"kubed/cmd"
	"kubed/yamlparser"
	//"kubed/yamlparser"
)

func main() {
	cmd.Execute()
	yamlparser.DirectoryList()
	//yamlparser.YamlParser()
}
