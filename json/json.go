package json

import (
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

// 結構體轉為json
func Struct2Json(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		log.Debug().Err(err).Msg("[Struct2Json]轉換異常")
	}
	return string(str)
}

// json轉為結構體
func Json2Struct(str string, obj interface{}) {
	// 將json轉為結構體
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		log.Debug().Err(err).Msg("[Json2Struct]轉換異常")
	}
}

// json interface轉為結構體
func JsonI2Struct(str interface{}, obj interface{}) {
	// 將json interface轉為string
	jsonStr, _ := str.(string)
	Json2Struct(jsonStr, obj)
}

// 結構體轉結構體, json為中間橋樑, struct2必須以指針方式傳遞, 否則可能獲取到空數據
func Struct2StructByJson(struct1 interface{}, struct2 interface{}) {
	// 轉換為響應結構體, 隱藏部分字段
	jsonStr := Struct2Json(struct1)
	Json2Struct(jsonStr, struct2)
}

// 兩結構體比對不同的字段, 不同時將取struct1中的字段返回, json為中間橋樑, struct3必須以指針方式傳遞, 否則可能獲取到空數據
func CompareDifferenceStructByJson(oldStruct interface{}, newStruct interface{}, update interface{}) {
	// 通過json先將其轉為map集合
	m1 := make(gin.H, 0)
	m2 := make(gin.H, 0)
	m3 := make(gin.H, 0)
	Struct2StructByJson(newStruct, &m1)
	Struct2StructByJson(oldStruct, &m2)
	for k1, v1 := range m1 {
		for k2, v2 := range m2 {
			switch v1.(type) {
			// 複雜結構不做對比
			case map[string]interface{}:
				continue
			}
			// key相同, 值不同
			if k1 == k2 && v1 != v2 {
				m3[k1] = v1
				break
			}
		}
	}
	Struct2StructByJson(m3, &update)
}
