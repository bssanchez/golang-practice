# Ejercicio 8: Descargador Concurrente

## Descripción
Implementa un sistema de descarga concurrente que pueda descargar múltiples archivos simultáneamente con límites de concurrencia.

## Requisitos
1. Implementa las siguientes estructuras y funciones en el archivo `downloader.go`:
   - Struct `DownloadResult` con campos para URL, ruta local, error y tiempo de descarga
   - Struct `Downloader` con campo para el número máximo de descargas concurrentes
   - Método `NewDownloader(maxConcurrent int) *Downloader` - Constructor
   - Método `Download(url, filepath string) error` - Descarga un solo archivo
   - Método `DownloadMany(urls []string, directory string) []DownloadResult` - Descarga múltiples archivos
   - Método `DownloadWithProgress(url, filepath string, progressChan chan<- int) error` - Descarga con notificación de progreso

2. Consideraciones:
   - Utiliza goroutines para descargas concurrentes
   - Limita el número de descargas simultáneas según `maxConcurrent`
   - Implementa un mecanismo para reportar el progreso de descarga
   - Maneja correctamente los errores y timeouts
   - Asegúrate de que los recursos se liberen adecuadamente

## Pruebas
Ejecuta `go test` para verificar tu implementación.

## Nota
Las pruebas utilizan un servidor HTTP mock para evitar dependencias externas.