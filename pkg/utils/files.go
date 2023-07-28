package utils

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
)

// Método para leer un archivo
func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Cannot read file:", err.Error())
		return ""
	}

	return string(file)
}

// Leer y contar las palabras/lineas de un fichero
func ReadAndCountFile(fileName string) {
	read := ReadFile(fileName)
}

// Leer los ficheros de un directorio
func ReadDir(fileName string) []fs.DirEntry {
	files, err := os.ReadDir(fileName)
	if err != nil {
		log.Fatal("Cannot read dir:", err.Error())
		return []fs.DirEntry{}
	}

	return files
}

// Filtrar los archivos .txt
func FilterDir(fileName string) string {
	files := ReadDir(fileName)
	out := ""
	for i, de := range files {
		if path.Ext(de.Name()) == ".txt" {
			out += fmt.Sprintf("%d. %s\n", (i + 1), de.Name())
		}
	}

	return out
}
