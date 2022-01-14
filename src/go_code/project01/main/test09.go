package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
)

type names struct {
	Data struct {
		Item []struct {
			Id         int    `json:"id"`
			Title      string `json:"title"`
			Content    string `json:"content"`
			CreateTime int    `json:"create_time"`
			Outline    string `json:"outline"`
		} `json:"item"`
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main()  {
	data := []byte(`{
    "data": {
        "item": [
            {
                "id": 72,
                "title": "2019/6/1",
                "content": "<p>系统公告尊敬的V社群粉丝们：                   </p><p>       大家好，为了大家反馈问题能得到及时的回复和处理，近段时间客服部门在调试智能化问题处理系统，导致一些反馈问题拖延了时间，平台在此向大家致以诚挚的歉意！经过紧张的调试和运行，目前已经能满足大家的反应问题的快速反馈，请粉丝们有问题及时反馈客服！但是由于前期积压问题过多，需要技术部门一一处理，处理时间会尽量从反馈之日起一个月内处理好，谢谢您的理解！      平台充提币功能经过持续的优化，已经陆续开通充提币，但是优化过程还在持续，平台将逐步实现内部和外部主流币种的充提！                                                </p><p>V-token客服部   </p><p>2019.5.28</p><p><br></p>",
                "create_time": 1559375867,
                "outline": "大家好，为了大家反馈问题能得到及时的回复和处理，近段时间客服部门在调试智能化问题处理系统，导致一些反馈问题拖延了时间，平台在此向大家致以诚挚的歉意！"
            },
            {
                "id": 71,
                "title": "2019/6/1",
                "content": "<p>시스템에 공고한 V 단체 팬 여러분:</p><p>        안녕하세요, 여러분의 피드백 문제를 위해 제때에 답변과 처리를 받을 수 있습니다. 근간 객복 부문은 지능화 문제 처리 시스템을 조율하고 있습니다. 일부 피드백 문제를 지연시켰습니다. 플랫폼은 여러분께 진심으로 사과드립니다!긴장된 디버깅과 운행을 거쳐 현재 여러분의 반응 문제의 빠른 피드백을 충족시킬 수 있으니, 팬들에게 질문이 있으면 고객복을 즉시 반영해 주십시오!하지만 전기 쌓인 문제가 너무 많아서 기술부문이 일일이 처리해야 할 때, 처리시간은 가능한 한 한 달 이내에 처리될 것입니다. 이해해 주셔서 감사합니다!플랫폼 충재통화 기능은 지속적인 최적화를 거쳐 이미 속속 충제폐를 개통했으나 최적화 과정은 계속되고 있으며 플랫폼은 내부와 외부 주류 종류의 충제로 이뤄질 것이다.</p><p>V -token 고객센터</p><p>2019.5.28</p><p><br></p>",
                "create_time": 1559375854,
                "outline": "大家好，为了大家反馈问题能得到及时的回复和处理，近段时间客服部门在调试智能化问题处理系统，导致一些反馈问题拖延了时间，平台在此向大家致以诚挚的歉意！"
            }
        ],
        "total": 51,
        "count": 10,
        "per_page": 10,
        "current_page": 1,
        "total_pages": 6,
        "links": {
            "next": "http://3.13.139.54:8883/api/notice?page=2"
        }
    },
    "code": 200,
    "msg": "success"
}`)


	//动态json 可以使用第三方包操作
	//s,_:=jsonparser.GetString(data, "data", "declarant", "[1]")
	//fmt.Println(string(s))

	var sdasd []string
	 jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		sss,_,_,_ :=jsonparser.Get(value,"title")
		 sdasd = append(sdasd,string(sss))
	}, "data", "item")
	fmt.Println(sdasd)


	//使用结构体反序列化
	var name names
	ste2 := "{\n    \"data\": {\n        \"item\": [\n            {\n                \"id\": 72,\n                \"title\": \"2019/6/1\",\n                \"content\": \"<p>系统公告尊敬的V社群粉丝们：                   </p><p>       大家好，为了大家反馈问题能得到及时的回复和处理，近段时间客服部门在调试智能化问题处理系统，导致一些反馈问题拖延了时间，平台在此向大家致以诚挚的歉意！经过紧张的调试和运行，目前已经能满足大家的反应问题的快速反馈，请粉丝们有问题及时反馈客服！但是由于前期积压问题过多，需要技术部门一一处理，处理时间会尽量从反馈之日起一个月内处理好，谢谢您的理解！      平台充提币功能经过持续的优化，已经陆续开通充提币，但是优化过程还在持续，平台将逐步实现内部和外部主流币种的充提！                                                </p><p>V-token客服部   </p><p>2019.5.28</p><p><br></p>\",\n                \"create_time\": 1559375867,\n                \"outline\": \"大家好，为了大家反馈问题能得到及时的回复和处理，近段时间客服部门在调试智能化问题处理系统，导致一些反馈问题拖延了时间，平台在此向大家致以诚挚的歉意！\"\n            },\n            {\n                \"id\": 71,\n                \"title\": \"2019/6/1\",\n                \"content\": \"<p>시스템에 공고한 V 단체 팬 여러분:</p><p>        안녕하세요, 여러분의 피드백 문제를 위해 제때에 답변과 처리를 받을 수 있습니다. 근간 객복 부문은 지능화 문제 처리 시스템을 조율하고 있습니다. 일부 피드백 문제를 지연시켰습니다. 플랫폼은 여러분께 진심으로 사과드립니다!긴장된 디버깅과 운행을 거쳐 현재 여러분의 반응 문제의 빠른 피드백을 충족시킬 수 있으니, 팬들에게 질문이 있으면 고객복을 즉시 반영해 주십시오!하지만 전기 쌓인 문제가 너무 많아서 기술부문이 일일이 처리해야 할 때, 처리시간은 가능한 한 한 달 이내에 처리될 것입니다. 이해해 주셔서 감사합니다!플랫폼 충재통화 기능은 지속적인 최적화를 거쳐 이미 속속 충제폐를 개통했으나 최적화 과정은 계속되고 있으며 플랫폼은 내부와 외부 주류 종류의 충제로 이뤄질 것이다.</p><p>V -token 고객센터</p><p>2019.5.28</p><p><br></p>\",\n                \"create_time\": 1559375854,\n                \"outline\": \"大家好，为了大家反馈问题能得到及时的回复和处理，近段时间客服部门在调试智能化问题处理系统，导致一些反馈问题拖延了时间，平台在此向大家致以诚挚的歉意！\"\n            }\n        ],\n        \"total\": 51,\n        \"count\": 10,\n        \"per_page\": 10,\n        \"current_page\": 1,\n        \"total_pages\": 6,\n        \"links\": {\n            \"next\": \"http://3.13.139.54:8883/api/notice?page=2\"\n        }\n    },\n    \"code\": 200,\n    \"msg\": \"success\"\n}"
	json.Unmarshal([]byte(string(ste2)),&name)

	for _,v := range name.Data.Item{
		fmt.Println(v.Outline)
	}










}