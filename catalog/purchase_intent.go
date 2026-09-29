package catalog

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

// ensures that the PurchaseIntent struct satisfies the core.RecordProxy interface
var _ core.RecordProxy = (*PurchaseIntent)(nil)

// IntentStatus tracks where in the follow-up flow a purchase intent currently is.
type IntentStatus string

const (
	IntentStatusPending   IntentStatus = "pending"
	IntentStatusContacted IntentStatus = "contacted"
	IntentStatusClosed    IntentStatus = "closed"
)

type PurchaseIntent struct {
	core.BaseRecordProxy
}

func (pi *PurchaseIntent) ProductID() string {
	return pi.GetString("product_id")
}

func (pi *PurchaseIntent) SetProductID(id string) {
	pi.Set("product_id", id)
}

func (pi *PurchaseIntent) CustomerName() string {
	return pi.GetString("customer_name")
}

func (pi *PurchaseIntent) SetCustomerName(name string) {
	pi.Set("customer_name", name)
}

func (pi *PurchaseIntent) CustomerEmail() string {
	return pi.GetString("customer_email")
}

func (pi *PurchaseIntent) SetCustomerEmail(email string) {
	pi.Set("customer_email", email)
}

func (pi *PurchaseIntent) CustomerPhone() string {
	return pi.GetString("customer_phone")
}

func (pi *PurchaseIntent) SetCustomerPhone(phone string) {
	pi.Set("customer_phone", phone)
}

func (pi *PurchaseIntent) Quantity() int {
	return pi.GetInt("quantity")
}

func (pi *PurchaseIntent) SetQuantity(qty int) {
	pi.Set("quantity", qty)
}

func (pi *PurchaseIntent) Message() string {
	return pi.GetString("message")
}

func (pi *PurchaseIntent) SetMessage(message string) {
	pi.Set("message", message)
}

func (pi *PurchaseIntent) Status() IntentStatus {
	return IntentStatus(pi.GetString("status"))
}

func (pi *PurchaseIntent) SetStatus(status IntentStatus) {
	pi.Set("status", string(status))
}

// PriceAtIntent is a snapshot of the product price at intent creation time.
// Stored separately so price changes on the product don't affect historical records.
func (pi *PurchaseIntent) PriceAtIntent() float64 {
	return pi.GetFloat("price_at_intent")
}

func (pi *PurchaseIntent) SetPriceAtIntent(price float64) {
	pi.Set("price_at_intent", price)
}

func (pi *PurchaseIntent) CreatedAt() types.DateTime {
	return pi.GetDateTime("created")
}

// FindPendingIntents returns all intents not yet followed up, oldest first.
func FindPendingIntents(app core.App) ([]*PurchaseIntent, error) {
	intents := []*PurchaseIntent{}
	err := app.RecordQuery("purchase_intents").
		AndWhere(dbx.HashExp{"status": string(IntentStatusPending)}).
		OrderBy("created ASC").
		All(&intents)
	if err != nil {
		return nil, err
	}
	return intents, nil
}

// FindIntentsByProduct returns all intents for a given product, newest first.
func FindIntentsByProduct(app core.App, productID string) ([]*PurchaseIntent, error) {
	intents := []*PurchaseIntent{}
	err := app.RecordQuery("purchase_intents").
		AndWhere(dbx.HashExp{"product_id": productID}).
		OrderBy("created DESC").
		All(&intents)
	if err != nil {
		return nil, err
	}
	return intents, nil
}
