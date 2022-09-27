package xoxoday

import (
	"encoding/json"
	"fmt"
	"time"
)

func (a *API) Vouchers(data VouchersRequestData) ([]VouchersResponse, error) {
	req := Request{
		Query: "plumProAPI.mutation.getVouchers",
		Tag:   "plumProAPI",
		Variables: RequestData{
			Data: data,
		},
	}

	resb, err := a.Run(req)
	if err != nil {
		return nil, err
	}

	var res struct {
		Data struct {
			Vouchers struct {
				Data []VouchersResponse `json:"data"`
			} `json:"getVouchers"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Data.Vouchers.Data, nil
}

type VouchersRequestData struct {
	Limit           int                         `json:"limit"`
	Page            int                         `json:"page"`
	IncludeProducts string                      `json:"includeProducts"`
	ExcludeProducts string                      `json:"excludeProducts"`
	ExchangeRate    float64                     `json:"exchangeRate"`
	Sort            VouchersRequestDataSort     `json:"sort"`
	Filters         []VouchersRequestDataFilter `json:"filters"`
}

type VouchersRequestDataSort struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

type VouchersRequestDataFilter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type VouchersResponse struct {
	ProductID                      int      `json:"productId"`
	Name                           string   `json:"name"`
	Description                    string   `json:"description"`
	TermsAndConditionsInstructions string   `json:"termsAndConditionsInstructions"`
	ExpiryAndValidity              string   `json:"expiryAndValidity"`
	RedemptionInstructions         string   `json:"redemptionInstructions"`
	Categories                     string   `json:"categories"`
	LastUpdateDate                 jsonTime `json:"lastUpdateDate"`
	ImageURL                       string   `json:"imageUrl"`
	CurrencyCode                   string   `json:"currencyCode"`
	CurrencyName                   string   `json:"currencyName"`
	CountryName                    string   `json:"countryName"`
	CountryCode                    string   `json:"countryCode"`
	Countries                      []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"countries"`
	ExchangeRateRule   float64 `json:"exchangeRateRule"`
	ValueType          string  `json:"valueType"`
	MaxValue           float64 `json:"maxValue"`
	MinValue           float64 `json:"minValue"`
	ValueDenominations string  `json:"valueDenominations"`
	TATInDays          int     `json:"tatInDays"`
	UsageType          string  `json:"usageType"`
	DeliveryType       string  `json:"deliveryType"`
	// TODO: This might need to be a boolean
	IsCommon     string  `json:"isCommon"`
	Fee          float64 `json:"fee"`
	Discount     float64 `json:"discount"`
	ExchangeRate int     `json:"exchangeRate"`
}

func (a *API) Filters(data FiltersRequestData) ([]FiltersResponse, error) {
	req := Request{
		Query: "plumProAPI.mutation.getFilters",
		Tag:   "plumProAPI",
		Variables: RequestData{
			Data: data,
		},
	}

	resb, err := a.Run(req)
	if err != nil {
		return nil, err
	}

	var res struct {
		Data struct {
			Filters struct {
				Data []FiltersResponse `json:"data"`
			} `json:"getFilters"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Data.Filters.Data, nil
}

type FiltersRequestData struct {
	FilterGroupCode string `json:"filterGroupCode"`
	IncludeFilters  string `json:"includeFilters"`
	ExcludeFilters  string `json:"excludeFilters"`
}

type FiltersResponse struct {
	FilterGroupName        string `json:"filterGroupName"`
	FilterGroupDescription string `json:"filterGroupDescription"`
	FilterGroupCode        string `json:"filterGroupCode"`
	Filters                []struct {
		FilterValue     string `json:"filterValue"`
		ISOCode         string `json:"isoCode"`
		FilterValueCode string `json:"filterValueCode"`
	} `json:"filters"`
}

func (a *API) Balance() (*BalanceResponse, error) {
	req := Request{
		Query: "plumProAPI.query.getBalance",
		Tag:   "plumProAPI",
		Variables: RequestData{
			Data: struct{}{},
		},
	}

	resb, err := a.Run(req)
	if err != nil {
		return nil, err
	}

	var res struct {
		Data struct {
			Balance struct {
				Data BalanceResponse `json:"data"`
			} `json:"getBalance"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &res.Data.Balance.Data, nil
}

type BalanceResponse struct {
	Points   float64 `json:"points"`
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

func (a *API) OrderHistory(data OrderHistoryRequestData) ([]OrderHistoryResponse, error) {
	req := Request{
		Query: "plumProAPI.mutation.getOrderHistory",
		Tag:   "plumProAPI",
		Variables: RequestData{
			Data: data,
		},
	}

	resb, err := a.Run(req)
	if err != nil {
		return nil, err
	}

	var res struct {
		Data struct {
			OrderHistory struct {
				Data []OrderHistoryResponse `json:"data"`
			} `json:"getOrderHistory"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Data.OrderHistory.Data, nil
}

type OrderHistoryRequestData struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
}

func (rd *OrderHistoryRequestData) SetStartEnd(start, end time.Time) *OrderHistoryRequestData {
	const format = "2006-01-02"

	rd.StartDate = start.Format(format)
	rd.EndDate = end.Format(format)

	return rd
}

type OrderHistoryResponse struct {
	OrderID        int      `json:"orderId"`
	OrderDate      jsonTime `json:"orderDate"`
	EMail          string   `json:"email"`
	DeliveryStatus string   `json:"deliveryStatus"`
	Tag            string   `json:"tag"`
	PONumber       string   `json:"poNumber"`
	Products       []struct {
		ProductName          string  `json:"productName"`
		ReceiverMobileNumber string  `json:"receiverMobileNumber"`
		Price                float64 `json:"price"`
		Quantity             int     `json:"quantity"`
		OrderProductStatus   string  `json:"orderProductStatus"`
	} `json:"products"`
}

func (a *API) OrderDetails(data OrderDetailsRequestData) (*OrderDetailsResponse, error) {
	req := Request{
		Query: "plumProAPI.mutation.getOrderDetails",
		Tag:   "plumProAPI",
		Variables: RequestData{
			Data: data,
		},
	}

	resb, err := a.Run(req)
	if err != nil {
		return nil, err
	}

	var res struct {
		Data struct {
			OrderDetails struct {
				Data OrderDetailsResponse `json:"data"`
			} `json:"getOrderDetails"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &res.Data.OrderDetails.Data, nil
}

type OrderDetailsRequestData struct {
	OrderID  int    `json:"orderId"`
	PONumber string `json:"poNumber"`
}

type OrderDetailsResponse struct {
	OrderID       int       `json:"orderId"`
	Vouchers      []Voucher `json:"vouchers"`
	AmountCharged float64   `json:"amountCharged"`
	CurrencyCode  string    `json:"currencyCode"`
	CurrencyValue float64   `json:"currencyValue"`
	Tag           string    `json:"tag"`
	// NOTE: API doc says that this is a float, when in fact it's a string
	DiscountPercent string `json:"discountPercent"`
	// NOTE: API doc says that this is a float, when in fact it's a string
	OrderDiscount  string  `json:"orderDiscount"`
	OrderTotal     float64 `json:"orderTotal"`
	OrderStatus    string  `json:"orderStatus"`
	DeliveryStatus string  `json:"deliveryStatus"`
}

type Voucher struct {
	Amount        int     `json:"amount"`
	Country       string  `json:"country"`
	Currency      string  `json:"currency"`
	OrderID       int     `json:"orderId"`
	Pin           string  `json:"pin"`
	ProductID     int     `json:"productId"`
	Tag           string  `json:"tag"`
	Type          string  `json:"type"`
	Validity      string  `json:"validity"`
	VoucherCode   string  `json:"voucherCode"`
	CurrencyValue float64 `json:"currencyValue"`
}

func (a *API) PlaceOrder(data PlaceOrderRequestData) (*PlaceOrderResponse, error) {
	req := Request{
		Query: "plumProAPI.mutation.placeOrder",
		Tag:   "plumProAPI",
		Variables: RequestData{
			Data: data,
		},
	}

	resb, err := a.Run(req)
	if err != nil {
		return nil, err
	}

	var res struct {
		Data struct {
			PlaceOrder struct {
				Data PlaceOrderResponse `json:"data"`
			} `json:"placeOrder"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &res.Data.PlaceOrder.Data, nil
}

type PlaceOrderRequestData struct {
	ProductID           int         `json:"productId"`
	Quantity            int         `json:"quantity"`
	Denomination        int         `json:"denomination"`
	EMail               string      `json:"email"`
	Contact             string      `json:"contact"`
	Tag                 string      `json:"tag"`
	PONumber            string      `json:"poNumber"`
	NotifyAdminEMail    jsonBoolInt `json:"notifyAdminEmail"`
	NotifyReceiverEMail jsonBoolInt `json:"notifyReceiverEmail"`
}

func (rd *PlaceOrderRequestData) SetNotifyAdminEMail(b bool) *PlaceOrderRequestData {
	rd.NotifyAdminEMail = jsonBoolInt{bool: b}

	return rd
}

func (rd *PlaceOrderRequestData) SetNotifyReceiverEMail(b bool) *PlaceOrderRequestData {
	rd.NotifyReceiverEMail = jsonBoolInt{bool: b}

	return rd
}

type PlaceOrderResponse struct {
	OrderID       int       `json:"orderId"`
	Vouchers      []Voucher `json:"vouchers"`
	AmountCharged float64   `json:"amountCharged"`
	CurrencyCode  string    `json:"currencyCode"`
	CurrencyValue float64   `json:"currencyValue"`
	Tag           string    `json:"tag"`
	// NOTE: API doc says that this is a float, when in fact it's a string
	DiscountPercent string `json:"discountPercent"`
	// NOTE: API doc says that this is a float, when in fact it's a string
	OrderDiscount  string  `json:"orderDiscount"`
	OrderTotal     float64 `json:"orderTotal"`
	OrderStatus    string  `json:"orderStatus"`
	DeliveryStatus string  `json:"deliveryStatus"`
}

// General io

type Request struct {
	Query     string      `json:"query"`
	Tag       string      `json:"tag"`
	Variables RequestData `json:"variables"`
}

type RequestData struct {
	Data interface{} `json:"data"`
}
