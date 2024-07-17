package main

import (
	"fmt"
)

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
	//fmt.Println(linkList)

	fmt.Println(configFile)
	fmt.Println(configFileFileName)
	fmt.Println(configFileFileFolder)

}
