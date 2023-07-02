package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
)
type QueryMap map[string]string
type ParamsMap map[string]string

type URLData struct {
	Params ParamsMap
	Query  QueryMap
}

type HandlerWithBodyFunc[T interface{}] func(ctx context.Context, urlData URLData, body T) (code int, obj interface{})

func ParamsAsMap(c *gin.Context) ParamsMap {
	res := make(ParamsMap)
	for _, p := range c.Params {
		res[p.Key] = p.Value
	}
	return res
}

func QueryParamsAsMap(c *gin.Context) QueryMap {
	res := make(QueryMap)
	queryMap := c.Request.URL.Query()
	for k, v := range queryMap {
		if len(v) > 0 {
			res[k] = v[0]
		}
	}
	return res
}


func NewHandlerWithBody[T interface{}](handler HandlerWithBodyFunc[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		urlData := URLData{
			Params: ParamsAsMap(c),
			Query:  QueryParamsAsMap(c),
		}
		var body T
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{
				"message": "invalid body",
			})
			return
		}
		code, obj := handler(c, urlData, body)
		c.JSON(code, obj)
	}
}