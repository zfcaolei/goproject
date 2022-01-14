package main

import (
	"encoding/json"
	"fmt"
)

//反序列化


//如果不知道反序列化类型 可以先直接map[string]interface{}  或者直接粘贴其中结构编辑器自动转换结构体
type ss map[string]interface{}

type tes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Data []struct {
			OrderNo     string `json:"order_no"`
			PayAmount   string `json:"pay_amount"`
			AssetCode   string `json:"asset_code"`
			PayStatus   int    `json:"pay_status"`
			OrderAmount string `json:"order_amount"`
			CreatedAt   string `json:"created_at"`
		} `json:"data"`
		PageParams string `json:"page_params"`
	} `json:"data"`
}

type aa struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		PaymentAssetLists []string `json:"payment_asset_lists"`
		CusOptLimit       struct {
			RefreshPriceTimeout int `json:"refresh_price_timeout"`
			SafeVerifyTimeout   int `json:"safe_verify_timeout"`
			SinglePrice         int `json:"single_price"`
			DailyPrice          int `json:"daily_price"`
		} `json:"cus_opt_limit"`
		Switch struct {
			ScanQrcode string `json:"scan_qrcode"`
		} `json:"switch"`
	} `json:"data"`
}



type sss struct {
	Name string
	Age int8

}

func main()  {
	strs := "{\"Name\":\"曹磊\",\"Age\":22}"

	var m1 map[string]interface{}

	err := json.Unmarshal([]byte(strs),&m1)
	if err !=nil {
		fmt.Println("反序列化失败 err=",err)
	}
	fmt.Println(m1)



	var sss tes
	str :=  "{\n  \"code\": 200,\n  \"msg\": \"success\",\n  \"data\": {\n    \"data\": [\n      {\n        \"order_no\": \"202010191144362002202010907376\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 2000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 11:44:36\"\n      },\n      {\n        \"order_no\": \"202010190954282002202010572585\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 2000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:54:28\"\n      },\n      {\n        \"order_no\": \"202010190949112002202010457064\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:49:11\"\n      },\n      {\n        \"order_no\": \"202010190948012002202010764588\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:48:01\"\n      },\n      {\n        \"order_no\": \"202010190947352002202010566850\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:47:35\"\n      },\n      {\n        \"order_no\": \"202010190939092002202010995681\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:39:09\"\n      },\n      {\n        \"order_no\": \"202010190929432002202010248295\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:29:43\"\n      },\n      {\n        \"order_no\": \"202010190923282002202010834341\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:23:28\"\n      },\n      {\n        \"order_no\": \"202010190923092002202010836807\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:23:09\"\n      },\n      {\n        \"order_no\": \"202010190913502002202010835083\",\n        \"pay_amount\": \"15.1651600\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-19 09:13:50\"\n      },\n      {\n        \"order_no\": \"202010161826562002202010230652\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-16 18:26:56\"\n      },\n      {\n        \"order_no\": \"202010161756432002202010675643\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-16 17:56:43\"\n      },\n      {\n        \"order_no\": \"202010161754222002202010661081\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-16 17:54:22\"\n      },\n      {\n        \"order_no\": \"202010161751432002202010424892\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"2020-10-16 17:51:43\"\n      },\n      {\n        \"order_no\": \"202010161744472002202010405133\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"1970-01-01 08:33:40\"\n      },\n      {\n        \"order_no\": \"202010161750262002202010811120\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"1970-01-01 08:33:40\"\n      },\n      {\n        \"order_no\": \"202010161749412002202010553120\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"1970-01-01 08:33:40\"\n      },\n      {\n        \"order_no\": \"202010161748542002202010587796\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"1970-01-01 08:33:40\"\n      },\n      {\n        \"order_no\": \"202010161748472002202010720905\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"1970-01-01 08:33:40\"\n      },\n      {\n        \"order_no\": \"202010161747532002202010880430\",\n        \"pay_amount\": \"15.1197500\",\n        \"asset_code\": \"PAYDEX\",\n        \"pay_status\": 1000,\n        \"order_amount\": \"1.00\",\n        \"created_at\": \"1970-01-01 08:33:40\"\n      }\n    ],\n    \"page_params\": \"ITlvEfDt+2wHDX0V\\/3o2eqBb\\/qxwbWO\\/DUsyLt67LHoZBtwPvogSL5FABcNiFZtw3MwEkZBrG6cluFRoA6ndIW1be4c28Dcwz\\/szfSYLvqJwjzxt65TSoHlWr4vt7nTby2pHWjiP43XEBT6EhCiEUbuz2kUJqfY6LRkJ3+BGaQNT5Wgn\\/ljkzum2Hod\\/4Bt\\/QRoiUZlKGKKZkZUpfJ3Vvq6ZatayqH+Esp+rBzhBGBIJfdXqWjdfyIAQcmJTDK6JdSL3Vp7sI2762fkCR9txLcmlTSN+E4jJkE7DgAaABFmO1ulkfDpN9Hzm\\/6G0z\\/j6RAb22a51rVIOkBrAmQhsRlpta26jXHWl8TSXoXCqVbv5raVROw\\/5tGL1yZz2s3hT4WC\\/rfaLeHlm7Hk+zZvjm2E9Ke3swK+m+j6LRq2dzQTVBOtFLk8LKJQ9bvPToBbQc\\/iAOCLf3U+OsIYrkRBNe\\/ruBev8P8kIM5koQiTw+PQJFCaJvxUTpvKER1p1+w\\/44\\/86PnLC360ixT4QftwPa2tqM7yUIMBs032TJN2c8a6NQCc1SO\\/Lf\\/wIE8suDUbIkLBSp+MXnlwFwTonlx8hTVKT3d+D6xPe2l95m7q3V9wfONK5r+1z27EpGyWCCQsgTHgdl3POe0PqFyk9Z0Kbf3UXVzsGUaTx7ZAdqVQkwCw=\"\n  }\n}"
	fmt.Println(str)
	json.Unmarshal([]byte(string(str)),&sss)




	fmt.Println(sss.Data.Data[0].OrderNo)


}
