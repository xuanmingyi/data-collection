package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

// 静态文件
type SPAHandler struct {
	IndexPath  string
	StaticPath string
}

func (h SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.StaticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}

// 服务
type Service struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Endpoints []string `json:"endpoints"`
	UpdateAt  string   `json:"update_at"`
}

type ServiceResponse struct {
	Code  int       `json:"code"`
	Msg   string    `json:"msg"`
	Count int       `json:"count"`
	Data  []Service `json:"data"`
}

func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")

	var rsp ServiceResponse

	rsp.Data = []Service{
		Service{ID: 111, Name: "aaaa", Endpoints: []string{"aaaa", "bbbb", "cccc"}, UpdateAt: "10 s 前"},
	}
	rsp.Count = len(rsp.Data)

	content, err := json.Marshal(rsp)
	if err != nil {
		rsp.Code = 1
		rsp.Msg = "error"
		rsp.Count = 0
		rsp.Data = []Service{}
		content, _ = json.Marshal(rsp)
		w.Write(content)
	}
	w.Write(content)
}
