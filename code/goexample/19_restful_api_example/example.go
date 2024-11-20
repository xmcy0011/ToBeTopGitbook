package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// LoginRequest 登录请求结构体
type LoginRequest struct {
	UserName string `json:"user_name"`
	UserPwd  string `json:"user_pwd"`
}

// LoginResponse 登录响应结构体
type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

// Error 错误响应结构体
type Error struct {
	Code    int    `json:"code"`
	Cause   string `json:"cause"`
	Message string `json:"message"`
}

func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, 40000001, err.Error(), "参数格式错误")
		return
	}

	// 验证参数
	if req.UserName == "" || req.UserPwd == "" {
		writeError(w, http.StatusBadRequest, 40000001, "用户名或密码为空", "参数不能为空")
		return
	}

	// 这里应该有真实的用户验证逻辑
	if req.UserName == "admin" && req.UserPwd == "111111" {
		resp := LoginResponse{
			ID:    1024,
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		}
		writeJSON(w, http.StatusOK, resp)
		return
	}

	writeError(w, http.StatusBadRequest, 40000001, "用户名或密码错误", "登录失败")
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status, code int, cause, message string) {
	writeJSON(w, status, Error{
		Code:    code,
		Cause:   cause,
		Message: message,
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/auth/login", login).Methods("POST")

	log.Printf("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
