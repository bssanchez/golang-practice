package chatbot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// RealWeatherAPI implementa la interfaz WeatherAPI usando una API real
type RealWeatherAPI struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// NewWeatherAPI crea una nueva instancia de la API de clima
func NewWeatherAPI(apiKey string) *RealWeatherAPI {
	return &RealWeatherAPI{
		APIKey:  apiKey,
		BaseURL: "https://api.openweathermap.org/data/2.5",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetWeather obtiene datos del clima para una ciudad
func (api *RealWeatherAPI) GetWeather(city string) (WeatherData, error) {
	// Construir URL de la API
	endpoint := fmt.Sprintf("%s/weather?q=%s&units=metric&lang=es&appid=%s",
		api.BaseURL, url.QueryEscape(city), api.APIKey)
	
	// Realizar petición HTTP
	resp, err := api.HTTPClient.Get(endpoint)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()
	
	// Verificar código de estado
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return WeatherData{}, ErrCityNotFound
		}
		return WeatherData{}, fmt.Errorf("error de la API: código %d", resp.StatusCode)
	}
	
	// Decodificar respuesta
	var apiResponse struct {
		Name string `json:"name"`
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return WeatherData{}, fmt.Errorf("error al decodificar respuesta: %w", err)
	}
	
	// Construir datos del clima
	weatherData := WeatherData{
		City:        apiResponse.Name,
		Temperature: apiResponse.Main.Temp,
		Humidity:    apiResponse.Main.Humidity,
	}
	
	// Obtener descripción del clima
	if len(apiResponse.Weather) > 0 {
		weatherData.Condition = apiResponse.Weather[0].Description
	}
	
	return weatherData, nil
}