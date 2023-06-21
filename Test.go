package main

import (

	"goSeckill/config"
 	"goSeckill/models"
	"sync"
)

func init(){

	config.InitConfig()

}

func main(){
	var wg = sync.WaitGroup{}

	concurrency := 102 //模拟请求
	wg.Add(concurrency)
	for i:=0;i<concurrency;i++ {
		go func() {
			models.BuyGoods()
			wg.Done()
		}()
	}
	wg.Wait()

}


