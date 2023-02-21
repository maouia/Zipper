package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getPath() string {
	var path string
	fmt.Print("please type the path : ")
	_, err := fmt.Scan(&path)
	if err != nil {
		panic(err)
	}
	return path
}

func Zipping() {
	file := getPath()

	path, name_of_file := filepath.Split(file)

	f, _ := os.Open(file)

	read := bufio.NewReader(f)

	// Now we would use the variable Read All to get all the bytes
	// So we just used variable data which will read all the bytes
	data, _ := ioutil.ReadAll(read)

	//name_of_file = strings.Replace(name_of_file, filepath.Ext(name_of_file), ".gz", -1)
	name_of_file = name_of_file + ".gz"
	// Open file for writing
	// Now using the Os.create method we would use the
	// To store the information of the file gz extension
	f, _ = os.Create(path + name_of_file)

	// Write compresses Data
	// We would use NewWriter to basically
	// copy all the compressed data

	w := gzip.NewWriter(f)

	// With the help of the Writer method, we would
	// write all the bytes in the data variable
	// copied from the original file
	w.Write(data)

	// We would now close the file.
	w.Close()

}
