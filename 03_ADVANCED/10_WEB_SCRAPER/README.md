# Ejercicio 10: Web Scraper

## Descripción
Implementa un web scraper concurrente que pueda extraer información de páginas web y seguir enlaces.

## Requisitos
1. Implementa las siguientes estructuras y funciones en el archivo `scraper.go`:
   - Struct `ScrapedPage` con campos para URL, título, enlaces y texto extraído
   - Struct `Scraper` con campos para profundidad máxima, número máximo de páginas y dominio permitido
   - Método `NewScraper(maxDepth, maxPages int, allowedDomain string) *Scraper` - Constructor
   - Método `Scrape(startURL string) ([]ScrapedPage, error)` - Inicia el scraping desde una URL
   - Método `ExtractLinks(html string, baseURL string) ([]string, error)` - Extrae enlaces de HTML
   - Método `ExtractTitle(html string) string` - Extrae el título de una página HTML
   - Método `ExtractText(html string) string` - Extrae el texto principal de una página HTML

2. Consideraciones:
   - Utiliza goroutines para scraping concurrente
   - Implementa un mecanismo para limitar la concurrencia
   - Respeta robots.txt y añade delays entre peticiones
   - Maneja correctamente URLs relativas y absolutas
   - Implementa un mecanismo para evitar ciclos (no visitar la misma URL dos veces)
   - Limita el scraping al dominio permitido

## Pruebas
Ejecuta `go test` para verificar tu implementación.

## Nota
Las pruebas utilizan un servidor HTTP mock para evitar dependencias externas.