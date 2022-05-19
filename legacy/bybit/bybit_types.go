package bybit

// Note: WsInstrument is different
type Instrument struct {
	Symbol               string  `json:"symbol"`
	BidPrice             string  `json:"bid_price"`
	AskPrice             string  `json:"ask_price"`
	LastPrice            string  `json:"last_price"`
	LastTickDirection    string  `json:"last_tick_direction"`
	PrevPrice24H         string  `json:"prev_price_24h"`
	Price24HPcnt         string  `json:"price_24h_pcnt"`
	HighPrice24H         string  `json:"high_price_24h"`
	LowPrice24H          string  `json:"low_price_24h"`
	PrevPrice1H          string  `json:"prev_price_1h"`
	Price1HPcnt          string  `json:"price_1h_pcnt"`
	MarkPrice            string  `json:"mark_price"`
	IndexPrice           string  `json:"index_price"`
	OpenInterest         float64 `json:"open_interest"`
	OpenValue            string  `json:"open_value"`
	TotalTurnover        string  `json:"total_turnover"`
	Turnover24H          string  `json:"turnover_24h"`
	TotalVolume          float64 `json:"total_volume"`
	Volume24H            float64 `json:"volume_24h"`
	FundingRate          string  `json:"funding_rate"`
	PredictedFundingRate string  `json:"predicted_funding_rate"`
	NextFundingTime      string  `json:"next_funding_time"`
	CountdownHour        int     `json:"countdown_hour"`
	//Symbol               string    `json:"symbol"`
	//BidPrice             float64    `json:"bid_price"`
	//AskPrice             float64    `json:"ask_price"`
	//LastPrice            float64    `json:"last_price"`
	//LastTickDirection    string    `json:"last_tick_direction"`
	//PrevPrice24H         float64    `json:"prev_price_24h"`
	//Price24HPcnt         float64    `json:"price_24h_pcnt"`
	//HighPrice24H         float64    `json:"high_price_24h"`
	//LowPrice24H          float64    `json:"low_price_24h"`
	//PrevPrice1H          float64    `json:"prev_price_1h"`
	//Price1HPcnt          float64    `json:"price_1h_pcnt"`
	//MarkPrice            float64    `json:"mark_price"`
	//IndexPrice           float64    `json:"index_price"`
	//OpenInterest         float64   `json:"open_interest"`
	//OpenValue            float64    `json:"open_value"`
	//TotalTurnover        float64    `json:"total_turnover"`
	//Turnover24H          float64    `json:"turnover_24h"`
	//TotalVolume          float64   `json:"total_volume"`
	//Volume24H            float64   `json:"volume_24h"`
	//FundingRate          float64    `json:"funding_rate"`
	//PredictedFundingRate float64    `json:"predicted_funding_rate"`
	//NextFundingTime      time.Time `json:"next_funding_time"`
	//CountdownHour        int       `json:"countdown_hour"`
}

type ResultWrapper struct {
	RetCode          int    `json:"ret_code"`
	RetMsg           string `json:"ret_msg"`
	ExtCode          string `json:"ext_code"`
	ExtInfo          string `json:"ext_info"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int    `json:"rate_limit_status"`
	RateLimitResetMs int64  `json:"rate_limit_reset_ms"`
	RateLimit        int    `json:"rate_limit"`
	//Result  interface{} `json:"result"`
}
