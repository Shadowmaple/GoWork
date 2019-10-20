package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
	custId		int		`gorm:"primary_key"`
	custName	string	`gorm:"default:xxx"`
	custAddress string
	custCity	string
	custState	string
	custZip		string
	custCountry string
	custContact	string
	custEmail	string
}

func main() {
	db, err := gorm.Open("mysql",
		"<username>:<password>@/<tableName>?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic(err)
		return
	} else {
		fmt.Println("open database successfully")
	}

//	if !db.HasTable(&User{}) {
//		fmt.Println("NO")
//		db.CreateTable(User{})
//	}

//	db.DropTable("users")
//	fmt.Println("Delete success")
//	return


//	cust := Customer{
//		custName: "zmc",
//		custAddress: "zj",
//		custCity: "China",
//	}

	var count int
	err = db.Model(
		&Customer{}).Count(&count).Error
	if err != nil {
		panic(err)
	    return
	} else {
		fmt.Println(count)
	}
//	if err := db.Create(cust).Error; err != nil {
//		panic(err)
//		return
//	} else {
//		fmt.Println("add successfully")
//	}
//	db.NewRecord(item)
//	db.Create(&user)

}
