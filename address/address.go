package address

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chimerakang/goutils/json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type IpResp struct {
	Status   string `json:"status"`
	Province string `json:"province"`
	City     string `json:"city"`
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

// 獲取IP真實地址
func GetIpRealLocation(ip string) string {
	resp, err := http.Get(fmt.Sprintf("https://restapi.amap.com/v3/ip?ip=%s&key=%s", ip, "9130aaac2b7a920b8bbd5dc9647fbe9e"))
	address := "unknown address"
	if err != nil {
		log.Error().Err(err).Msg("[GetIpRealLocation]IP ADDRESS 查詢失敗")
		return address
	}
	defer resp.Body.Close()
	// 读取响应数据
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("[GetIpRealLocation]IP地址查询失败")
		return address
	}
	// json数据转结构体
	var result IpResp
	json.Json2Struct(string(data), &result)
	if result.Status == "1" {
		address = result.Province
		if result.City != "" && result.Province != result.City {
			address += result.City
		}
	}
	return address
}
