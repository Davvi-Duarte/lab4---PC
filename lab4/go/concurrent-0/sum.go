package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filePath string, out chan int)  {
	data, _ := readFile(filePath)
// 	if err != nil {
// 		return 0, err
// 	}

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	out <- _sum
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
	sumsChannel := make(chan int)
	for _, path := range os.Args[1:] {
		go sum(path, sumsChannel)
    }
    close(sumsChannel)
    for v := range sumsChannel {
        totalSum += int64(v)
    }

		//sums[_sum] = append(sums[_sum], path)

	//fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}