package http

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

const (
	quoteUrl       = "https://stock.xueqiu.com/v5/stock/batch/quote.json?symbol=SH000001,SZ399001,SZ399006,SH000688,HKHSI,HKHSCEI,HKHSCCI,.DJI,.IXIC,.INX"
	hotStockUrl    = "https://stock.xueqiu.com/v5/stock/hot_stock/list.json?size=8&_type=10&type=10"
	snbArticleUrl  = "https://xueqiu.com/statuses/hot/listV2.json?since_id=-1&max_id=-1&size=15"
	newsArticleUrl = "https://xueqiu.com/statuses/hot/listV2.json?since_id=-1&max_id=-1&size=15"
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
