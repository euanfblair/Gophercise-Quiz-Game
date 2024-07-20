package reader

import (
	"encoding/csv"
	"flag"
	"os"
)

func ReadCsvToSlice(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func FileFlag(flagName, defaultVal, desc string) string {

	fileFlag := flag.String(flagName, defaultVal, desc)
	fileProvided := false
	for _, arg := range os.Args {
		if arg == "-"+flagName || arg == "--"+flagName {
			fileProvided = true
			break
		}
	}

	flag.Parse()

	if fileProvided {
		return *fileFlag
	}

	return defaultVal
}
