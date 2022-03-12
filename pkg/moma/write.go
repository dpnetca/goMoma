package moma

import (
	"encoding/csv"
	"os"
	"time"
)

func WriteCsv(filePrefix string, records [][]string) error {
	// TODO add support for directory path

	// format YYYMMDD-HHMMSS
	t := time.Now().Format("20060102-151405")
	fileName := filePrefix + "_" + t + ".csv"
	csvFile, err := os.Create(fileName)
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
