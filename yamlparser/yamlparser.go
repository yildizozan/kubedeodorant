package yamlparser

import (
     "fmt"
     "io/ioutil"
     "log"
     "gopkg.in/yaml.v3"
)

func YamlParser() {
	
     yfile, err := ioutil.ReadFile("./yamlparser/test.yaml")

     if err != nil {

          log.Fatal(err)
     }

     data := make(map[interface{}]interface{})

     err2 := yaml.Unmarshal(yfile, &data)

     if err2 != nil {

          log.Fatal(err2)
     }

     for k, v := range data {

          fmt.Printf("\n%s ->\n\n %d\n\n", k, v)
     }
}

func Execute() {
     
}

func init() {

}
