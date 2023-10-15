package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BalanceTransaction struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	CompanyID       string             `json:"company_id" bson:"company_id"`
	ServiceID       int                `json:"service_id" bson:"service_id"`
	Type            string             `json:"type" bson:"type"`
	SubType         string             `json:"sub_type" bson:"sub_type"`
	Value           float64            `json:"value" bson:"value"`
	PreviousBalance float64            `json:"previous_balance" bson:"previous_balance"`
	CurrencyID      int                `json:"currency_id" bson:"currency_id"`
	ExchangeRate    float64            `json:"exchange_rate" bson:"exchange_rate"`
	ValueInUsd      float64            `json:"value_in_usd" bson:"value_in_usd"`
	Status          string             `json:"status" bson:"status"`
	Pulse           int                `json:"pulse" bson:"pulse"`
	Quantity        int                `json:"quantity" bson:"quantity"`
	Balance         float64            `json:"balance" bson:"balance"`
	TraceID         string             `json:"trace_id" bson:"trace_id"`
	SourceID        string             `json:"source_id" bson:"source_id"`
	Reason          string             `json:"reason" bson:"reason"`
	ServiceName     string             `json:"service_name" bson:"service_name"`
	ServiceCategory string             `json:"service_category" bson:"service_category"`
	RefID           string             `json:"ref_id" bson:"ref_id"`
	BillingType     int                `json:"billing_type" bson:"billing_type"`
	Date            time.Time          `json:"date" bson:"date"`

	PostpaidBalance         float64 `json:"postpaid_balance" bson:"postpaid_balance"`
	PostpaidPreviousBalance float64 `json:"postpaid_previous_balance" bson:"postpaid_previous_balance"`

	//UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`
	CreatedBy             string `json:"created_by" bson:"created_by"`
	CreatedAtTZ           string `json:"created_at_tz" bson:"created_at_tz"`
	IsWalletUsageMovement bool   `json:"is_wallet_usage_movement" bson:"is_wallet_usage_movement"`
	IsWalletDeducted      int    `json:"is_wallet_deducted" bson:"is_wallet_deducted"`
	UserID                string `json:"user_id" bson:"user_id"`
	WhatsappID            string `json:"whatsapp_id" bson:"whatsapp_id"`
	IsPackageDeducted     int    `json:"is_package_deducted" bson:"is_package_deducted"`
	CurrencyCode          string `json:"currency_code" bson:"currency_code"`
	Symbol                string `json:"symbol" bson:"symbol"`

	// Subaccount fields
	Subaccount *Subaccount `json:"subaccount,omitempty" bson:"subaccount,omitempty"`

	// Usage
	Usage []*UsageInfo `json:"usage" bson:"usage"`

	IsKYCApproved bool `json:"is_kyc_approved" bson:"is_kyc_approved"`
}

type Subaccount struct { // TODO: move to entity package?
	// ParentCompanyID is the parent company id of the subaccount. "" for parent company
	ParentCompanyID string `json:"parent_company_id" bson:"parent_company_id"`
	// SubaccountCompanyID of the subaccount. "" for Subaccount company
	SubaccountCompanyID string `json:"subaccount_company_id" bson:"subaccount_company_id"`
	// SubaccountType can be assigned or shared
	SubaccountType string `json:"subaccount_type" bson:"subaccount_type"`
	// IsSubaccountUsage is true if subaccount used it, false if it is a money transfer
	IsSubaccountUsage bool `json:"is_subaccount_usage" bson:"is_subaccount_usage"`
}

// usage info struct
type UsageInfo struct {
	Charge      float64 `json:"charge" bson:"charge"`
	ChargeUnits int     `json:"charge_units" bson:"charge_units"`
	CostUnits   int     `json:"cost_units" bson:"cost_units"`
	CountryID   int     `json:"country_id" bson:"country_id"`
	CountryName string  `json:"country_name" bson:"country_name"`
	NetworkID   string  `json:"network_id" bson:"network_id"`
	CircleID    string  `json:"circle_id" bson:"circle_id"`
	NumberID    string  `json:"number_id" bson:"number_id"`
	NumberType  int     `json:"number_type" bson:"number_type"`
	CampaignID  string  `json:"campaign_id" bson:"campaign_id"`
	Price       float64 `json:"price" bson:"price"`
	Quantity    int     `json:"quantity" bson:"quantity"`
}
