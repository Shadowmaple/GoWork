package main

import (
	"fmt"
)

type UserInfo struct {
	Id       int64
	Name     string
	CourseId int64
}

func main() {
	users, err := GetUserList([]int64{})
	if err != nil {
		return
	}
	fmt.Println(users)
	return
}

// 根据用户id获取用户信息，以映射的形式返回
func GetUserList(userId []int64) (map[int64]*UserInfo, error) {
	userList := make(map[int64]*UserInfo, 0)

	for _, id := range userId {
		// 获取姓名
		name, err := GetName(id)
		if err != nil {
			return nil, err
		}

		// 获取课堂号
		courseId, err := GetCourseId(id)
		if err != nil {
			return nil, err
		}

		// 添加用户
		userList[id] = &UserInfo{
			Id:       id,
			Name:     name,
			CourseId: courseId,
		}
	}

	return userList, nil
}

// 获取姓名
func GetName(id int64) (string, error) {
	// ...
	return "", nil
}

// 获取课堂号
func GetCourseId(id int64) (int64, error) {
	// ...
	return 0, nil
}
