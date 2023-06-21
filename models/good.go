package models

import (
	"fmt"
	"goSeckill/config"
	"sync"
)

var GM  = NewGood() //初始化good

type Good struct {
	GoodSync sync.WaitGroup
	GoodMutex *sync.Mutex //全局写锁
	GoodRMutex *sync.RWMutex
	GoodChan chan GoodData

}
type GoodData struct {
	Id int `json:"id"`
	GoodName string `json:"good_name"`
	GoodNumber int `json:"good_number"`//商品数量
	GoodPrice float64 `json:"good_price"` //商品价格
}

//实例化商品
func NewGood() *Good {
	return &Good{
		GoodSync: sync.WaitGroup{},
		GoodMutex: &sync.Mutex{},
		GoodRMutex: &sync.RWMutex{},
		GoodChan: make(chan GoodData),
	}
}
//读取商品
func ReadGood() GoodData{
	gs:= getConfigGood()
	g:= GoodData{
		GoodNumber: gs.GoodNumber,
		GoodName: gs.Name,
		GoodPrice: gs.Price,
	}
	//fmt.Println(g)
	return g
}
//购买一个商品

func BuyGoods(){

	//修改商品商量
	GM.GoodMutex.Lock() //加锁
	defer GM.GoodMutex.Unlock() //解锁

	GM.GoodSync.Add(1)
	goods := ReadGood() //读取商品
	//判定用户是否购买--等其它逻辑
	if goods.GoodNumber > 0 {
		buyNum:= goods.GoodNumber -1
		config.SetConfigGoodNumber(buyNum)
		fmt.Println("商品购买成功:",goods)

	}else{
		fmt.Println("商品数量小于0")
	}
	GM.GoodSync.Done()

	GM.GoodSync.Wait()

}

func getConfigGood() config.Goods{
	c :=config.GetConfigGood()
	return c
}
