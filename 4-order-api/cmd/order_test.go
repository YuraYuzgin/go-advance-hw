package main

import (
	"bytes"
	"cart-api/internal/order"
	"cart-api/internal/product"
	"cart-api/internal/user"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Phone:     "89331112233",
		SessionId: "XKY0BTHM91C7EEk",
		Code:      1234,
	})
	db.Create(&product.Product{
		Name:        "TV",
		Description: "43",
	})
}

func RemoveData(db *gorm.DB) {
	db.Unscoped().
		Where("phone = ?", "89331112233").
		Delete(&user.User{})
	db.Unscoped().
		Where("name = ?", "TV")
}

func TestCreateOrderSuccess(t *testing.T) {
	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&order.OrderRequest{
		Products: []uint{1},
	})
	req, err := http.NewRequest("POST", ts.URL+"/order", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6Ijg5MzMxMTEyMjMzIn0.hsbO52AR8ZIjLkVzJxE4Tx_PQ-5oI8MgRZ9giz_US0Y")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData order.Order
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.ID == 0 {
		t.Fatal("Order invalid")
	}
	if len(resData.Products) == 0 {
		t.Fatal("Order should contain products")
	}

	RemoveData(db)
}
