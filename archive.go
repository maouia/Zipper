package main

import (
	"archive/zip"
	"fmt"
	"io"
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

func main() {

	//creating simple archive folder to store the data in
	println("Creating zip archive.....")
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	//opening file
	println("Opening first file.....")
	f1, err := os.Open(getPath())
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	_, fileName := filepath.Split(f1.Name())

	println("Adding data file to archive")
	w1, err := zipWriter.Create(fileName)
	if err != nil {
		panic(err)
	}

	//copy uncompressed file to archive
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}

	//Closing archive
	println("Every thing complited we just closing the archive now")
	archive.Close()

}
