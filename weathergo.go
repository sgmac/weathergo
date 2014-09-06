// Package to get weather forecast from OpenWeatherMap API
package weathergo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/sgmac/weathergo/helpers"
)

// JSON response
type Response struct {
	Coord struct {
		Lon float32
		Lat float32
	}

	Sys struct {
		Type    int
		Id      int
		Message float32
		Sunrise int64
		Sunset  int64
	}

	Weather []struct {
		Id          int
		Main        string
		Description string
		Icon        string
	}

	Base string

	Main struct {
		Temp     float32
		Pressure float32
		Humidity int
		Temp_min float32
		Temp_max float32
	}

	Wind struct {
		Speed  float32
		Deg    float32
		VarBeg int `json:"var_beg"`
		VarEnd int `json:"var_end"`
	}

	Clouds struct {
		All int
	}

	Rain struct {
		Dt int32
	}

	Dt   int32
	Id   int
	Name string
	Cod  int
}

func (r *Response) String() string {

	header := helpers.DoLine(r.Name)
	coords := "Lon: %f\nLat: %f\n"
	weather := "Main: %s\nCondition: %s\n"
	main := "Temp: %.1f\nTemp Max:%.1f ºC\nTemp Min: %.1f\nHumidity: %d"
	sunrise := "Sunrise:%s\n"
	sunset := "Sunset:%s\n"

	s := []string{header, coords, sunrise, sunset, weather, main}
	ret := strings.Join(s, "")

	return fmt.Sprintf(ret, r.Name,
		r.Coord.Lon, r.Coord.Lat,
		helpers.EpochToTime(r.Sys.Sunrise), helpers.EpochToTime(r.Sys.Sunset),
		r.Weather[0].Main, r.Weather[0].Description,
		(r.Main.Temp - 273), (r.Main.Temp_max - 273), (r.Main.Temp_min - 273), r.Main.Humidity)

}

// Returns the weather forecast by city name
func ByCity(city string) (*Response, error) {

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s", city)
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	to_json := new(Response)

	err = json.NewDecoder(response.Body).Decode(to_json)

	if err != nil {
		return nil, err
	}

	return to_json, nil

}
