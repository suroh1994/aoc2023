package aoccommon

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
)

func ReadInputLineByLine(inputFileName string) []string {
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		return nil
	}

	inputScanner := bufio.NewScanner(bytes.NewReader(content))
	inputScanner.Split(bufio.ScanLines)
	var inputLines []string

	for inputScanner.Scan() {
		inputLines = append(inputLines, inputScanner.Text())
	}
	return inputLines
}

func WriteLineByLine(inputFileName string, lines interface{}) error {
	val := reflect.ValueOf(lines)
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("invalid type, expected a slice but received a %q", val.Kind())
	}

	count := val.Len()
	slice := make([]interface{}, count)
	for i := 0; i < count; i++ {
		slice[i] = val.Index(i).Interface()
	}

	f, err := os.Create(inputFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, line := range slice {
		_, err = writer.WriteString(fmt.Sprintf("%v\n", line))
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	return err
}
