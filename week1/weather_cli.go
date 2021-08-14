package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	jsoniter "github.com/json-iterator/go"
)

func GetWeather(citycode string) string {
	type Weatherinfo struct {
		Sd      string `json:"SD"`
		Wd      string `json:"WD"`
		Ws      string `json:"WS"`
		Wse     string `json:"WSE"`
		City    string `json:"city"`
		Cityid  string `json:"cityid"`
		Isradar string `json:"isRadar"`
		Radar   string `json:"radar"`
		Temp    string `json:"temp"`
		Time    string `json:"time"`
		Weather string `json:"weather"`
	}
	type Autogenerated struct {
		Weatherinfo Weatherinfo `json:"weatherinfo"`
	}
	var QueryUrl string = fmt.Sprintf(`http://weatherapi.market.xiaomi.com/wtr-v2/temp/realtime?cityId=%s`, citycode)
	client := &http.Client{}
	req, err := http.NewRequest("GET", QueryUrl, nil)
	if err != nil {
		fmt.Println("Fatal Error occured:", err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Fatal Error occured:", err.Error())
	}
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var WeatherData Autogenerated
	json.Unmarshal(bodyByte, &WeatherData)
	wstrmodel := `当前时间: %s/当前城市: %s /%s %s`
	return fmt.Sprintf(wstrmodel, WeatherData.Weatherinfo.Time, WeatherData.Weatherinfo.City, WeatherData.Weatherinfo.Temp, WeatherData.Weatherinfo.Weather)
}
func GetPos() string {
	type Data struct {
		CountryCode string `json:"countryCode"`
		Country     string `json:"country"`
		Province    string `json:"province"`
		City        string `json:"city"`
		IP          string `json:"ip"`
		Latitude    string `json:"latitude"`
		Longitude   string `json:"longitude"`
		Zipcode     string `json:"zipcode"`
		Timezone    string `json:"timezone"`
		Refer       string `json:"refer"`
	}
	type IPAddr struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	var uri string = `https://www.douyacun.com/api/openapi/geo/location`
	var atoken string = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SWQiOiJlZWQ4ZmQ1ODBmYTRmNjkyIn0.d7qF_mjdXMC0R5M6f04Lnh6x61kaU4lqHT0Axt9xUOY`
	var quri string = fmt.Sprintf("%s?token=%s", uri, atoken)
	var dict map[string]string = makeCityDict("citycode.json")
	client := &http.Client{}
	req, err := http.NewRequest("GET", quri, nil)
	if err != nil {
		fmt.Println("Fatal Error occured:", err.Error())
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Fatal Error occured:", err.Error())
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var ipaddr IPAddr
	json.Unmarshal(bodyByte, &ipaddr)
	if (err == nil) && (ipaddr.Message == "success") {
		println(ipaddr.Data.City)
		return dict[ipaddr.Data.City]
	} else {
		return "error"
	}
}
func makeCityDict(fname string) map[string]string {
	type CityCode []struct {
		ID       int    `json:"id"`
		Pid      int    `json:"pid"`
		CityCode string `json:"city_code"`
		CityName string `json:"city_name"`
		PostCode string `json:"post_code"`
		AreaCode string `json:"area_code"`
		Ctime    string `json:"ctime"`
	}
	jsonfile, err := os.Open(fname)
	if err != nil {
		fmt.Println("Fatal Error occured in reading city code:", err.Error())
	}
	defer jsonfile.Close()
	byteValue, _ := ioutil.ReadFile(fname)
	var cname CityCode
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	retcc := make(map[string]string)
	err = json.Unmarshal(byteValue, &cname)
	for _, cval := range cname {
		retcc[cval.CityName] = cval.CityCode
	}
	return retcc
}
