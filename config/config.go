package config

import (

	"fmt"
	"github.com/spf13/viper"

)

type Goods struct {
	GoodNumber int
	Name string
	Price float64
}

var vip *viper.Viper
var gs Goods

func InitConfig()  {

	// 读取YAML配置文件
	v := viper.New()

	v.SetConfigFile("config/config.yaml")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to read config file:", err)
	}
	vip = v

}

func GetConfigGood() Goods {
	gs = Goods{
		GoodNumber: vip.GetInt("good.goodnumber"),
		Name: vip.GetString("good.name"),
		Price: vip.GetFloat64("good.price"),
	}
 	return gs
}

//修改配置的值
func SetConfigGoodNumber(number int) {

	vip.Set("good.goodNumber", number)
//	fmt.Println("number:",vip.GetInt("good.goodnumber"))
	//// 保存修改后的配置到文件
	err := vip.WriteConfig()
	if err != nil {
		fmt.Println("Failed to write config file: %v", err)
	}
}

