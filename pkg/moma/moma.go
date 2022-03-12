package moma

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type OutputFile struct {
	Path   string
	Name   string
	Prefix string
}

type Flags struct {
	ApiKey     string
	OutputFile OutputFile
}

func HandleFlags() Flags {
	apiKey := flag.String("key", "", "Meraki API Key")
	path := flag.String("path", "./output", "output directory")
	name := flag.String("name", "", "output file name")
	prefix := flag.String("prefix", "", "output file name prefix")

	flag.Parse()

	if *apiKey == "" {
		*apiKey = os.Getenv("MERAKI_API_KEY")
	}
	if *apiKey == "" {
		fmt.Println("API Key not defined, Meraki API Key must either be passed as a runtime flag or an Environmental Variable")
		fmt.Println("   To pass as a flag include `-key <api key>`")
		fmt.Println("   To pass as an environmental variable set `MERAKI_API_KEY` to your API key")
		os.Exit(1)
	}

	if *name != "" && *prefix != "" {
		fmt.Printf("WARNING: both file name and prefix can not be used together, prefix will be ignored")
	}

	flags := Flags{
		ApiKey: *apiKey,
		OutputFile: OutputFile{
			Path:   *path,
			Name:   *name,
			Prefix: *prefix,
		},
	}
	return flags

}

func SetFileName(outputFile OutputFile, defaultPrefix string) string {
	if outputFile.Name != "" {
		return outputFile.Name
	}
	prefix := defaultPrefix
	if outputFile.Prefix != "" {
		prefix = outputFile.Prefix

	}
	// format YYYMMDD-HHMMSS
	t := time.Now().Format("20060102-151405")
	fileName := prefix + "_" + t + ".csv"
	return fileName
}

func WriteCsv(path string, fileName string, records [][]string) error {

	if path == "" {
		path = "."
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {

			return fmt.Errorf("directory does not exist, and unable to create directory: %s", errDir)
		}
	}

	fullPath := filepath.Join(path, fileName)
	csvFile, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	err = csvWriter.WriteAll(records)
	if err != nil {
		return err
	}
	return nil
}
