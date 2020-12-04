package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserModel struct {
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Password string `json:"password"` // 密码
	Grade    string `json:"grade"`    // 年级
	Age      int8   `json:"age"`      // 年龄
}

// 保存用户数据
var userData = make(map[string]UserModel)

// 注册
func Register(w http.ResponseWriter, req *http.Request) {
	// parse json data in body
	var newUser UserModel
	if err := json.NewDecoder(req.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// 用户名和密码不允许为空
	if newUser.Username == "" || newUser.Password == "" {
		http.Error(w, "username and password are required", 400)
		return
	}

	// 验证用户是否存在
	if _, ok := userData[newUser.Username]; ok {
		fmt.Fprintln(w, "The user has existed")
		return
	}

	// 保存用户数据
	userData[newUser.Username] = newUser

	w.WriteHeader(200)
	fmt.Fprintln(w, "Register OK")
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录
func Login(w http.ResponseWriter, req *http.Request) {
	// 解析 json 请求数据
	var request LoginRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// 用户是否存在
	user, ok := userData[request.Username]
	if !ok {
		http.Error(w, "the user is not found", 404)
		return
	}

	// 校验密码
	if user.Password != request.Password {
		fmt.Fprintln(w, "password is incorrect")
		return
	}

	fmt.Fprintln(w, "login OK")
}

type UpdatePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// 修改密码
func UpdatePassword(w http.ResponseWriter, req *http.Request) {
	// 解析 json
	var request UpdatePasswordRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// 用户是否存在
	user, ok := userData[request.Username]
	if !ok {
		http.Error(w, "the user is not found", 404)
		return
	}

	// 校验旧密码
	if user.Password != request.OldPassword {
		fmt.Fprintln(w, "Password is incorrect")
		return
	}

	// 更新密码
	user.Password = request.NewPassword
	userData[user.Username] = user

	fmt.Fprintln(w, "Update Password OK")
}

// 获取用户信息
func GetUserInfo(w http.ResponseWriter, req *http.Request) {
	// 从 url query 中获取 username
	queries := req.URL.Query()
	username := queries.Get("username")
	if username == "" {
		http.Error(w, "username is required", 400)
		return
	}

	// 用户是否存在
	user, ok := userData[username]
	if !ok {
		http.Error(w, "user is not found", 404)
		return
	}

	// 将用户信息以 json 形式返回
	// json 序列化
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "server error; marshaling json failed", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Grade    string `json:"grade"` // 年级
	Age      int8   `json:"age"`   // 年龄
}

// 修改用户信息
func UpdateUser(w http.ResponseWriter, req *http.Request) {
	var request UpdateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, ok := userData[request.Username]
	if !ok {
		http.Error(w, "the user is not found", 404)
		return
	}

	// 更新字段
	user.Nickname = request.Nickname
	user.Grade = request.Grade
	user.Age = request.Age

	userData[user.Username] = user

	fmt.Fprintln(w, "Update User OK")
}

// 删除用户
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	// 从 query 获取
	queries := req.URL.Query()
	username := queries.Get("username")
	if username == "" {
		http.Error(w, "username is required", 400)
		return
	}

	if _, ok := userData[username]; !ok {
		http.Error(w, "user is not found", 404)
		return
	}

	// 根据 username 移除用户数据
	delete(userData, username)

	fmt.Fprintln(w, "remove user OK")
}

func main() {
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/user/password/update", UpdatePassword)
	http.HandleFunc("/user/info", GetUserInfo)
	http.HandleFunc("/user/update", UpdateUser)
	http.HandleFunc("/user/delete", DeleteUser)

	fmt.Println("Server is running, listening on 8080...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
