package chatbot

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// MockWeatherAPI simula una API de clima para pruebas
type MockWeatherAPI struct {
	Data map[string]WeatherData
}

// NewMockWeatherAPI crea una API de clima simulada
func NewMockWeatherAPI() *MockWeatherAPI {
	return &MockWeatherAPI{
		Data: map[string]WeatherData{
			"madrid": {
				City:        "Madrid",
				Temperature: 25.5,
				Condition:   "Soleado",
				Humidity:    45,
			},
			"barcelona": {
				City:        "Barcelona",
				Temperature: 23.0,
				Condition:   "Parcialmente nublado",
				Humidity:    60,
			},
			"valencia": {
				City:        "Valencia",
				Temperature: 28.0,
				Condition:   "Soleado",
				Humidity:    50,
			},
		},
	}
}

// GetWeather obtiene datos del clima simulados
func (m *MockWeatherAPI) GetWeather(city string) (WeatherData, error) {
	city = strings.ToLower(city)
	if data, ok := m.Data[city]; ok {
		return data, nil
	}
	return WeatherData{}, ErrCityNotFound
}

// TestNewChatbotEngine prueba la creación de un motor de chatbot
func TestNewChatbotEngine(t *testing.T) {
	config := ChatbotConfig{
		Name:             "TestBot",
		KnowledgeBasePath: filepath.Join("data", "intents.json"),
		EntitiesPath:     filepath.Join("data", "entities.json"),
	}
	
	engine, err := NewChatbotEngine(config)
	if err != nil {
		t.Fatalf("NewChatbotEngine returned error: %v", err)
	}
	
	if engine == nil {
		t.Fatal("NewChatbotEngine returned nil engine")
	}
	
	if engine.Config.Name != "TestBot" {
		t.Errorf("Engine name = %s, want TestBot", engine.Config.Name)
	}
	
	// Verificar que se cargaron las intenciones
	if len(engine.Intents) == 0 {
		t.Error("No intents were loaded")
	}
	
	// Verificar que se registraron los manejadores de intención
	if len(engine.IntentHandlers) == 0 {
		t.Error("No intent handlers were registered")
	}
}

// TestLoadKnowledgeBase prueba la carga de la base de conocimiento
func TestLoadKnowledgeBase(t *testing.T) {
	engine := &ChatbotEngine{
		Intents: make(map[string]Intent),
		Config: ChatbotConfig{
			Name: "TestBot",
		},
	}
	
	err := engine.LoadKnowledgeBase(filepath.Join("data", "intents.json"))
	if err != nil {
		t.Fatalf("LoadKnowledgeBase returned error: %v", err)
	}
	
	// Verificar que se cargaron las intenciones
	expectedIntents := []string{"greeting", "farewell", "thanks", "help", "weather", "fallback"}
	for _, intentName := range expectedIntents {
		if _, ok := engine.Intents[intentName]; !ok {
			t.Errorf("Intent %s was not loaded", intentName)
		}
	}
	
	// Verificar que las intenciones tienen patrones y respuestas
	for name, intent := range engine.Intents {
		if len(intent.Patterns) == 0 && name != "fallback" {
			t.Errorf("Intent %s has no patterns", name)
		}
		if len(intent.Responses) == 0 {
			t.Errorf("Intent %s has no responses", name)
		}
	}
	
	// Probar con archivo inexistente
	err = engine.LoadKnowledgeBase("nonexistent.json")
	if err == nil {
		t.Error("LoadKnowledgeBase should return error for non-existent file")
	}
}

// TestDetectIntent prueba la detección de intenciones
func TestDetectIntent(t *testing.T) {
	// Crear un motor de chatbot con intenciones predefinidas
	engine := &ChatbotEngine{
		Intents: map[string]Intent{
			"greeting": {
				Name:     "greeting",
				Patterns: []string{"hola", "buenos días", "saludos"},
				Responses: []string{"¡Hola!", "¡Saludos!"},
			},
			"weather": {
				Name:     "weather",
				Patterns: []string{"clima", "temperatura", "lluvia"},
				Responses: []string{"El clima es..."},
			},
			"fallback": {
				Name:     "fallback",
				Patterns: []string{},
				Responses: []string{"No entiendo"},
			},
		},
	}
	
	tests := []struct {
		name         string
		message      string
		expectedIntent string
		minConfidence float64
	}{
		{"greeting exact", "hola", "greeting", 0.9},
		{"greeting partial", "hola, ¿cómo estás?", "greeting", 0.7},
		{"weather exact", "clima", "weather", 0.9},
		{"weather context", "¿cómo está el clima hoy?", "weather", 0.7},
		{"unknown", "quiero ordenar una pizza", "fallback", 0.0},
		{"empty", "", "fallback", 0.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intent, confidence := engine.DetectIntent(tt.message)
			
			if intent.Name != tt.expectedIntent {
				t.Errorf("DetectIntent(%q) intent = %s, want %s", tt.message, intent.Name, tt.expectedIntent)
			}
			
			if confidence < tt.minConfidence && tt.expectedIntent != "fallback" {
				t.Errorf("DetectIntent(%q) confidence = %f, want at least %f", tt.message, confidence, tt.minConfidence)
			}
		})
	}
}

// TestProcessMessage prueba el procesamiento de mensajes
func TestProcessMessage(t *testing.T) {
	// Crear un motor de chatbot
	config := ChatbotConfig{
		Name:             "TestBot",
		KnowledgeBasePath: filepath.Join("data", "intents.json"),
		EntitiesPath:     filepath.Join("data", "entities.json"),
	}
	
	engine, err := NewChatbotEngine(config)
	if err != nil {
		t.Fatalf("NewChatbotEngine returned error: %v", err)
	}
	
	// Reemplazar la API de clima con un mock
	mockWeatherAPI := NewMockWeatherAPI()
	engine.WeatherAPI = mockWeatherAPI
	
	tests := []struct {
		name           string
		userID         string
		message        string
		expectContains string
	}{
		{"greeting", "user1", "hola", "Hola"},
		{"farewell", "user1", "adiós", "Adiós"},
		{"thanks", "user1", "gracias", "nada"},
		{"help", "user1", "necesito ayuda", "ayudarte"},
		{"weather query", "user1", "¿cómo está el clima?", "ubicación"},
		{"weather with city", "user1", "clima en Madrid", "Madrid"},
		{"unknown", "user1", "quiero ordenar una pizza", "No estoy seguro"},
		{"empty", "user1", "", "No entendí"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := engine.ProcessMessage(tt.userID, tt.message)
			
			if err != nil {
				t.Fatalf("ProcessMessage returned error: %v", err)
			}
			
			if !strings.Contains(strings.ToLower(response), strings.ToLower(tt.expectContains)) {
				t.Errorf("ProcessMessage(%q) = %q, should contain %q", tt.message, response, tt.expectContains)
			}
		})
	}
}

// TestConversationContext prueba el mantenimiento del contexto de conversación
func TestConversationContext(t *testing.T) {
	// Crear un motor de chatbot
	config := ChatbotConfig{
		Name:             "TestBot",
		KnowledgeBasePath: filepath.Join("data", "intents.json"),
		EntitiesPath:     filepath.Join("data", "entities.json"),
	}
	
	engine, err := NewChatbotEngine(config)
	if err != nil {
		t.Fatalf("NewChatbotEngine returned error: %v", err)
	}
	
	// Reemplazar la API de clima con un mock
	mockWeatherAPI := NewMockWeatherAPI()
	engine.WeatherAPI = mockWeatherAPI
	
	// Simular una conversación con contexto
	userID := "user2"
	
	// Primer mensaje: preguntar por el clima
	response1, err := engine.ProcessMessage(userID, "¿cómo está el clima?")
	if err != nil {
		t.Fatalf("ProcessMessage returned error: %v", err)
	}
	
	if !strings.Contains(strings.ToLower(response1), "ubicación") {
		t.Errorf("First response should ask for location, got: %s", response1)
	}
	
	// Segundo mensaje: proporcionar ubicación
	response2, err := engine.ProcessMessage(userID, "Madrid")
	if err != nil {
		t.Fatalf("ProcessMessage returned error: %v", err)
	}
	
	// Debería responder con información del clima para Madrid
	if !strings.Contains(strings.ToLower(response2), "madrid") {
		t.Errorf("Second response should contain city name, got: %s", response2)
	}
	
	// Verificar que el contexto se mantiene
	conversation, exists := engine.Conversations[userID]
	if !exists {
		t.Fatal("Conversation context was not stored")
	}
	
	if conversation.Context["city"] != "Madrid" {
		t.Errorf("Context city = %s, want Madrid", conversation.Context["city"])
	}
	
	// Tercer mensaje: preguntar "y mañana" (debería mantener el contexto de la ciudad)
	response3, err := engine.ProcessMessage(userID, "¿y mañana?")
	if err != nil {
		t.Fatalf("ProcessMessage returned error: %v", err)
	}
	
	if !strings.Contains(strings.ToLower(response3), "madrid") {
		t.Errorf("Third response should maintain city context, got: %s", response3)
	}
}

// TestRegisterIntentHandler prueba el registro de manejadores de intención
func TestRegisterIntentHandler(t *testing.T) {
	engine := &ChatbotEngine{
		IntentHandlers: make(map[string]IntentHandlerFunc),
	}
	
	// Registrar un manejador de prueba
	testHandler := func(engine *ChatbotEngine, userID string, intent Intent, entities map[string]string) (string, error) {
		return "Test response", nil
	}
	
	engine.RegisterIntentHandler("test_intent", testHandler)
	
	// Verificar que el manejador se registró correctamente
	handler, exists := engine.IntentHandlers["test_intent"]
	if !exists {
		t.Fatal("Intent handler was not registered")
	}
	
	// Probar el manejador
	response, err := handler(engine, "user1", Intent{Name: "test_intent"}, nil)
	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}
	
	if response != "Test response" {
		t.Errorf("Handler response = %s, want 'Test response'", response)
	}
}