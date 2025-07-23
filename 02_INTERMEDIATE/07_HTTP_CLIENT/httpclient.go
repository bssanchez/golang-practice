package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

// Get realiza una petición HTTP GET y devuelve el cuerpo de la respuesta
func Get(url string) ([]byte, error) {
	// TODO: Implementar esta función
	return nil, nil
}

// GetJSON realiza una petición HTTP GET y deserializa la respuesta JSON
func GetJSON(url string, v interface{}) error {
	// TODO: Implementar esta función
	return nil
}

// Post realiza una petición HTTP POST y devuelve el cuerpo de la respuesta
func Post(url string, contentType string, body []byte) ([]byte, error) {
	// TODO: Implementar esta función
	return nil, nil
}

// DownloadFile descarga un archivo desde una URL y lo guarda en la ruta especificada
func DownloadFile(url, filepath string) error {
	// TODO: Implementar esta función
	return nil
}

// FetchWithTimeout realiza una petición HTTP GET con un timeout especificado
func FetchWithTimeout(url string, timeout time.Duration) ([]byte, error) {
	// TODO: Implementar esta función
	return nil, nil
}