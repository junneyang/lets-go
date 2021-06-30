package main

import (
	"bufio"
	"errors"
	"os"
)

func fileOps(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		// fmt.Println(err)
		return nil, errors.New(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// line, err := reader.ReadBytes('\n')
	// if err != nil {
	// 	fmt.Println(line)
	// }
	res := make([]string, 0)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		res = append(res, string(line))
	}
	return res, nil
}
