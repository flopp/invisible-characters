package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Character struct {
	Code        string
	Char        string
	Name        string
	Description string
	Url         string
}

func loadCharacters(fileName string) ([]*Character, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	jsonBlob, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var characters []Character
	err = json.Unmarshal(jsonBlob, &characters)
	if err != nil {
		return nil, err
	}

	var result []*Character
	for i := range characters {
		c := &characters[i]

		code, err := strconv.ParseUint(c.Code, 16, 32)
		if err != nil {
			return nil, err
		}
		c.Char = string(rune(int(code)))

		c.Url = fmt.Sprintf("%s-%s.html", c.Code, strings.Replace(c.Name, " ", "-", -1))

		result = append(result, c)
	}

	return result, nil
}

type Data struct {
	Characters []*Character
}

type Context struct {
	OutDir string
	Data   Data
	Urls   []string
}

func createCharacterFile(context *Context, c *Character, t *template.Template) {
	fileName := fmt.Sprintf("%s/%s", context.OutDir, c.Url)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := t.ExecuteTemplate(f, "base", c); err != nil {
		panic(err)
	}

	context.Urls = append(context.Urls, c.Url)
}

func createIndexFile(context *Context, t *template.Template) {
	fileName := fmt.Sprintf("%s/index.html", context.OutDir)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := t.ExecuteTemplate(f, "base", context.Data); err != nil {
		panic(err)
	}

	context.Urls = append(context.Urls, "index.html")
}

func createOtherFile(context *Context, url string) {
	t := template.Must(template.ParseFiles("templates/base.html", fmt.Sprintf("templates/%s", url)))

	fileName := fmt.Sprintf("%s/%s", context.OutDir, url)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := t.ExecuteTemplate(f, "base", context.Data); err != nil {
		panic(err)
	}

	context.Urls = append(context.Urls, url)
}

func createSitemapFile(context *Context) {
	fileName := fmt.Sprintf("%s/sitemap.txt", context.OutDir)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, url := range context.Urls {
		if _, err := f.WriteString(url); err != nil {
			panic(err)
		}
		if _, err := f.WriteString("\n"); err != nil {
			panic(err)
		}
	}
}

func main() {
	outDir := ".out"
	charactersFile := "characters.json"

	characters, err := loadCharacters(charactersFile)
	if err != nil {
		panic(err)
	}
	context := Context{outDir, Data{characters}, nil}

	tCharacter := template.Must(template.ParseFiles("templates/base.html", "templates/character.html"))
	tIndex := template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))

	if err := os.RemoveAll(outDir); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		panic(err)
	}

	createIndexFile(&context, tIndex)
	for _, c := range context.Data.Characters {
		createCharacterFile(&context, c, tCharacter)
	}
	createOtherFile(&context, "legal.html")
	createOtherFile(&context, "view.html")
	createOtherFile(&context, "empty-tweet.html")
	createOtherFile(&context, "empty-whatsapp.html")
	createOtherFile(&context, "404.html")

	createSitemapFile(&context)
}
