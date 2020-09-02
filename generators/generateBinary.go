package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fileName := "dirFile.txt"
	content, _ := ioutil.ReadFile(fileName)
	out, _ := os.Create("generated.go")
	out.Write([]byte("package main\n"))
	out.Write([]byte("var dirFile = map[string][]byte{\n"))
	out.Write([]byte("\"" + fileName + "\": []byte(`"))
	out.Write(content)
	out.Write([]byte("`),\n};"))

	fileName = "profFile.txt"
	content, _ = ioutil.ReadFile(fileName)
	out.Write([]byte("var profFile = map[string][]byte{\n"))
	out.Write([]byte("\"" + fileName + "\": []byte(`"))
	out.Write(content)
	out.Write([]byte("`),\n};"))

	fileName = "appsFile.txt"
	content, _ = ioutil.ReadFile(fileName)
	out.Write([]byte("var appsFile = map[string][]byte{\n"))
	out.Write([]byte("\"" + fileName + "\": []byte(`"))
	out.Write(content)
	out.Write([]byte("`),\n};"))

	fmt.Println("Generated Embed dir, prof, apps!")

}
