package utils

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

// Método para leer un archivo
func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Cannot read file: ", err.Error())
		return ""
	}

	return string(file)
}

func ReadAndCountWordAndLines(fileName string) (wordCount, lineCount uint) {
	read := ReadFile(fileName)

	for _, line := range strings.Split(read, "\n") {
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}

		lineCount++

		for _, word := range strings.Split(line, " ") {
			word = strings.Trim(word, " ")
			if len(word) != 0 && word != " " {
				wordCount++
			}
		}
	}

	return
}

// Leer y contar las palabras/lineas de un fichero
func ReadAndCountWords(fileName string) uint {
	read := ReadFile(fileName)

	var wordCount uint
	for _, word := range strings.Split(read, " ") {
		word = strings.Trim(word, " ")
		if len(word) != 0 && word != " " {
			wordCount++
		}
	}

	return wordCount
}

func ReadAndCountLines(fileName string) int {
	read := ReadFile(fileName)

	return len(strings.Split(read, "\n"))
}

// Leer los ficheros de un directorio
func ReadDir(fileName string) []fs.DirEntry {
	files, err := os.ReadDir(fileName)
	if err != nil {
		log.Fatal("Cannot read dir: ", err.Error())
		return []fs.DirEntry{}
	}

	return files
}

// Filtrar los archivos .txt
func FilterDir(fileName string) string {
	files := ReadDir(fileName)
	out := ""
	cnt := 1
	for _, de := range files {
		if path.Ext(de.Name()) == ".txt" {
			out += fmt.Sprintf("%d. %s\n", cnt, de.Name())
			cnt++
		}
	}

	return out
}
