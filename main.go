package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

const path = "tmp/.pmc.out"

const MockString = `key1: 490942    key2: 0:0:0
    key3: 0:0:23    key4: 2:232:32`

func main() {
	CreateFile()
	WriteFile()
	ReadFile()
}

func getData() string {
	out, err := exec.Command("syncappmmi.sh").Output()
	check(err)
	return string(out)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateFile() {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		check(err)
		file.Close()
	}
}

func WriteFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	check(err)
	_, err = file.WriteString(getData())
	check(err)
	err = file.Sync()
	check(err)
	file.Close()
}

func ReadFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	check(err)
	defer file.Close()
	out := make([]byte, 1024)
	for {
		n, err := file.Read(out)
		if err != io.EOF {
			check(err)
		}
		if n == 0 {
			break
		}
	}
	file.Close()
	Parse(string(out))
	check(err)
}

func Parse(s string) {
	fmt.Println(s)
}
