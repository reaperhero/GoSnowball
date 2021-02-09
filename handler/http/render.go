package http

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	quoteUrl       = "https://stock.xueqiu.com/v5/stock/batch/quote.json?symbol=SH000001,SZ399001,SZ399006,SH000688,HKHSI,HKHSCEI,HKHSCCI,.DJI,.IXIC,.INX"
	hotStockUrl    = "https://stock.xueqiu.com/v5/stock/hot_stock/list.json?size=8&_type=10&type=10"
	snbArticleUrl  = "https://xueqiu.com/statuses/hot/listV2.json?since_id=-1&max_id=-1&size=15"
	newsArticleUrl = "https://xueqiu.com/statuses/livenews/list.json?since_id=-1&max_id=-1&count=15"
	industriesUrl  = "https://xueqiu.com/service/screener/industries?category=CN&_=1612874545440"
	areaUrl        = "https://xueqiu.com/service/screener/areas?_=1612874545438"
	toolUrl        = "https://xueqiu.com/hq/screener"
)

func (h *handler) index(ctx echo.Context) error {
	req, _ := http.NewRequest("GET", quoteUrl, nil)
	setHeader(req)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ctx.JSON(200, "参数错误")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	data := new(interface{})
	json.Unmarshal(body, &data)
	return ctx.JSON(200, data)
}

func (h *handler) hotStock(ctx echo.Context) error {
	req, _ := http.NewRequest("GET", hotStockUrl, nil)
	setHeader(req)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ctx.JSON(200, "参数错误")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	data := new(interface{})
	json.Unmarshal(body, &data)
	return ctx.JSON(200, data)
}

func (h *handler) news(ctx echo.Context) error {
	name := ctx.QueryParam("category")
	if name == "recommand" {
		req, _ := http.NewRequest("GET", snbArticleUrl, nil)
		setHeader(req)
		resp, err := (&http.Client{}).Do(req)
		if err != nil {
			return ctx.JSON(200, "参数错误")
		}
		body, _ := ioutil.ReadAll(resp.Body)
		data := new(interface{})
		json.Unmarshal(body, &data)
		return ctx.JSON(200, data)
	}
	if name == "dayinfo" {
		req, _ := http.NewRequest("GET", newsArticleUrl, nil)
		setHeader(req)
		resp, err := (&http.Client{}).Do(req)
		if err != nil {
			return ctx.JSON(200, "参数错误")
		}
		body, _ := ioutil.ReadAll(resp.Body)
		data := new(interface{})
		json.Unmarshal(body, &data)
		return ctx.JSON(200, data)
	}
	return nil
}

func (h *handler) industries(ctx echo.Context) error {
	req, _ := http.NewRequest("GET", industriesUrl, nil)
	setHeader(req)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ctx.JSON(200, "参数错误")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	data := new(interface{})
	json.Unmarshal(body, &data)
	return ctx.JSON(200, data)
}

func (h *handler) chooseAreas(ctx echo.Context) error {
	req, _ := http.NewRequest("GET", areaUrl, nil)
	setHeader(req)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ctx.JSON(200, "参数错误")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	data := new(interface{})
	json.Unmarshal(body, &data)
	return ctx.JSON(200, data)
}

func (h *handler) chooseTools(ctx echo.Context) error {
	req, _ := http.NewRequest("GET", toolUrl, nil)
	setHeader(req)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ctx.JSON(200, "参数错误")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	reg1 := regexp.MustCompile(`SNB.data.condition =(?s:(.*?));`)

	result1 := reg1.FindAllStringSubmatch(string(body), -1)
	return ctx.String(200, result1[0][1])
}
