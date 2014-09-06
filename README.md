# WeatherGo

A small library to get weather forecast from [OpenWeatherMap](http://openweathermap.org/).


### Overview

After watching this great video, [Get started with Go](https://www.youtube.com/watch?v=2KmHtgtEZ1s) I came up with 
this. A library to get weather information using a free API.


```
func ByCity(city string) (*Response, error)
```

### Installation

Getting a library with Go is pretty straighforward: 

```
go get github.com/sgmac/weathergo
```
Build the cli application

```
make
```
Run it!!

```
$ cli/cli-weather Madrid
Madrid
======
Lon: -3.700000
Lat: 40.419998
Sunrise:7:46
Sunset:20:40
Main: Clear
Condition: Sky is Clear
Temp: 32.0
Temp Max:32.1 ºC
Temp Min: 31.8
Humidity: 18
```

### Todo

 - Process cities using gorutines
 - Travis CI
 - Testing

### License

MIT

