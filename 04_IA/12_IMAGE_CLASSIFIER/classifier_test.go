package classifier

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// MockImageClassifier implementa un clasificador de imágenes simulado para pruebas
type MockImageClassifier struct {
	Classes []string
}

// NewMockClassifier crea un clasificador simulado para pruebas
func NewMockClassifier() *MockImageClassifier {
	return &MockImageClassifier{
		Classes: []string{
			"persona", "gato", "perro", "coche", "bicicleta",
			"árbol", "edificio", "teléfono", "libro", "computadora",
		},
	}
}

// MockClassifyImage simula la clasificación de una imagen
func (m *MockImageClassifier) MockClassifyImage(imagePath string) []Prediction {
	// Extraer el nombre del archivo sin extensión
	filename := filepath.Base(imagePath)
	ext := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(ext)]
	
	// Generar predicciones simuladas basadas en el nombre del archivo
	var predictions []Prediction
	
	// Si el nombre contiene alguna de nuestras clases, darle alta confianza
	highConfidenceClass := ""
	for _, class := range m.Classes {
		if name == class || name == "imagen_"+class {
			highConfidenceClass = class
			break
		}
	}
	
	// Generar predicciones
	for i, class := range m.Classes {
		confidence := 0.1 // Confianza base baja
		
		if class == highConfidenceClass {
			confidence = 0.95 // Alta confianza para la clase que coincide
		} else if i < 3 {
			confidence = 0.3 - float64(i)*0.1 // Confianzas medias para las primeras clases
		}
		
		predictions = append(predictions, Prediction{
			Class:      class,
			Confidence: confidence,
		})
	}
	
	// Ordenar por confianza (mayor a menor)
	sort.Slice(predictions, func(i, j int) bool {
		return predictions[i].Confidence > predictions[j].Confidence
	})
	
	return predictions
}

// TestNewImageClassifier prueba la creación de un clasificador
func TestNewImageClassifier(t *testing.T) {
	// Paths a archivos de modelo simulados
	modelPath := filepath.Join("models", "model.caffemodel")
	configPath := filepath.Join("models", "model.prototxt")
	classesPath := filepath.Join("models", "classes.txt")
	
	// Crear archivos temporales para la prueba
	if err := os.WriteFile(modelPath, []byte("mock model data"), 0644); err != nil {
		t.Fatalf("Failed to create mock model file: %v", err)
	}
	if err := os.WriteFile(configPath, []byte("mock config data"), 0644); err != nil {
		t.Fatalf("Failed to create mock config file: %v", err)
	}
	
	// El archivo classes.txt ya debería existir
	
	// Intentar crear un clasificador
	classifier, err := NewImageClassifier(modelPath, configPath, classesPath)
	
	// Limpiar archivos temporales
	os.Remove(modelPath)
	os.Remove(configPath)
	
	// Verificar resultado
	if err != nil {
		t.Fatalf("NewImageClassifier returned error: %v", err)
	}
	
	if classifier == nil {
		t.Fatal("NewImageClassifier returned nil classifier")
	}
	
	// Verificar que se cargaron las clases
	if len(classifier.Classes) == 0 {
		t.Error("No classes were loaded")
	}
	
	// Probar con archivos inexistentes
	_, err = NewImageClassifier("nonexistent.model", configPath, classesPath)
	if err == nil {
		t.Error("NewImageClassifier should return error for non-existent model file")
	}
	
	_, err = NewImageClassifier(modelPath, "nonexistent.config", classesPath)
	if err == nil {
		t.Error("NewImageClassifier should return error for non-existent config file")
	}
	
	_, err = NewImageClassifier(modelPath, configPath, "nonexistent.classes")
	if err == nil {
		t.Error("NewImageClassifier should return error for non-existent classes file")
	}
}

// TestClassifyImage prueba la clasificación de una imagen
func TestClassifyImage(t *testing.T) {
	// Crear un clasificador mock
	mockClassifier := NewMockClassifier()
	
	// Crear un clasificador real con el mock interno
	classifier := &ImageClassifier{
		Classes: mockClassifier.Classes,
		// Los campos Model y Config serían nil en este test
	}
	
	// Sobreescribir la función de clasificación para usar el mock
	originalClassifyFunc := classifyImageFunc
	defer func() { classifyImageFunc = originalClassifyFunc }()
	
	classifyImageFunc = func(ic *ImageClassifier, imagePath string) ([]Prediction, error) {
		return mockClassifier.MockClassifyImage(imagePath), nil
	}
	
	// Probar clasificación con diferentes imágenes simuladas
	testCases := []struct {
		name          string
		imagePath     string
		expectedClass string
	}{
		{"cat image", "imagen_gato.jpg", "gato"},
		{"dog image", "perro.png", "perro"},
		{"person image", "persona.jpg", "persona"},
		{"unknown image", "desconocido.jpg", mockClassifier.Classes[0]}, // Primera clase por defecto
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			predictions, err := classifier.ClassifyImage(tc.imagePath)
			
			if err != nil {
				t.Fatalf("ClassifyImage returned error: %v", err)
			}
			
			if len(predictions) == 0 {
				t.Fatal("ClassifyImage returned no predictions")
			}
			
			// La primera predicción debería ser la de mayor confianza
			if predictions[0].Class != tc.expectedClass {
				t.Errorf("Top prediction = %s, want %s", predictions[0].Class, tc.expectedClass)
			}
			
			// Verificar que las predicciones estén ordenadas por confianza
			for i := 1; i < len(predictions); i++ {
				if predictions[i-1].Confidence < predictions[i].Confidence {
					t.Errorf("Predictions not sorted by confidence: %v", predictions)
					break
				}
			}
		})
	}
}

// TestClassifyImageBatch prueba la clasificación de múltiples imágenes
func TestClassifyImageBatch(t *testing.T) {
	// Crear un clasificador mock
	mockClassifier := NewMockClassifier()
	
	// Crear un clasificador real con el mock interno
	classifier := &ImageClassifier{
		Classes: mockClassifier.Classes,
		// Los campos Model y Config serían nil en este test
	}
	
	// Sobreescribir la función de clasificación para usar el mock
	originalClassifyFunc := classifyImageFunc
	defer func() { classifyImageFunc = originalClassifyFunc }()
	
	classifyImageFunc = func(ic *ImageClassifier, imagePath string) ([]Prediction, error) {
		return mockClassifier.MockClassifyImage(imagePath), nil
	}
	
	// Imágenes para clasificar
	imagePaths := []string{
		"imagen_gato.jpg",
		"perro.png",
		"persona.jpg",
	}
	
	// Clasificar lote de imágenes
	results, err := classifier.ClassifyImageBatch(imagePaths)
	
	if err != nil {
		t.Fatalf("ClassifyImageBatch returned error: %v", err)
	}
	
	// Verificar que tenemos resultados para todas las imágenes
	if len(results) != len(imagePaths) {
		t.Errorf("ClassifyImageBatch returned %d results, want %d", len(results), len(imagePaths))
	}
	
	// Verificar resultados individuales
	expectedClasses := map[string]string{
		"imagen_gato.jpg": "gato",
		"perro.png":       "perro",
		"persona.jpg":     "persona",
	}
	
	for path, predictions := range results {
		if len(predictions) == 0 {
			t.Errorf("No predictions for image: %s", path)
			continue
		}
		
		expectedClass := expectedClasses[path]
		if predictions[0].Class != expectedClass {
			t.Errorf("Top prediction for %s = %s, want %s", path, predictions[0].Class, expectedClass)
		}
	}
}

// TestGetTopNPredictions prueba la obtención de las N predicciones principales
func TestGetTopNPredictions(t *testing.T) {
	// Crear predicciones de prueba
	predictions := []Prediction{
		{Class: "gato", Confidence: 0.8},
		{Class: "perro", Confidence: 0.7},
		{Class: "persona", Confidence: 0.6},
		{Class: "coche", Confidence: 0.5},
		{Class: "bicicleta", Confidence: 0.4},
	}
	
	// Probar con diferentes valores de N
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"top 3", 3, 3},
		{"top 1", 1, 1},
		{"top 0", 0, 0},
		{"top 10", 10, 5}, // Solo hay 5 predicciones
		{"top -1", -1, 0}, // N negativo debería devolver 0
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			topN := GetTopNPredictions(predictions, tc.n)
			
			if len(topN) != tc.expected {
				t.Errorf("GetTopNPredictions returned %d predictions, want %d", len(topN), tc.expected)
			}
			
			// Verificar que las predicciones están ordenadas por confianza
			for i := 1; i < len(topN); i++ {
				if topN[i-1].Confidence < topN[i].Confidence {
					t.Errorf("Predictions not sorted by confidence: %v", topN)
					break
				}
			}
		})
	}
}

// Variable global para permitir sobreescribir la función de clasificación en tests
var classifyImageFunc = func(ic *ImageClassifier, imagePath string) ([]Prediction, error) {
	// Esta implementación sería reemplazada en los tests
	return nil, nil
}