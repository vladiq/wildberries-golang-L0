package internal

import (
	"wb_l0/internal/repo"
	"wb_l0/internal/services/data_generator"
)

func MapGeneratedToStored(mj *data_generator.ModelJSON) *repo.OrderData {
	od := repo.OrderData{}

	od.OrderUid = mj.OrderUid
	od.TrackNumber = mj.TrackNumber
	od.Entry = mj.Entry
	od.Delivery.Name = mj.Delivery.Name
	od.Delivery.Phone = mj.Delivery.Phone
	od.Delivery.Zip = mj.Delivery.Zip
	od.Delivery.City = mj.Delivery.City
	od.Delivery.Address = mj.Delivery.Address
	od.Delivery.Region = mj.Delivery.Region
	od.Delivery.Email = mj.Delivery.Email
	od.Payment.RequestId = mj.Payment.RequestId
	od.Payment.Currency = mj.Payment.Currency
	od.Payment.Transaction = mj.Payment.Transaction
	od.Payment.Provider = mj.Payment.Provider
	od.Payment.Amount = mj.Payment.Amount
	od.Payment.PaymentDt = mj.Payment.PaymentDt
	od.Payment.Bank = mj.Payment.Bank
	od.Payment.DeliveryCost = mj.Payment.DeliveryCost
	od.Payment.GoodsTotal = mj.Payment.GoodsTotal
	od.Payment.CustomFee = mj.Payment.CustomFee

	numItems := len(mj.Items)
	for i := 0; i < numItems; i++ {
		curItems := struct {
			ChrtId      int    `db:"chrt_id"`
			TrackNumber string `db:"track_number"`
			Price       int    `db:"price"`
			RId         string `db:"rid"`
			Name        string `db:"name"`
			Sale        int    `db:"sale"`
			Size        string `db:"size"`
			TotalPrice  int    `db:"total_price"`
			NmId        int    `db:"nm_id"`
			Brand       string `db:"brand"`
			Status      int    `db:"status"`
		}{
			ChrtId:      mj.Items[i].ChrtId,
			TrackNumber: mj.Items[i].TrackNumber,
			Price:       mj.Items[i].Price,
			RId:         mj.Items[i].RId,
			Name:        mj.Items[i].Name,
			Sale:        mj.Items[i].Sale,
			Size:        mj.Items[i].Size,
			TotalPrice:  mj.Items[i].TotalPrice,
			NmId:        mj.Items[i].NmId,
			Brand:       mj.Items[i].Brand,
			Status:      mj.Items[i].Status,
		}
		od.Items = append(od.Items, curItems)
	}
	od.Locale = mj.Locale
	od.InternalSignature = mj.InternalSignature
	od.CustomerId = mj.CustomerId
	od.DeliveryService = mj.DeliveryService
	od.ShardKey = mj.ShardKey
	od.SmId = mj.SmId
	od.DateCreated = mj.DateCreated
	od.OofShard = mj.OofShard

	return &od
}

func MapStoredToDisplayed(od *repo.OrderData) *data_generator.ModelJSON {
	mj := data_generator.ModelJSON{}

	mj.OrderUid = od.OrderUid
	mj.TrackNumber = od.TrackNumber
	mj.Entry = od.Entry
	mj.Delivery.Name = od.Delivery.Name
	mj.Delivery.Phone = od.Delivery.Phone
	mj.Delivery.Zip = od.Delivery.Zip
	mj.Delivery.City = od.Delivery.City
	mj.Delivery.Address = od.Delivery.Address
	mj.Delivery.Region = od.Delivery.Region
	mj.Delivery.Email = od.Delivery.Email
	mj.Payment.RequestId = od.Payment.RequestId
	mj.Payment.Currency = od.Payment.Currency
	mj.Payment.Transaction = od.Payment.Transaction
	mj.Payment.Provider = od.Payment.Provider
	mj.Payment.Amount = od.Payment.Amount
	mj.Payment.PaymentDt = od.Payment.PaymentDt
	mj.Payment.Bank = od.Payment.Bank
	mj.Payment.DeliveryCost = od.Payment.DeliveryCost
	mj.Payment.GoodsTotal = od.Payment.GoodsTotal
	mj.Payment.CustomFee = od.Payment.CustomFee

	numItems := len(od.Items)
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
			ChrtId:      od.Items[i].ChrtId,
			TrackNumber: od.Items[i].TrackNumber,
			Price:       od.Items[i].Price,
			RId:         od.Items[i].RId,
			Name:        od.Items[i].Name,
			Sale:        od.Items[i].Sale,
			Size:        od.Items[i].Size,
			TotalPrice:  od.Items[i].TotalPrice,
			NmId:        od.Items[i].NmId,
			Brand:       od.Items[i].Brand,
			Status:      od.Items[i].Status,
		}
		mj.Items = append(mj.Items, curItems)
	}
	mj.Locale = od.Locale
	mj.InternalSignature = od.InternalSignature
	mj.CustomerId = od.CustomerId
	mj.DeliveryService = od.DeliveryService
	mj.ShardKey = od.ShardKey
	mj.SmId = od.SmId
	mj.DateCreated = od.DateCreated
	mj.OofShard = od.OofShard

	return &mj
}
