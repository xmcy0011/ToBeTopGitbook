package main

import "net/http"

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get("city")
		if city == "上海" {
			w.Write([]byte("晴天\n"))
		} else {
			w.Write([]byte("未知\n"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
