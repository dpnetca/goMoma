package outfile

import (
	"encoding/csv"
	"os"
	"time"
)

func WriteOutput(filePrefix string, records [][]string) error {
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
	for _, record := range records {
		err := csvWriter.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
