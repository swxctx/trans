package trans

import (
	"encoding/csv"
	"os"
)

// readCsv 读取csv文件
func readCsv(file string) ([][]string, error) {
	var (
		rows [][]string
		err  error
	)
	csvFile, err := os.Open(file)
	if err != nil {
		return rows, err
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	rows, err = csvReader.ReadAll()
	if err != nil {
		return rows, err
	}
	return rows, nil
}
