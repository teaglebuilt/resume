package generator

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func generateHTML(filePath string) {
	out, err := os.Create(filePath)
	defer out.Close()

	data := map[string]interface{}{}

	file, err := ioutil.ReadFile(("data.json"))
	check(err)

	json.Unmarshal(file, &data)
	check(err)

	t, err := template.ParseGlob("template/*")
	t.Execute(out, data)
}
