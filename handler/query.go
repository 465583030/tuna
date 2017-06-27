package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mylxsw/tuna/libs"
	"github.com/mylxsw/tuna/storage"
)

// Query 方法用于查询hash值与url的对应关系
func Query(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	if driver := storage.Default(); driver != nil {
		url := driver.Get(hash)
		if url == "" {
			libs.SendNotFoundResponse(w)
			return
		}

		libs.Redirect(url, w)

		return
	}

	w.Write(libs.Failed("操作失败"))
}