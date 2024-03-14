package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

//todo input sanity checks for city input
//allow lat and long input
//create a weather page to view extensive data
//convert from kelvim to celsius

type WeatherData struct {
	Name        string      `json:"Name"`
	Coord       coord       `json:"coord"`
	Weather     weather     `json:"weather"`
	WeatherTemp weathertemp `json:"weathertemp"`
}

type coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type weather struct {
	Id          float64 `json:"id"`
	Main        string  `json:"main"`
	Description string  `json:"description"`
}

type weathertemp struct {
	Temp       float64 `json:"temp"`
	Feels_like float64 `json:"feelsLike"`
	Temp_min   float64 `json:"tempMin"`
	Temp_max   float64 `json:"tempMax"`
	Pressure   float64 `json:"pressure"`
	Humidity   float64 `json:"humidity"`
}

func (w *WeatherData) getWeather(city string) string {
	var data map[string]interface{}

	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file: %v", err)
	}

	// Access environment variables
	apiKey := os.Getenv("API_KEY")

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "+&appid=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return err.Error()
	}
	errn := json.Unmarshal(resBody, &data)
	if errn != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return err.Error()
	}
	//assigning var from body

	fmt.Printf("json map: %v\n", data)

	//name
	name, ok := data["name"].(string)
	if !ok || len(name) == 0 {
		fmt.Println("Name data not found or is not of the expected type")
		return err.Error()
	}
	w.Name = name

	//coordinates
	coord, ok := data["coord"].(map[string]interface{})
	if !ok {
		fmt.Println("coord data not found or is not of the expected type")
		return err.Error()
	}
	lat, ok := coord["lat"].(float64)

	if !ok || lat == 0 {
		fmt.Println("Conversion failed")
		return err.Error()
	}

	lon, ok := coord["lon"].(float64)
	if !ok || lon == 0 {
		fmt.Println("Conversion failed")
		return err.Error()
	}

	w.Coord.Lat = lat
	w.Coord.Lon = lon

	//weather
	weather := data["weather"].([]interface{})
	if !ok || len(weather) == 0 {
		fmt.Println("Weather data not found or is not of the expected type")
		return err.Error()
	}

	// Access the first element of the "weather" array
	firstWeatherData, ok := weather[0].(map[string]interface{})
	if !ok {
		fmt.Println("Weather data not found or is not of the expected type")
		return err.Error()
	}

	id, ok := firstWeatherData["id"].(float64)
	if !ok {
		fmt.Println("id data not found or is not of the expected type")
		return err.Error()
	}

	main, ok := firstWeatherData["main"].(string)
	if !ok || main == "" {
		fmt.Println("main data not found or is not of the expected type")
		return err.Error()
	}

	desc, ok := firstWeatherData["description"].(string)
	if !ok || desc == "" {
		fmt.Println("desc data not found or is not of the expected type")
		return err.Error()
	}

	w.Weather.Id = id
	w.Weather.Main = main
	w.Weather.Description = desc

	//WeatherTemp
	weathertemp, ok := data["main"].(map[string]interface{})
	if !ok {
		fmt.Println("temp data not found or is not of the expected type")
		return err.Error()
	}

	temp, ok := weathertemp["temp"].(float64)
	if !ok || temp == 0 {
		fmt.Println("temp data not found or is not of the expected type")
		return err.Error()
	}

	feels_like, ok := weathertemp["temp"].(float64)
	if !ok || feels_like == 0 {
		fmt.Println("feels like data not found or is not of the expected type")
		return err.Error()
	}

	temp_min, ok := weathertemp["temp"].(float64)
	if !ok || temp_min == 0 {
		fmt.Println("temp min data not found or is not of the expected type")
		return err.Error()
	}

	temp_max, ok := weathertemp["temp"].(float64)
	if !ok || temp_max == 0 {
		fmt.Println("temp max data not found or is not of the expected type")
		return err.Error()
	}

	pressure, ok := weathertemp["temp"].(float64)
	if !ok || pressure == 0 {
		fmt.Println("pressure data not found or is not of the expected type")
		return err.Error()
	}

	humidity, ok := weathertemp["temp"].(float64)
	if !ok || humidity == 0 {
		fmt.Println("humidity data not found or is not of the expected type")
		return err.Error()
	}

	w.WeatherTemp.Temp = kelvinToCelsius(temp)
	w.WeatherTemp.Feels_like = kelvinToCelsius(feels_like)
	w.WeatherTemp.Temp_min = kelvinToCelsius(temp_min)
	w.WeatherTemp.Temp_max = kelvinToCelsius(temp_max)
	w.WeatherTemp.Pressure = kelvinToCelsius(pressure)
	w.WeatherTemp.Humidity = kelvinToCelsius(humidity)

	return ""
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func weatherDataToJson(weather WeatherData) []byte {
	json, err := json.Marshal(weather)
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func (w *WeatherData) GetWeather(city string) string {
	err := w.getWeather(city)
	if err != "" {
		fmt.Println(err)
		return err
	}
	weatherjson := weatherDataToJson(*w)
	return string(weatherjson)
}
