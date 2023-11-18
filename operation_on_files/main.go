package main

import (
	//"bufio"
	"fmt"
	"log"
	"os"
)

func checkErr(err error, file *os.File) error {
	if err != nil { // there is an error
		log.Fatal(err)
	}
	defer file.Close()
	return nil
}

func main() {
	// The common way to work with files is to use the package os.

	//
	//

	// Creating a file
	// newFile, err := os.Create("a.txt")
	// checkErr(err, newFile)

	// Truncating the file
	// err = os.Truncate("a.txt", 0)
	// checkErr(err)

	// Files must be closed when done with.
	//newFile.Close()

	//
	//
	//

	//OPENING AND CLOSING AN EXISTING FILE.

	//file, err := os.Open("a.txt") using the open fucntion

	//Using the OpenFile fucntion gives even more power to the file
	// file, err := os.OpenFile("instruct.txt", os.O_CREATE|os.O_APPEND, 0644)
	// checkErr(err, file)

	//
	//
	//

	//GETTING FILE INFORMATION OF A FILE

	// fileInfo, err := os.Stat("a.txt")
	//p := fmt.Println
	// p(fileInfo.Name(), fileInfo.Mode(), fileInfo.Size())

	// To check if the file doesn't exist,

	// fileInfo, err = os.Stat("b.txt")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal("File doesnt exist")
	// 	}
	// }

	//
	//
	//

	//Renaming or Moving a File

	// oldFile := "a.txt"
	// new := "new.txt"
	// err = os.Rename(oldFile, new)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// REMOVING A FILE
	// err = os.Remove("instruct.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//
	//
	//

	//WRITING A BYTE SLICE TO A FILE
	// file, err := os.OpenFile("new.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fileInfo, err := os.Stat("new.txt")
	// p(fileInfo.Name(), fileInfo.Mode(), fileInfo.Size())

	// To write a byte slice to a file, we first convert the string to a byte slice then write
	// each rune/byte to the file; this can also give room to modify each byte since its mutable

	// byteSlice := []byte("I am enjoyin Golang")
	// bytesWritten, err := file.Write(byteSlice)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("No of bytes written is %d", bytesWritten)

	// Another way is to use the os.WriteFile function.

	// bc := []byte("Still enjoying Golang")
	//err = os.WriteFile("new.txt", bc, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//
	//
	//

	//READING A FILE

	// file, err := os.Open("new.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// byteSlice := make([]byte, 3)
	// numberOfBytesRead, err := io.ReadFull(file, byteSlice)

	// fmt.Printf("Number of bytes Read is %v\n", numberOfBytesRead)
	// fmt.Printf("Data read is %q\n", byteSlice) // This retuurned the number of bytes read which was 3.

	// To read the actual byte without specifying the number of bytes presumably; we use io.ReadAll.

	// file, err = os.Open("new.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// dataReader, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Data read is \n %s\n", (dataReader)) // This will print out the entire data read from the file new.txt.

	//
	//Another way is to use os.ReadFile() to read the entire file.
	// file, err := os.Open("new.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	dataReader, err := os.ReadFile("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read is \n %s\n", (dataReader)) // This will print out the entire data read from the file new.txt.

	//
	//
	//

	//SCANNING FOR USER INPUT, READING FROM STDIN
	// scanner := bufio.NewScanner(os.Stdin)

	// scanner.Scan()

	// text := scanner.Text()
	//bytes := scanner.Bytes()
	//fmt.Printf("Text entered is %s\n", text)

	// To continuously scan the stdin until a particualr string is scaned, we scan using the for loop
	// for scanner.Scan() {
	// 	text = scanner.Text()
	// 	fmt.Printf("You entered %s\n", text)
	// 	if text == "exit" {
	// 		fmt.Printf("No more entry needed\n")
	// 		break
	// 	}
	// }
	// if err := scanner.Err(); err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("Waiting for instructions...\n")
}
