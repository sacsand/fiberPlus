package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var exists = func(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func createFile(filePath, content string) (err error) {
	var f *os.File
	if f, err = os.Create(filePath); err != nil {
		return
	}

	defer func() { _ = f.Close() }()

	_, err = f.WriteString(content)

	return
}

func createDirectory(path string) {
	_ = os.Mkdir(path, 0700)
	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// 	os.Mkdir(path, 755)
	// }
}

type input struct {
	param       string
	path        string
	errorLine   string
	successLine string
}

type config struct {
	ModelPath      string `yaml:"modelpath"`
	PkgPath        string `yaml:"pkgpath"`
	ControllerPath string `yaml:"controllerpath"`
}

func (c *config) getConfig() *config {

	yamlFile, err := ioutil.ReadFile(".fiberplus.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// if yaml not exist use below values

	if len(c.ModelPath) < 1 {
		c.ModelPath = "models"
	}
	if len(c.PkgPath) < 1 {
		c.PkgPath = "pkg"
	}
	if len(c.ControllerPath) < 1 {
		c.ControllerPath = "controllers"
	}

	createDirectory(c.ModelPath)
	createDirectory(c.ControllerPath)
	createDirectory(c.PkgPath)

	return c
}
