package categorytest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"pzn-restful-api/model/domain"
	"pzn-restful-api/repository/category_repo"
	"pzn-restful-api/test/setup"
	templatemethod "pzn-restful-api/test/template_method"
	"strconv"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

const (
	url = "http://localhost:3000/api/categories"
)

func Test_CreateCategorySuccess(t *testing.T){
	//Setup Db, Router, dan Truncate Table
	router := templatemethod.RouterAndTruncateTable()

	// buat request body
	requestBody := strings.NewReader(`{"name":"Gadget"}`)

	// buat objek request
	request := templatemethod.CreateRequest(http.MethodPost, url,requestBody)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,200,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t,200,int(responseBody["code"].(float64)))
	fmt.Println(responseBody);
}
func Test_CreateCategoryFailed(t *testing.T){
	//Setup Db, Router, dan Truncate Table
	router := templatemethod.RouterAndTruncateTable()

	// buat request body
	requestBody := strings.NewReader(`{"name":""}`)

	// buat objek request
	request := templatemethod.CreateRequest(http.MethodPost,url,requestBody)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,400,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t,400,int(responseBody["code"].(float64)))
	fmt.Println(responseBody);
}
func Test_UpdateCategorySuccess(t *testing.T){
	// koneksi ke db test
	db := setup.SetupTestDB()
	setup.TruncateCategory(db)
	
	// create data terlebih dahulu
	tx, _ := db.Begin()
	categoryRepo := category_repo.NewCategoryRepository()
	category := categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setup.SetupRouter(db)
	
	// buat request body
	requestBody := strings.NewReader(`{"name":"Gadget Update"}`)

	// buat objek request
	request := templatemethod.CreateRequest(http.MethodPut,url+"/"+strconv.Itoa(category.Id),requestBody)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,200,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// assert.Equal(t,200,int(responseBody["code"].(float64)))
	fmt.Println(responseBody);
}
func Test_UpdateCategoryFailed(t *testing.T){
	// koneksi ke db test
	db := setup.SetupTestDB()
	setup.TruncateCategory(db)
	
	// create data terlebih dahulu
	tx, _ := db.Begin()
	categoryRepo := category_repo.NewCategoryRepository()
	category := categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setup.SetupRouter(db)
	
	// buat request body
	requestBody := strings.NewReader(`{"name":""}`)
	// requestBody := strings.NewReader(`{"name":"Update Name"}`)

	// buat objek request
	request := templatemethod.CreateRequest(http.MethodPut,url+"/"+strconv.Itoa(category.Id),requestBody)
	// request := templatemethod.CreateRequest(http.MethodPut,url+"/2",requestBody)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,404,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// assert.Equal(t,200,int(responseBody["code"].(float64)))
	fmt.Println(responseBody);
}
func Test_DeleteCategorySuccess(t *testing.T){
	// koneksi ke db test
	db := setup.SetupTestDB()
	setup.TruncateCategory(db)
	
	// create data terlebih dahulu
	tx, _ := db.Begin()
	categoryRepo := category_repo.NewCategoryRepository()
	category := categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setup.SetupRouter(db)
	
	// buat objek request
	request := templatemethod.CreateRequest(http.MethodDelete,url+"/"+strconv.Itoa(category.Id),nil)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,200,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody);
}
func Test_DeleteCategoryFailed(t *testing.T){
	// koneksi ke db test
	db := setup.SetupTestDB()
	setup.TruncateCategory(db)
	
	// create data terlebih dahulu
	tx, _ := db.Begin()
	categoryRepo := category_repo.NewCategoryRepository()
	categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setup.SetupRouter(db)
	
	// buat objek request
	request := templatemethod.CreateRequest(http.MethodDelete,url+"/300",nil)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,404,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody);
}
func Test_GetCategorySuccess(t *testing.T){
	// koneksi ke db test
	db := setup.SetupTestDB()
	setup.TruncateCategory(db)
	
	// create data terlebih dahulu
	tx, _ := db.Begin()
	categoryRepo := category_repo.NewCategoryRepository()
	categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "F&B",
	})
	categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "FMCG",
	})
	tx.Commit()

	router := setup.SetupRouter(db)
	
	// buat objek request
	request := templatemethod.CreateRequest(http.MethodGet,url,nil)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,200,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody);
}

func Test_FindByIdCategorySuccess(t *testing.T){
	// koneksi ke db test
	db := setup.SetupTestDB()
	setup.TruncateCategory(db)
	
	// create data terlebih dahulu
	tx, _ := db.Begin()
	categoryRepo := category_repo.NewCategoryRepository()
	category := categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setup.SetupRouter(db)
	
	// buat objek request
	request := templatemethod.CreateRequest(http.MethodGet,url+"/"+strconv.Itoa(category.Id),nil)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,200,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody);
}
func Test_FindByIdCategoryFailed(t *testing.T){
	// koneksi ke db test
	db := setup.SetupTestDB()
	setup.TruncateCategory(db)
	
	// create data terlebih dahulu
	tx, _ := db.Begin()
	categoryRepo := category_repo.NewCategoryRepository()
	categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setup.SetupRouter(db)
	
	// buat objek request
	request := templatemethod.CreateRequest(http.MethodGet,url+"/300",nil)
	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,404,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody);
}
func Test_UnauthorizedFailed(t *testing.T){
	// koneksi ke db test
	db := setup.SetupTestDB()
	setup.TruncateCategory(db)
	
	// create data terlebih dahulu
	tx, _ := db.Begin()
	categoryRepo := category_repo.NewCategoryRepository()
	category := categoryRepo.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setup.SetupRouter(db)

	// buat objek request
	request := httptest.NewRequest(http.MethodGet,url+"/"+strconv.Itoa(category.Id),nil)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("X-Api-Key", "salah")

	// buat objek response
	recorder := httptest.NewRecorder()

	// jalankan servernya
	router.ServeHTTP(recorder, request)

	//ambil responsenya
	response := recorder.Result()

	// test
	assert.Equal(t,401,response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody);
}