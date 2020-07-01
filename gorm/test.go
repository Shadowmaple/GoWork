package main

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	//	Id		int		`gorm:primary_key`
	Name  string `gorm:"not null;unique"`
	Age   int
	Email string
}

func main() {
	var lock sync.Mutex
	var wait, wait2 sync.WaitGroup
	wait.Add(1)
	wait2.Add(1)

	db, err := gorm.Open("mysql",
		"<username>:<passward>@/<dbName>?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic(err)
		return
	}

	db.DropTable(&User{})

	func() {
		lock.Lock()
		defer lock.Unlock()
		defer wait.Done()

		if !db.HasTable(&User{}) {
			db.CreateTable(User{})
			fmt.Println("create a new table")

			var names = []string{"zmc", "maple", "nick", "shdw", "jzc", "wzh"}
			for _, name := range names {
				user := User{
					Name:  name,
					Age:   20,
					Email: "123@qq.com",
				}
				db.Create(&user)
			}
		}
	}()

	var user User
	go func() {
		wait.Wait()
		lock.Lock()
		defer lock.Unlock()
		defer wait2.Done()

		db.Where("Name = ?", "nick").Find(&user)
		fmt.Println(user)
	}()

	wait2.Wait()
	if user.Name != "nick" {
		fmt.Println("user fetch error")
		return
	}
	user.Age = 18
	user.Email = "unknow@xx.com"
	db.Save(&user)

	fmt.Println(user)
}
