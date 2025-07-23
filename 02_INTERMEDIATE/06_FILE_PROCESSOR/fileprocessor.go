package fileprocessor

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// ReadFile lee un archivo y devuelve sus líneas como un slice de strings
func ReadFile(filename string) ([]string, error) {
	// TODO: Implementar esta función
	return nil, nil
}

// WriteFile escribe un slice de strings a un archivo, una línea por elemento
func WriteFile(filename string, lines []string) error {
	// TODO: Implementar esta función
	return nil
}

// CountWords cuenta el número de palabras en un archivo
func CountWords(filename string) (int, error) {
	// TODO: Implementar esta función
	return 0, nil
}

// FindPattern encuentra líneas en un archivo que coinciden con un patrón regex
func FindPattern(filename, pattern string) ([]string, error) {
	// TODO: Implementar esta función
	return nil, nil
}

// ReplaceInFile reemplaza todas las ocurrencias de 'old' con 'new' en un archivo
// Devuelve el número de reemplazos realizados
func ReplaceInFile(filename, old, new string) (int, error) {
	// TODO: Implementar esta función
	return 0, nil
}