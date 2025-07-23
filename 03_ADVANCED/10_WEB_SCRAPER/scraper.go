package scraper

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ScrapedPage contiene la información extraída de una página web
type ScrapedPage struct {
	URL     string   // URL de la página
	Title   string   // Título de la página
	Text    string   // Texto principal de la página
	Links   []string // Enlaces encontrados en la página
	Depth   int      // Profundidad de la página en el árbol de scraping
}

// Scraper implementa un web scraper concurrente
type Scraper struct {
	MaxDepth      int    // Profundidad máxima de scraping
	MaxPages      int    // Número máximo de páginas a scrapear
	AllowedDomain string // Dominio permitido para el scraping
}

// NewScraper crea un nuevo scraper con los límites especificados
func NewScraper(maxDepth, maxPages int, allowedDomain string) *Scraper {
	// TODO: Implementar esta función
	return nil
}

// Scrape inicia el proceso de scraping desde una URL inicial
// Devuelve las páginas scrapeadas y un posible error
func (s *Scraper) Scrape(startURL string) ([]ScrapedPage, error) {
	// TODO: Implementar esta función
	return nil, nil
}

// ExtractLinks extrae todos los enlaces de un documento HTML
// Convierte URLs relativas a absolutas usando la URL base
func (s *Scraper) ExtractLinks(html string, baseURL string) ([]string, error) {
	// TODO: Implementar esta función
	return nil, nil
}

// ExtractTitle extrae el título de un documento HTML
func (s *Scraper) ExtractTitle(html string) string {
	// TODO: Implementar esta función
	return ""
}

// ExtractText extrae el texto principal de un documento HTML
// Elimina tags HTML, scripts, estilos, etc.
func (s *Scraper) ExtractText(html string) string {
	// TODO: Implementar esta función
	return ""
}