package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManger struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManger) ReadLine() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println(err)
		file.Close()

		return nil, err
	}

	file.Close()

	return lines, nil
}

func (fm FileManger) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return err
	}

	enconder := json.NewEncoder(file)

	err = enconder.Encode(data)

	file.Close()

	return err
}

func New(inputPath, outputPath string) FileManger {
	return FileManger{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
