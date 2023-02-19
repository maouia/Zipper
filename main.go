package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		print("1) : to zip \n2) : to unzip \n3) : to exit   ")
		var varibale string
		fmt.Scan(&varibale)
		switch varibale {
		case "1":
			Zipping()
		case "2":
			Unzipping()
		case "3":
			os.Exit(1)
		default:
			println("your choice is invalid make sure you choose something exist please !!")
			continue
		}
	}

}
