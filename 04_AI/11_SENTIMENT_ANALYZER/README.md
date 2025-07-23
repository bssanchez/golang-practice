# Ejercicio 11: Analizador de Sentimientos

## Descripción
Implementa un analizador de sentimientos simple que pueda clasificar textos como positivos, negativos o neutros utilizando un enfoque basado en diccionario.

## Requisitos
1. Implementa las siguientes funciones en el archivo `sentiment.go`:
   - `LoadSentimentDictionary(positiveWordsFile, negativeWordsFile string) (*SentimentDictionary, error)` - Carga palabras positivas y negativas desde archivos
   - `AnalyzeSentiment(text string, dictionary *SentimentDictionary) SentimentResult` - Analiza el sentimiento de un texto
   - `AnalyzeMultipleTexts(texts []string, dictionary *SentimentDictionary) []SentimentResult` - Analiza múltiples textos
   - `GetSentimentScore(text string, dictionary *SentimentDictionary) float64` - Calcula una puntuación numérica de sentimiento

2. Implementa las siguientes estructuras:
   - `SentimentDictionary` con campos para palabras positivas y negativas
   - `SentimentResult` con campos para el texto original, la clasificación y la puntuación

3. Consideraciones:
   - Normaliza el texto (minúsculas, eliminar puntuación) antes del análisis
   - Implementa un sistema de puntuación que considere la intensidad del sentimiento
   - Maneja correctamente casos especiales como negaciones ("no es bueno")

## Pruebas
Ejecuta `go test` para verificar tu implementación.

## Datos de Ejemplo
El directorio `data` contiene archivos con palabras positivas y negativas para entrenar el analizador.