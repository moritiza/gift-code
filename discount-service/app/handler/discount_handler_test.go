package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moritiza/gift-code/discount-service/app/dto"
	"github.com/moritiza/gift-code/discount-service/app/helper"
	"github.com/moritiza/gift-code/discount-service/app/repository"
	"github.com/moritiza/gift-code/discount-service/app/service"
	"github.com/moritiza/gift-code/discount-service/config"
	"github.com/moritiza/gift-code/discount-service/test"
)

var (
	testDiscountRepository     repository.DiscountRepository
	testUsedDiscountRepository repository.UsedDiscountRepository
	testDiscountService        service.DiscountService
	testDiscountHandler        DiscountHandler
)

func TestCreate(t *testing.T) {
	var (
		jsonString = []byte(`{"code": "aazAAyy", "code_credit": 120000, "count": 1000}`)
		response   helper.Response
		discount   dto.Discount
	)

	req, err := http.NewRequest(http.MethodPost, "/api/discount", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	cfg := test.Prepare()
	defer config.DisconnectDatabase(cfg)

	testDiscountRepository = repository.NewDiscountRepository(cfg.Database)
	testUsedDiscountRepository = repository.NewUsedDiscountRepository(cfg.Database)
	testDiscountService = service.NewDiscountService(*cfg.Logger, testDiscountRepository, testUsedDiscountRepository)
	testDiscountHandler = NewDiscountHandler(*cfg.Logger, *cfg.Validator, testDiscountService)

	testDiscountHandler.Create(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status created; got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		t.Fatalf("Response is invalid. Error: %v", err)
	}

	if response.Data == nil {
		t.Fatalf("Response data is empty")
	}

	discount.Code = response.Data.(map[string]interface{})["code"].(string)
	discount.CodeCredit = uint64(response.Data.(map[string]interface{})["code_credit"].(float64))
	discount.Count = uint(response.Data.(map[string]interface{})["count"].(float64))

	if discount.Code != "aazAAyy" || discount.CodeCredit != 120000 || discount.Count != 1000 {
		t.Fatalf("Unexpected Response")
	}
}
