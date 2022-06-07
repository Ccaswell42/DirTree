package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func get_size(file fs.FileInfo) string {
	var c string
	if file.IsDir() {
		return file.Name()
	} else {
		if file.Size() == 0 {
			c = " (empty)"
		} else {
			c = " ("
			c = c + strconv.Itoa(int(file.Size()))
			c = c + "b)"
		}
	}
	return file.Name() + c

}
func getPrefix(file fs.FileInfo, prefix string, islast bool) string {
	if islast {
		prefix += "└───"
	} else {
		prefix += "├───"
	}
	prefix += get_size(file)

	return prefix
}
func get_Tab(tab string, islast bool) string {
	if !islast {
		return tab + "│\t"
	}
	return tab + "\t"
}
func get_path(str string, prefix string) ([]string, error) {
	var a []string
	filesFromDir, err := ioutil.ReadDir(str)

	if err != nil {
		return a, err
	}
	lstElement := len(filesFromDir) - 1
	for i, file := range filesFromDir {
		isLast := lstElement == i
		a = append(a, getPrefix(file, prefix, isLast))
		if file.IsDir() {
			items, err := get_path(str+"/"+file.Name(), get_Tab(prefix, isLast))
			if err != nil {
				return a, err
			}
			a = append(a, items...)
		}
	}
	return a, nil
}

func main() {

	path := os.Args[1]
	str, err := get_path(path, "")
	if err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
	fmt.Println(strings.Join(str, "\n"))

}
