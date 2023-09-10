package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

var HTTP_SERVER_ADDRESS string = "localhost"
var HTTP_SERVER_PORT string = "8889"
var HTTP_SERVER_URI string = HTTP_SERVER_ADDRESS + ":" + HTTP_SERVER_PORT

var ROUTE_BOOKS string = "/books"
var ROUTE_BOOK string = "/book"

//gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local
var SQL_USERNAME string = "myuser"
var SQL_PASSWORD string = "mypassword"
var SQL_ADDR string = "127.0.0.1"
var SQL_CONNECTION string = SQL_USERNAME + ":" + SQL_PASSWORD + "@tcp(" + SQL_ADDR + ")/mydb?charset=utf8&parseTime=True&loc=Local"

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		println(body)
		if err := json.Unmarshal([]byte(body), x); err == nil {
			return
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
