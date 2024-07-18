package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	//fmt.Printf("viper")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
func initSlog() {

	slog.SetLogLoggerLevel(slog.LevelDebug)
	//change the slog output to file

	// Open a file for logging
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(logFile)
	//change the slog output tconsole
	//slog.SetOutputConsole(true)

}

// loadLinksFromFile lê um ficheiro de texto que contém uma lista de links e retorna uma slice de strings com esses links
func loadLinksFromFile(filePath string) ([]string, error) {

	slog.Debug("---- inicio loadLinksFromFile ----")

	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("erro ao abrir o ficheiro: %v", err)
		return nil, fmt.Errorf("erro ao abrir o ficheiro: %v", err)
	}
	defer file.Close()

	var links []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		slog.Debug("linha do ficheiro: " + line)
		if line != "" {
			links = append(links, line)
			slog.Debug("link a ser carregado: " + line)
		}
	}

	if err := scanner.Err(); err != nil {
		slog.Error("erro ao ler o ficheiro: %v", err)
		return nil, fmt.Errorf("erro ao ler o ficheiro: %v", err)
	}

	slog.Debug("---- fim loadLinksFromFile ----")
	return links, nil
}

// downloadFile faz o download do conteúdo de um link e salva-o num ficheiro local
func downloadFile(url, folderPath string) error {

	slog.Debug("---- inicio downloadFile ----")

	// Obter o nome do ficheiro a partir da URL
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1]
	filePath := filepath.Join(folderPath, fileName)
	slog.Debug("filePath: " + filePath)

	// Fazer o pedido HTTP
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("erro ao fazer o pedido HTTP: %v", err)
		return fmt.Errorf("erro ao fazer o pedido HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Criar o ficheiro local
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("erro ao criar o ficheiro: %v", err)
	}
	defer outFile.Close()

	// Copiar o conteúdo da resposta HTTP para o ficheiro local
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao copiar o conteúdo: %v", err)
	}

	slog.Debug("---- fim downloadFile ----")
	return nil
}

// downloadLinks faz o download de todos os links para uma pasta especificada
func downloadLinks(links []string, folderPath string) error {
	for _, link := range links {
		//fmt.Printf("A fazer download de %s...\n", link)
		err := downloadFile(link, folderPath)
		if err != nil {
			fmt.Printf("Erro ao fazer download de %s: %v\n", link, err)
		} else {
			//fmt.Printf("Download de %s concluído.\n", link)
		}
	}
	return nil
}
func main() {

	initSlog()
	fmt.Println("Wellcame to Mass Link Downloder")
	slog.Info("---------------------------------")
	slog.Info("Program Started")
	slog.Info("---------------------------------")

	initConfig()

	slog.Debug("Config file loaded")

	//define the o ficheiro de links
	linkListFile := filepath.Join(viper.GetString("linkListFileFolder"), viper.GetString("linkListFile"))
	slog.Debug("linkListFile: " + linkListFile)

	//define the download folder
	downloadFolder := viper.GetString("downloadFolder")
	slog.Debug("downloadFolder: " + downloadFolder)

	slog.Debug("config data loaded")

	links, err := loadLinksFromFile(linkListFile)
	if err != nil {
		slog.Error("Erro: %v\n", err)
		fmt.Printf("Erro: %v\n", err)
		return
	}

	//fmt.Println("3")
	//#DEBUG
	//fmt.Println("Loaded links:")
	//for _, link := range links {
	//	fmt.Println(link)
	//

	// Criar a pasta de downloads se não existir
	if _, err := os.Stat(downloadFolder); os.IsNotExist(err) {
		os.Mkdir(downloadFolder, 0755)
	}

	// Fazer o download de todos os links
	err = downloadLinks(links, downloadFolder)
	if err != nil {
		fmt.Printf("Erro ao fazer o download dos links: %v\n", err)
	} else {
		fmt.Println("Todos os downloads foram concluídos.")
	}

	//close logFile
	slog.Info("---------------------------------")
	slog.Info("- 3 - %d", "Program Ended")
	slog.Info("---------------------------------")

}
