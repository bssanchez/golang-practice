package downloader

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// DownloadResult contiene el resultado de una operación de descarga
type DownloadResult struct {
	URL      string        // URL del archivo descargado
	FilePath string        // Ruta donde se guardó el archivo
	Error    error         // Error durante la descarga (nil si fue exitosa)
	Duration time.Duration // Tiempo que tomó la descarga
}

// Downloader gestiona descargas concurrentes de archivos
type Downloader struct {
	maxConcurrent int // Número máximo de descargas concurrentes
}

// NewDownloader crea un nuevo descargador con un límite de concurrencia
func NewDownloader(maxConcurrent int) *Downloader {
	// TODO: Implementar esta función
	return nil
}

// Download descarga un archivo desde una URL y lo guarda en la ruta especificada
func (d *Downloader) Download(url, filepath string) error {
	// TODO: Implementar esta función
	return nil
}

// DownloadMany descarga múltiples archivos concurrentemente
// Respeta el límite de concurrencia especificado en el Downloader
func (d *Downloader) DownloadMany(urls []string, directory string) []DownloadResult {
	// TODO: Implementar esta función
	return nil
}

// DownloadWithProgress descarga un archivo y reporta el progreso a través de un canal
// El progreso se reporta como un porcentaje (0-100)
func (d *Downloader) DownloadWithProgress(url, filepath string, progressChan chan<- int) error {
	// TODO: Implementar esta función
	return nil
}