package repo

type OrderData struct {
	OrderUid    string `db:"order_uid"`
	TrackNumber string `db:"track_number"`
	Entry       string `db:"entry"`
	Delivery    struct {
		Name    string `db:"name"`
		Phone   string `db:"phone"`
		Zip     string `db:"zip"`
		City    string `db:"city"`
		Address string `db:"address"`
		Region  string `db:"region"`
		Email   string `db:"email"`
	} `db:"delivery"`
	Payment struct {
		Transaction  string `db:"transaction"`
		RequestId    string `db:"request_id"`
		Currency     string `db:"currency"`
		Provider     string `db:"provider"`
		Amount       int    `db:"amount"`
		PaymentDt    int    `db:"payment_dt"`
		Bank         string `db:"bank"`
		DeliveryCost int    `db:"delivery_cost"`
		GoodsTotal   int    `db:"goods_total"`
		CustomFee    int    `db:"custom_fee"`
	} `db:"payment"`
	Items []struct {
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
	}
	Locale            string `db:"locale"`
	InternalSignature string `db:"internal_signature"`
	CustomerId        string `db:"customer_id"`
	DeliveryService   string `db:"delivery_service"`
	ShardKey          string `db:"shard_key"`
	SmId              int    `db:"sm_id"`
	DateCreated       string `db:"date_created"`
	OofShard          string `db:"oof_shard"`
}
