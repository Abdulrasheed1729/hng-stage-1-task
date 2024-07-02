package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type APIServer struct {
	listenAddr string
}

var (
	ipToLocationAPIKey = os.Getenv("IP_TO_LOCATION_API_KEY")
	weatherAPIKey      = os.Getenv("WEATHER_API_KEY")
)

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	log.Println("JSON API server running on port: ", s.listenAddr)

	router.HandleFunc("/hello", makeHTTPHandleFunc(s.handleHello))

	http.ListenAndServe(s.listenAddr, router)
}

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

	var ip2locationResponse IPToLocationResponse

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

	var weatherAPIResponse WeatherAPIResponse

	err = json.Unmarshal(body, &weatherAPIResponse)

	if err != nil {
		return 0.0, err
	}

	return weatherAPIResponse.Current.TempC, nil
}

func (s *APIServer) handleHello(w http.ResponseWriter, r *http.Request) error {
	clientIp, _, err := net.SplitHostPort(r.RemoteAddr)

	if err != nil {
		return err
	}

	cityName, err := getLocation(clientIp)

	if err != nil {
		return err
	}

	temperature, err := getWeather(cityName)

	if err != nil {
		return err
	}

	visitorName := r.URL.Query().Get("visitor_name")

	if visitorName == "" {
		visitorName = "Guest"
	}

	if cityName == "-" {
		cityName = "localhost"
	}

	var greeting string

	var resp Response

	if clientIp == "127.0.0.1" {
		greeting = fmt.Sprintf("Hello, %s!, the temperature is %0.f degrees Celsius at localhost.", visitorName, 21.0)
		resp = Response{
			ClientIp: clientIp,
			Location: cityName,
			Greeting: greeting,
		}

		return WriteJSON(w, http.StatusOK, resp)
	}

	greeting = fmt.Sprintf("Hello, %s!, the temperature is %0.f degrees Celsius in %s", visitorName, temperature, cityName)

	resp = Response{
		ClientIp: clientIp,
		Location: cityName,
		Greeting: greeting,
	}

	return WriteJSON(w, http.StatusOK, resp)

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
