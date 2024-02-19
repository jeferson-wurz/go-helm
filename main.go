package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	// Carregar o arquivo de template
	templateFile := "templates/values.yaml.template"
	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	// Parse o conteúdo do arquivo de template
	tmpl, err := template.New("values").Parse(string(templateContent))
	if err != nil {
		panic(err)
	}

	// Carregar os placeholders do arquivo JSON
	placeholdersFile := "placeholders/placeholders.json"
	placeholderData, err := ioutil.ReadFile(placeholdersFile)
	if err != nil {
		panic(err)
	}

	// Decodificar o conteúdo JSON
	var placeholders map[string]interface{}
	err = json.Unmarshal(placeholderData, &placeholders)
	if err != nil {
		panic(err)
	}

	// Criar o arquivo de saída
	outputFile := "values.yaml"
	f, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Preencher o template com os valores dos placeholders
	err = tmpl.Execute(f, placeholders)
	if err != nil {
		panic(err)
	}

	println("Arquivo de valores gerado com sucesso:", outputFile)
}
