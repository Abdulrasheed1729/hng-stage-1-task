package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	ht "github.com/Abdulrasheed1729/hng-stage-1/types"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ipToLocationAPIKey = os.Getenv("IP_TO_LOCATION_API_KEY")
	weatherAPIKey      = os.Getenv("WEATHER_API_KEY")
)

func getLocation(ip string) (string, error) {

	url := fmt.Sprintf("https://api.ip2location.io/?key=%s&ip=%s", ipToLocationAPIKey, ip)

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var ip2locationResponse ht.IPToLocationResponse

	err = json.Unmarshal(body, &ip2locationResponse)

	if err != nil {
		return "", err
	}

	return ip2locationResponse.CityName, nil

}

func getWeather(city string) (float64, error) {

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", weatherAPIKey, city)

	resp, err := http.Get(url)

	if err != nil {
		return 0.0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return 0.0, err
	}

	var weatherAPIResponse ht.WeatherAPIResponse

	err = json.Unmarshal(body, &weatherAPIResponse)

	if err != nil {
		return 0.0, err
	}

	return weatherAPIResponse.Current.TempC, nil
}

func Hello(w http.ResponseWriter, r *http.Request) {
	clientIp, _, err := net.SplitHostPort(r.RemoteAddr)

	if err != nil {
		log.Fatal(err)
	}

	cityName, err := getLocation(clientIp)

	if err != nil {
		log.Fatal(err)

	}

	temperature, err := getWeather(cityName)

	if err != nil {
		log.Fatal(err)

	}

	visitorName := r.URL.Query().Get("visitor_name")

	if visitorName == "" {
		visitorName = "Guest"
	}

	var greeting string

	var resp ht.Response

	if clientIp == "127.0.0.1" {
		greeting = fmt.Sprintf("Hello, %s!, the temperature is %0.f degrees Celsius at home.", visitorName, 21.0)
		resp = ht.Response{
			ClientIp: clientIp,
			Location: cityName,
			Greeting: greeting,
		}

		err = WriteJSON(w, http.StatusOK, resp)

		if err != nil {
			log.Fatal(err)
		}
	}

	greeting = fmt.Sprintf("Hello, %s!, the temperature is %0.f degrees Celsius in %s", visitorName, temperature, cityName)

	resp = ht.Response{
		ClientIp: clientIp,
		Location: cityName,
		Greeting: greeting,
	}

	err = WriteJSON(w, http.StatusOK, resp)

	if err != nil {
		log.Fatal(err)
	}

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
