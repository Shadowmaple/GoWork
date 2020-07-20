package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	GetElectricCharge("")
}

func GetElectricCharge(RoomId string) {
	// var charge model.ElectricCharge

	// 第二个URL的 query要求以“2020/5/30”格式输入昨天的日期
	// now := time.Now().AddDate(0, 0, -1)
	// query := now.Format("2006/01/02")
	firUrl := "http://jnb.ccnu.edu.cn/icbs/PurchaseWebService.asmx/getReserveHKAM?AmMeter_ID=" + RoomId
	// secUrl := "http://jnb.ccnu.edu.cn/icbs/PurchaseWebService.asmx/getMeterDayValue?AmMeter_ID=" + RoomId + "&startDate=" + query + "&endDate=" + query

	resp, err := http.Get(firUrl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// xml.Unmarshal(body, &charge)
	// res, err := http.Get(secUrl)
	// if err != nil {
	// 	panic(err)
	// }

	// defer res.Body.Close()
	// newBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// charge.Id = RoomId
	// xml.Unmarshal(newBody, &charge)
	// return charge, nil
}
