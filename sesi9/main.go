package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var url = "https://jsonplaceholder.typicode.com/posts"
var port = ":4000"

type Post struct {
	UserId int
	Title  string
	Body   string
}

func main() {
	http.HandleFunc("/posts", AllowOnlyGet(Auth(GetPostById)))
	http.HandleFunc("/posted", AllowOnlyPost(Auth(PostBody)))
	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)

}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	resp, err := HttpGet(url + "/" + query.Get("id"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": resp,
	})

}

func PostBody(w http.ResponseWriter, r *http.Request) {
	resp, err := HttpPost(url)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": resp,
	})
}

func AllowOnlyGet(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"erros": "method not allowed",
			})
			return
		}
		next(w, r)
	}
}

func AllowOnlyPost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"erros": "method not allowed",
			})
			return
		}
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "need basic auth",
			})
			return
		}
		if username != "admin" || password != "admin" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "usename / password salah.",
			})
			return
		}
		next(w, r)
	}

}

func HttpGet(url string) (map[string]interface{}, error) {
	data, err := req(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)

	return data, nil
}

func HttpPost(url string) (map[string]interface{}, error) {
	reqPayload := Post{
		UserId: 1,
		Title:  "Ini Title",
		Body:   "Ini Body",
	}
	data, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, err
	}
	resp, err := req(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return resp, nil
}

func req(method, url string, body []byte) (map[string]interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
