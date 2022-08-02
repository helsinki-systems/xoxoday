package xoxoday

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func (a *API) DoGetVouchersRequest(data GetVouchersRequestData) ([]GetVouchersResponse, error) {
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
			GetVouchers struct {
				Data []GetVouchersResponse `json:"data"`
			} `json:"getVouchers"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Data.GetVouchers.Data, nil
}

type GetVouchersRequestData struct {
	Limit           int     `json:"limit"`
	Page            int     `json:"page"`
	IncludeProducts string  `json:"includeProducts"`
	ExcludeProducts string  `json:"excludeProducts"`
	ExchangeRate    float64 `json:"exchangeRate"`
	Sort            struct {
		Field string `json:"field"`
		Order string `json:"order"`
	} `json:"sort"`
	Filters []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"filters"`
}

type GetVouchersResponse struct {
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
	IsCommon           string  `json:"isCommon"`
	Fee                float64 `json:"fee"`
	Discount           float64 `json:"discount"`
	ExchangeRate       int     `json:"exchangeRate"`
}

type jsonTime time.Time

func (jt *jsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}

	*jt = jsonTime(t)

	return nil
}

func (a *API) DoGetFiltersRequest(data GetFiltersRequestData) ([]GetFiltersResponse, error) {
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
			GetFilters struct {
				Data []GetFiltersResponse `json:"data"`
			} `json:"getFilters"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Data.GetFilters.Data, nil
}

type GetFiltersRequestData struct {
	FilterGroupCode string `json:"filterGroupCode"`
	IncludeFilters  string `json:"includeFilters"`
	ExcludeFilters  string `json:"excludeFilters"`
}

type GetFiltersResponse struct {
	FilterGroupName        string `json:"filterGroupName"`
	FilterGroupDescription string `json:"filterGroupDescription"`
	FilterGroupCode        string `json:"filterGroupCode"`
	Filters                []struct {
		FilterValue     string `json:"filterValue"`
		ISOCode         string `json:"isoCode"`
		FilterValueCode string `json:"filterValueCode"`
	} `json:"filters"`
}

func (a *API) DoGetBalanceRequest() (*GetBalanceResponse, error) {
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
			GetBalance struct {
				Data GetBalanceResponse `json:"data"`
			} `json:"getBalance"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &res.Data.GetBalance.Data, nil
}

type GetBalanceResponse struct {
	Points   int    `json:"points"`
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

func (a *API) DoGetOrderHistoryRequest(data GetOrderHistoryRequestData) ([]GetOrderHistoryResponse, error) {
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
			GetOrderHistory struct {
				Data []GetOrderHistoryResponse `json:"data"`
			} `json:"getOrderHistory"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Data.GetOrderHistory.Data, nil
}

type GetOrderHistoryRequestData struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
}

type GetOrderHistoryResponse struct {
	OrderID        string   `json:"orderId"`
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

func (a *API) DoGetOrderDetailsRequest(data GetOrderDetailsRequestData) (*GetOrderDetailsResponse, error) {
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
			GetOrderDetails struct {
				Data GetOrderDetailsResponse `json:"data"`
			} `json:"getOrderDetails"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &res.Data.GetOrderDetails.Data, nil
}

type GetOrderDetailsRequestData struct {
	PONumber string `json:"poNumber"`
	OrderID  string `json:"orderId"`
}

type GetOrderDetailsResponse struct {
	OrderID         string    `json:"orderId"`
	Vouchers        []Voucher `json:"vouchers"`
	AmountCharged   float64   `json:"amountCharged"`
	CurrencyCode    string    `json:"currencyCode"`
	CurrencyValue   float64   `json:"currencyValue"`
	DiscountPercent float64   `json:"discountPercent"`
	OrderDiscount   float64   `json:"orderDiscount"`
	OrderTotal      float64   `json:"orderTotal"`
	OrderStatus     string    `json:"orderStatus"`
	DeliveryStatus  string    `json:"deliveryStatus"`
}

type Voucher struct {
	Amount        int     `json:"amount,string"`
	Country       string  `json:"country"`
	Currency      string  `json:"currency"`
	OrderID       string  `json:"orderId"`
	Pin           string  `json:"pin"`
	ProductID     string  `json:"productId"`
	Tag           string  `json:"tag"`
	Type          string  `json:"type"`
	Validity      string  `json:"validity"`
	VoucherCode   string  `json:"voucherCode"`
	CurrencyValue float64 `json:"currencyValue"`
}

func (a *API) DoPlaceOrderRequest(data PlaceOrderRequestData) (*PlaceOrderResponse, error) {
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

type PlaceOrderResponse struct {
	OrderID         int       `json:"orderId"`
	Vouchers        []Voucher `json:"vouchers"`
	AmountCharged   float64   `json:"amountCharged"`
	CurrencyCode    string    `json:"currencyCode"`
	CurrencyValue   float64   `json:"currencyValue"`
	Tag             string    `json:"tag"`
	DiscountPercent float64   `json:"discountPercent"`
	OrderDiscount   float64   `json:"orderDiscount"`
	OrderTotal      float64   `json:"orderTotal"`
	OrderStatus     string    `json:"orderStatus"`
	DeliveryStatus  string    `json:"deliveryStatus"`
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
