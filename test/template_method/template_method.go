package templatemethod

import (
	"io"
	"net/http"
	"net/http/httptest"
	"pzn-restful-api/test/setup"
)

func RouterAndTruncateTable() http.Handler{
	// koneksi ke db test
	db := setup.SetupTestDB()
	// buat router dan kosong data setiap kali test dijalankan
	router := setup.SetupRouter(db)
	setup.TruncateCategory(db)

	return router
}

func CreateRequest(method string,url string, requestBody io.Reader) *http.Request{
	// buat objek request
	request := httptest.NewRequest(method, url, requestBody)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("X-Api-Key", "1118020")

	return request
}