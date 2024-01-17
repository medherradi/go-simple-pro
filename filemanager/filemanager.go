package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadFilesLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error occured while opening file: " + err.Error())
		// file.Close() we dont need to close at first error the file
		return nil, errors.New("error occured while opening file: " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() { // we use append bcs we use scan otherwise we should use lines[idx] = val
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("reading content in file failed: " + err.Error())
		file.Close()
		return nil, errors.New("reading content in file failed: " + err.Error())
	}
	file.Close()
	return lines, nil
}

// storing data as json
func StoreDataAsJson(pathFile string, data interface{}) error {
	file, err := os.Create(pathFile)
	if err != nil {
		return errors.New("failed to create file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json ")
	}
	file.Close()
	return nil
}
