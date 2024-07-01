package handler

import "time"

type Response struct {
	ClientIp string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

type IPToLocationResponse struct {
	IP                 string  `json:"ip"`
	CountryCode        string  `json:"country_code"`
	CountryName        string  `json:"country_name"`
	RegionName         string  `json:"region_name"`
	District           string  `json:"district"`
	CityName           string  `json:"city_name"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	ZipCode            string  `json:"zip_code"`
	TimeZone           string  `json:"time_zone"`
	Asn                string  `json:"asn"`
	As                 string  `json:"as"`
	Isp                string  `json:"isp"`
	Domain             string  `json:"domain"`
	NetSpeed           string  `json:"net_speed"`
	IddCode            string  `json:"idd_code"`
	AreaCode           string  `json:"area_code"`
	WeatherStationCode string  `json:"weather_station_code"`
	WeatherStationName string  `json:"weather_station_name"`
	Mcc                string  `json:"mcc"`
	Mnc                string  `json:"mnc"`
	MobileBrand        string  `json:"mobile_brand"`
	Elevation          int     `json:"elevation"`
	UsageType          string  `json:"usage_type"`
	AddressType        string  `json:"address_type"`
	AdsCategory        string  `json:"ads_category"`
	AdsCategoryName    string  `json:"ads_category_name"`
	Continent          struct {
		Name        string   `json:"name"`
		Code        string   `json:"code"`
		Hemisphere  []string `json:"hemisphere"`
		Translation struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"translation"`
	} `json:"continent"`
	Country struct {
		Name        string `json:"name"`
		Alpha3Code  string `json:"alpha3_code"`
		NumericCode int    `json:"numeric_code"`
		Demonym     string `json:"demonym"`
		Flag        string `json:"flag"`
		Capital     string `json:"capital"`
		TotalArea   int    `json:"total_area"`
		Population  int    `json:"population"`
		Currency    struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"currency"`
		Language struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"language"`
		Tld         string `json:"tld"`
		Translation struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"translation"`
	} `json:"country"`
	Region struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		Translation struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"translation"`
	} `json:"region"`
	City struct {
		Name        string `json:"name"`
		Translation struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"translation"`
	} `json:"city"`
	TimeZoneInfo struct {
		Olson       string    `json:"olson"`
		CurrentTime time.Time `json:"current_time"`
		GmtOffset   int       `json:"gmt_offset"`
		IsDst       bool      `json:"is_dst"`
		Sunrise     string    `json:"sunrise"`
		Sunset      string    `json:"sunset"`
	} `json:"time_zone_info"`
	Geotargeting struct {
		Metro any `json:"metro"`
	} `json:"geotargeting"`
	IsProxy bool `json:"is_proxy"`
	Proxy   struct {
		LastSeen                   int    `json:"last_seen"`
		ProxyType                  string `json:"proxy_type"`
		Threat                     string `json:"threat"`
		Provider                   string `json:"provider"`
		IsVpn                      bool   `json:"is_vpn"`
		IsTor                      bool   `json:"is_tor"`
		IsDataCenter               bool   `json:"is_data_center"`
		IsPublicProxy              bool   `json:"is_public_proxy"`
		IsWebProxy                 bool   `json:"is_web_proxy"`
		IsWebCrawler               bool   `json:"is_web_crawler"`
		IsResidentialProxy         bool   `json:"is_residential_proxy"`
		IsConsumerPrivacyNetwork   bool   `json:"is_consumer_privacy_network"`
		IsEnterprisePrivateNetwork bool   `json:"is_enterprise_private_network"`
		IsSpammer                  bool   `json:"is_spammer"`
		IsScanner                  bool   `json:"is_scanner"`
		IsBotnet                   bool   `json:"is_botnet"`
	} `json:"proxy"`
}

type WeatherAPIResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}
