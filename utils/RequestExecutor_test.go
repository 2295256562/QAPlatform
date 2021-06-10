package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_sssss(t *testing.T) {

	var dataStr string = `
{
    "store": {
        "book": [
            {
                "category": "reference",
                "author": "Nigel Rees",
                "title": "Sayings of the Century",
                "price": 8.95
            },
            {
                "category": "fiction",
                "author": "Evelyn Waugh",
                "title": "Sword of Honour",
                "price": 12.99
            },
            {
                "category": "fiction",
                "author": "Herman Melville",
                "title": "Moby Dick",
                "isbn": "0-553-21311-3",
                "price": 8.99
            },
            {
                "category": "fiction",
                "author": "J. R. R. Tolkien",
                "title": "The Lord of the Rings",
                "isbn": "0-395-19395-8",
                "price": 22.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    },
    "expensive": 10
}
`
	res, err := JsonPathExtract(dataStr, "$.store.book[?(@.price < $.expensive)].price")
	fmt.Println("结果 = ", res, "err = ", err)
	return

}

func Test_replaceKeyFromMap(t *testing.T) {
	str := "name=${name}&page=${page}&dd=${flag}&price=${price}"
	url.Values{}.Encode()
	data := make(map[string]interface{})
	data["name"] = "张三"
	data["page"] = 1
	data["flag"] = true
	data["price"] = 3.14
	replaceKeyFromMap(str, data)
}

func Test_Case(t *testing.T) {
	data := make(map[string]interface{})
	data["id"] = 9
	data["password"] = "123456"

	//RequestExecutor(&ApiCase{
	//	Id: 1,
	//	Name: "百度",
	//	Query: "?limit=${id}",
	//	Domain: "https://studygolang.com/",
	//	Url: "users/newest",
	//	Method: "GET",
	//	GVars: data,
	//})
	RequestExecutor(&ApiCase{
		Id:         1,
		Name:       "百度",
		Domain:     "http://127.0.0.1:3000/",
		Url:        "api/v1/login",
		Method:     "POST",
		Type:       "form",
		Parameters: `{"user_name":"admin", "password":"${password}"}`,
		GVars:      data,
	})
}
