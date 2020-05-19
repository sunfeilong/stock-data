package main

import (
	"../enums"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	stockExchange   string
	companyInfoUrl  string
	stockInfoUrl    string
	realTimeInfoUrl string
}

type Configs struct {
	configs string
}

func getConfigStockExchange(sc enums.StockExchange) (Configs, error) {
	c := Configs{}
	fileName := "configs.yml"
	yamlFIle, eil := ioutil.ReadFile(fileName)
	if nil != eil {
		log.Fatal("读取配置文件出错,fileName: %s, error info: %s", fileName, eil.Error())
		return c, eil
	}
	eil = yaml.Unmarshal(yamlFIle, &c)
	if nil != eil {
		log.Fatal("解析配置文件出错,fileName: %s, error info: %s", fileName, eil.Error())
		return c, eil
	}
	return c, nil
}

func main() {
	configs, eil := getConfigStockExchange(enums.SZ)
	if nil != eil {
		log.Fatal("获取配置信息失败,%s", eil.Error())
	}
	fmt.Println(configs)
}
