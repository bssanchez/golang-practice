# Ejercicio 7: Cliente HTTP

## Descripción
Implementa un cliente HTTP simple que pueda realizar peticiones y procesar respuestas.

## Requisitos
1. Implementa las siguientes funciones en el archivo `httpclient.go`:
   - `Get(url string) ([]byte, error)` - Realiza una petición GET y devuelve el cuerpo de la respuesta
   - `GetJSON(url string, v interface{}) error` - Realiza una petición GET y deserializa la respuesta JSON
   - `Post(url string, contentType string, body []byte) ([]byte, error)` - Realiza una petición POST
   - `DownloadFile(url, filepath string) error` - Descarga un archivo desde una URL
   - `FetchWithTimeout(url string, timeout time.Duration) ([]byte, error)` - Realiza una petición GET con timeout

2. Consideraciones:
   - Maneja correctamente los errores HTTP (códigos de estado, timeouts, etc.)
   - Para `GetJSON`, usa la librería estándar `encoding/json`
   - Para `DownloadFile`, escribe el contenido en el archivo de forma eficiente
   - Implementa un mecanismo para cancelar peticiones que tarden demasiado

## Pruebas
Ejecuta `go test` para verificar tu implementación.

## Nota
Las pruebas utilizan un servidor HTTP mock para evitar dependencias externas.