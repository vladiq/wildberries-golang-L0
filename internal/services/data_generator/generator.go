package data_generator

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

var givenValues = map[string][]string{
	"track_numbers": {"WBILMTESTTRACK", "WBSECONDTRACK", "WBTHIRDTRACK"},
	"entries":       {"WBIL", "WBILM", "WBIM"},
	"currencies":    {"RUR", "USD", "EUR", "GBP", "HNL"},
	"providers":     {"wbpay", "sberpay", "applepay", "cash"},
	"banks":         {"alpha", "sberbank", "vtb", "tinkoff"},
	"emails":        {"gmail.com", "mail.ru", "wb.ru"},
	"locales":       {"en", "ru", "de", "kz"},
}

type ModelJSON struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		RId         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	}
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerId        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	ShardKey          string `json:"shard_key"`
	SmId              int    `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OofShard          string `json:"oof_shard"`
}

func (mj *ModelJSON) fill(givenValues map[string][]string) {
	mj.OrderUid = generateAlphaNumeric(19)
	mj.TrackNumber = givenValues["track_numbers"][rand.Intn(len(givenValues["track_numbers"]))]
	mj.Entry = getRandomValue(givenValues, "entries")

	mj.Delivery.Name = generateWord(3+rand.Intn(5)) + " " + generateWord(3+rand.Intn(5))
	mj.Delivery.Phone = "+79" + generateNumber(9)
	mj.Delivery.Zip = generateNumber(7)
	mj.Delivery.City = generateWord(4 + rand.Intn(5))
	mj.Delivery.Address = generateWord(5+rand.Intn(10)) + " " + generateNumber(3)
	mj.Delivery.Region = generateWord(4 + rand.Intn(5))
	suf := getRandomValue(givenValues, "emails")
	mj.Delivery.Email = generateWord(1+rand.Intn(4)) + "_" + generateAlphaNumeric(1+rand.Intn(4)) + "@" + suf

	mj.Payment.Transaction = mj.OrderUid
	mj.Payment.RequestId = generateAlphaNumeric(6)
	mj.Payment.Currency = getRandomValue(givenValues, "currencies")
	mj.Payment.Provider = getRandomValue(givenValues, "providers")
	mj.Payment.Amount = rand.Intn(10000)
	mj.Payment.PaymentDt = rand.Intn(1000000)
	mj.Payment.Bank = getRandomValue(givenValues, "banks")
	mj.Payment.DeliveryCost = rand.Intn(3000)
	mj.Payment.GoodsTotal = rand.Intn(1000)
	mj.Payment.CustomFee = rand.Intn(500)

	numItems := 1 + rand.Intn(3)
	for i := 0; i < numItems; i++ {
		curItems := struct {
			ChrtId      int    `json:"chrt_id"`
			TrackNumber string `json:"track_number"`
			Price       int    `json:"price"`
			RId         string `json:"rid"`
			Name        string `json:"name"`
			Sale        int    `json:"sale"`
			Size        string `json:"size"`
			TotalPrice  int    `json:"total_price"`
			NmId        int    `json:"nm_id"`
			Brand       string `json:"brand"`
			Status      int    `json:"status"`
		}{
			ChrtId:      rand.Intn(99999),
			TrackNumber: mj.TrackNumber,
			Price:       rand.Intn(10000),
			RId:         generateAlphaNumeric(21),
			Name:        generateWord(3 + rand.Intn(10)),
			Sale:        rand.Intn(500),
			Size:        strconv.Itoa(rand.Intn(100)),
			TotalPrice:  rand.Intn(10000),
			NmId:        rand.Intn(9999999999),
			Brand:       generateWord(1 + rand.Intn(10)),
			Status:      200 + rand.Intn(20),
		}
		mj.Items = append(mj.Items, curItems)
	}

	mj.Locale = getRandomValue(givenValues, "locales")
	mj.InternalSignature = generateAlphaNumeric(1 + rand.Intn(3))
	mj.CustomerId = generateAlphaNumeric(5)
	mj.DeliveryService = generateWord(5)
	mj.ShardKey = generateNumber(3)
	mj.SmId = rand.Intn(228)
	mj.DateCreated = generateDate()
	mj.OofShard = generateNumber(1)
}

func getRandomValue(givenValues map[string][]string, key string) string {
	randomIdx := rand.Intn(len(givenValues[key]))
	return givenValues[key][randomIdx]
}

func Generate() []byte {
	rand.Seed(time.Now().UnixNano())

	// generate rubbish with certain probability or a valid JSON otherwise
	switch flip := rand.Intn(6); flip % 3 {
	case 0:
		rubbishWord := generateWord(10)
		return []byte(rubbishWord)
	default:
		mj := ModelJSON{}
		mj.fill(givenValues)

		if strJson, err := json.Marshal(mj); err != nil {
			panic(err)
		} else {
			return strJson
		}
	}
}
