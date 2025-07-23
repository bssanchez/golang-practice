# Ejercicio 13: Motor de Chatbot

## Descripción
Implementa un motor de chatbot avanzado que pueda procesar lenguaje natural, mantener contexto de conversación y conectarse a APIs externas para obtener información.

## Requisitos
1. Implementa las siguientes estructuras y funciones en el archivo `chatbot.go`:
   - Struct `ChatbotEngine` con campos para configuración, estado y procesadores de intención
   - Struct `Conversation` para mantener el historial y contexto de una conversación
   - Struct `Message` para representar mensajes del usuario y del chatbot
   - Struct `Intent` para representar la intención detectada en un mensaje
   - Método `NewChatbotEngine(config ChatbotConfig) (*ChatbotEngine, error)` - Constructor
   - Método `ProcessMessage(userID string, message string) (string, error)` - Procesa un mensaje y genera respuesta
   - Método `DetectIntent(message string) (Intent, float64)` - Detecta la intención del usuario
   - Método `GenerateResponse(userID string, intent Intent) (string, error)` - Genera respuesta basada en intención
   - Método `LoadKnowledgeBase(filePath string) error` - Carga base de conocimiento desde archivo
   - Método `RegisterIntentHandler(intentName string, handler IntentHandlerFunc)` - Registra manejador para intención

2. Implementa los siguientes procesadores de intención:
   - `WeatherIntentHandler` - Obtiene información del clima
   - `GreetingIntentHandler` - Maneja saludos
   - `FarewellIntentHandler` - Maneja despedidas
   - `HelpIntentHandler` - Proporciona ayuda sobre capacidades del chatbot
   - `FallbackIntentHandler` - Maneja intenciones no reconocidas

3. Consideraciones:
   - Implementa un sistema de detección de intenciones basado en palabras clave y patrones
   - Mantén el contexto de la conversación para referencias pronominales
   - Implementa un sistema de memoria a corto y largo plazo
   - Conecta con APIs externas para obtener información (clima, noticias, etc.)
   - Implementa un sistema de retroalimentación para mejorar respuestas

## Pruebas
Ejecuta `go test` para verificar tu implementación.

## Datos de Ejemplo
El directorio `data` contiene archivos con patrones de intención y respuestas para entrenar el chatbot.