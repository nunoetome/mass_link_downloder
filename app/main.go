package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var SYS_COMFIG_LINK_LIST = []string{}
var SYS_COMFIG_DESTINATION_FOLDER string

// download
func download(link2Download string) {

}

// loadLinksFromFile lê um ficheiro de texto que contém uma lista de links e retorna uma slice de strings com esses links
func loadLinksFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o ficheiro: %v", err)
	}
	defer file.Close()

	var links []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			links = append(links, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erro ao ler o ficheiro: %v", err)
	}

	return links, nil
}
func main() {

	//var linkListFile string = ""
	var linkListFileName string = "linkListFileName.txt"
	var linkListFileFolder string = "C:\\Users\\nunot\\Documents\\_Documentos\\Projectos\\mass link downloder\\repositorio\\app"
	linkListFile := linkListFileFolder + "\\" + linkListFileName

	//config file defenition
	/*var configFileFileName string = "configFileFileName"
	var configFileFileFolder string = "C:\\Users\\nunot\\Documents\\_Documentos\\Projectos\\mass link downloder\\repositorio\\app"
	configFile := configFileFileName + "\\" + configFileFileFolder
	*/

	links, err := loadLinksFromFile(linkListFile)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println("Loaded links:")
	for _, link := range links {
		fmt.Println(link)
	}
}

/*
func main() {

	//var linkListFile string = ""
	var linkListFileName string = "linkListFileName.txt"
	var linkListFileFolder string = "C:\\Users\\nunot\\Documents\\_Documentos\\Projectos\\mass link downloder\\repositorio\\app"
	linkListFile := linkListFileFolder + "\\" + linkListFileFolder

	//var configFile string = ""
	var configFileFileName string = "configFileFileName"
	var configFileFileFolder string = "C:\\Users\\nunot\\Documents\\_Documentos\\Projectos\\mass link downloder\\repositorio\\app"
	configFile := configFileFileName + "\\" + configFileFileFolder

	fmt.Println(linkListFile)
	fmt.Println(linkListFileName)
	fmt.Println(linkListFileFolder)
	fmt.Println(linkList)

	fmt.Println(configFile)
	fmt.Println(configFileFileName)
	fmt.Println(configFileFileFolder)

	load_link_list_from_file(linkListFile)

}
*/
