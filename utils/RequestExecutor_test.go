package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_sssss(t *testing.T) {

	var dataStr string = `
{
	"data": [{
		"uid": 66455,
		"username": "MichaelKong",
		"email": "kongyouji@gmail.com",
		"open": 0,
		"name": "",
		"avatar": "gopher20.png",
		"city": "",
		"company": "",
		"github": "",
		"gitea": "",
		"weibo": "",
		"website": "",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 0,
		"is_third": 0,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 0,
		"is_root": false,
		"ctime": "2021-06-10 16:20:52",
		"mtime": "2021-06-10T16:20:52+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 0,
		"copper": 0,
		"is_online": false
	}, {
		"uid": 51156,
		"username": "wuwentao",
		"email": "wuwentao.1024@gmail.com",
		"open": 0,
		"name": "",
		"avatar": "gopher22.png",
		"city": "",
		"company": "",
		"github": "",
		"gitea": "",
		"weibo": "",
		"website": "",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 0,
		"is_third": 0,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 1,
		"is_root": false,
		"ctime": "2021-06-10 15:44:36",
		"mtime": "2021-06-10T20:44:35+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 0,
		"copper": 0,
		"is_online": false
	}, {
		"uid": 66454,
		"username": "gogogo0610",
		"email": "cnfreebsd@qq.com",
		"open": 0,
		"name": "",
		"avatar": "gopher21.png",
		"city": "",
		"company": "",
		"github": "",
		"gitea": "",
		"weibo": "",
		"website": "",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 0,
		"is_third": 0,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 1,
		"is_root": false,
		"ctime": "2021-06-10 14:50:48",
		"mtime": "2021-06-10T14:52:44+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 0,
		"copper": 0,
		"is_online": false
	}, {
		"uid": 66453,
		"username": "MrLiuhhh",
		"email": "MrLiuhhh@github.com",
		"open": 0,
		"name": "",
		"avatar": "https://avatars.githubusercontent.com/u/20110182?v=4",
		"city": "",
		"company": "",
		"github": "MrLiuhhh",
		"gitea": "",
		"weibo": "",
		"website": "",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 0,
		"is_third": 1,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 1,
		"is_root": false,
		"ctime": "2021-06-10 14:24:09",
		"mtime": "2021-06-10T14:24:09+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 0,
		"copper": 0,
		"is_online": false
	}, {
		"uid": 66457,
		"username": "peate",
		"email": "356522375@qq.com",
		"open": 1,
		"name": "peateDeng",
		"avatar": "gopher07.png",
		"city": "广东广州",
		"company": "",
		"github": "",
		"gitea": "",
		"weibo": "",
		"website": "",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 0,
		"is_third": 0,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 1,
		"is_root": false,
		"ctime": "2021-06-10 13:54:19",
		"mtime": "2021-06-10T18:54:52+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 0,
		"copper": 0,
		"is_online": false
	}, {
		"uid": 66456,
		"username": "smile_yfc",
		"email": "yao19981001@163.com",
		"open": 0,
		"name": "",
		"avatar": "gopher16.png",
		"city": "",
		"company": "",
		"github": "",
		"gitea": "",
		"weibo": "",
		"website": "",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 2000,
		"is_third": 0,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 1,
		"is_root": false,
		"ctime": "2021-06-10 12:45:45",
		"mtime": "2021-06-10T17:46:43+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 20,
		"copper": 0,
		"is_online": false
	}, {
		"uid": 66452,
		"username": "fastzhong",
		"email": "fastzhong@github.com",
		"open": 0,
		"name": "Alan",
		"avatar": "gopher_aqua.jpg",
		"city": "Singapore",
		"company": "352app",
		"github": "fastzhong",
		"gitea": "",
		"weibo": "",
		"website": "https://fastzhong.com",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 0,
		"is_third": 1,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 1,
		"is_root": false,
		"ctime": "2021-06-10 08:17:54",
		"mtime": "2021-06-10T13:18:47+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 0,
		"copper": 0,
		"is_online": false
	}, {
		"uid": 66451,
		"username": "rongqinzheng",
		"email": "rongqinzheng@github.com",
		"open": 0,
		"name": "",
		"avatar": "https://avatars.githubusercontent.com/u/25027301?v=4",
		"city": "",
		"company": "",
		"github": "rongqinzheng",
		"gitea": "",
		"weibo": "",
		"website": "",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 0,
		"is_third": 1,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 1,
		"is_root": false,
		"ctime": "2021-06-10 04:38:20",
		"mtime": "2021-06-10T09:38:19+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 0,
		"copper": 0,
		"is_online": false
	}, {
		"uid": 66450,
		"username": "moogoxu",
		"email": "moogoxu@163.com",
		"open": 0,
		"name": "",
		"avatar": "gopher23.png",
		"city": "",
		"company": "",
		"github": "",
		"gitea": "",
		"weibo": "",
		"website": "",
		"monlog": "",
		"introduce": "",
		"unsubscribe": 0,
		"balance": 2000,
		"is_third": 0,
		"dau_auth": 87,
		"is_vip": false,
		"vip_expire": 0,
		"status": 1,
		"is_root": false,
		"ctime": "2021-06-10 04:07:38",
		"mtime": "2021-06-10T09:11:47+08:00",
		"Roleids": null,
		"Rolenames": null,
		"weight": 0,
		"gold": 0,
		"silver": 20,
		"copper": 0,
		"is_online": false
	}],
	"msg": "操作成功",
	"ok": 1
}
`
	res, err := JsonPathExtract(dataStr, "$.data[0].uid")
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

	RequestExecutor(&ApiCase{
		Id:     1,
		Name:   "百度",
		Query:  "?limit=${id}",
		Domain: "https://studygolang.com/",
		Url:    "users/newest",
		Method: "GET",
		GVars:  data,
		Asserts: []Asserts{
			{AssertType: "response_json", Check: "$.data[0].uid", Expect: "66455", Comparator: "相等"},
			{AssertType: "response_json", Check: "$.data[0].uid", Expect: "66469", Comparator: "相等"},
			{AssertType: "status_code", Check: "", Expect: "200", Comparator: "相等"},
		},
	})
	//RequestExecutor(&ApiCase{
	//	Id:         1,
	//	Name:       "百度",
	//	Domain:     "http://127.0.0.1:3000/",
	//	Url:        "api/v1/login",
	//	Method:     "POST",
	//	Type:       "form",
	//	Parameters: `{"user_name":"admin", "password":"${password}"}`,
	//	GVars:      data,
	//})
}
