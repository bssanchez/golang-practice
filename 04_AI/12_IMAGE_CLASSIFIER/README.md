# Ejercicio 12: Clasificador de Imágenes

## Descripción
Implementa un clasificador de imágenes simple utilizando un modelo pre-entrenado y la biblioteca GoCV (OpenCV para Go).

## Requisitos
1. Implementa las siguientes funciones en el archivo `classifier.go`:
   - `NewImageClassifier(modelPath, configPath, classesPath string) (*ImageClassifier, error)` - Carga un modelo pre-entrenado
   - `ClassifyImage(imagePath string) ([]Prediction, error)` - Clasifica una imagen y devuelve predicciones
   - `ClassifyImageBatch(imagePaths []string) (map[string][]Prediction, error)` - Clasifica múltiples imágenes
   - `GetTopNPredictions(predictions []Prediction, n int) []Prediction` - Obtiene las N predicciones con mayor confianza

2. Implementa las siguientes estructuras:
   - `ImageClassifier` con campos para el modelo, configuración y clases
   - `Prediction` con campos para la clase predicha y su confianza

3. Consideraciones:
   - Utiliza GoCV para cargar y procesar imágenes
   - Implementa preprocesamiento de imágenes (redimensionar, normalizar)
   - Maneja correctamente los errores de carga de modelo e imágenes
   - Optimiza el rendimiento para clasificación por lotes

## Dependencias
- GoCV (github.com/hybridgroup/gocv)

## Instalación de Dependencias
```bash
go get -u github.com/hybridgroup/gocv
```

## Pruebas
Ejecuta `go test` para verificar tu implementación.

## Nota
Para las pruebas, se incluye un modelo pre-entrenado simplificado en el directorio `models`.