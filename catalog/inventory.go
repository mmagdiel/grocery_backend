package catalog

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

// ensures that the Inventory struct satisfies the core.RecordProxy interface
var _ core.RecordProxy = (*Inventory)(nil)

// InventoryStatus represents the availability state of a product.
type InventoryStatus string

const (
	InventoryStatusInStock InventoryStatus = "in_stock"
	InventoryStatusLow     InventoryStatus = "low"
	InventoryStatusOut     InventoryStatus = "out"

	// LowStockThreshold is the quantity at or below which the status becomes "low".
	LowStockThreshold = 5
)

type Inventory struct {
	core.BaseRecordProxy
}

func (i *Inventory) ProductID() string {
	return i.GetString("product_id")
}

func (i *Inventory) SetProductID(id string) {
	i.Set("product_id", id)
}

func (i *Inventory) Quantity() int {
	return i.GetInt("quantity")
}

func (i *Inventory) SetQuantity(qty int) {
	i.Set("quantity", qty)
	i.recalculateStatus()
}

func (i *Inventory) Reserved() int {
	return i.GetInt("reserved")
}

func (i *Inventory) SetReserved(reserved int) {
	i.Set("reserved", reserved)
	i.recalculateStatus()
}

// Available is a computed value (quantity - reserved). It is NOT stored in the DB.
func (i *Inventory) Available() int {
	return i.Quantity() - i.Reserved()
}

func (i *Inventory) Status() InventoryStatus {
	return InventoryStatus(i.GetString("status"))
}

func (i *Inventory) SetStatus(status InventoryStatus) {
	i.Set("status", string(status))
}

func (i *Inventory) UpdatedAt() types.DateTime {
	return i.GetDateTime("updated")
}

// IsAvailableForSale returns true when at least one unit is available.
func (i *Inventory) IsAvailableForSale() bool {
	return i.Available() > 0 && i.Status() != InventoryStatusOut
}

// Reserve blocks units for a pending purchase intent.
// Call app.Save(inventory) after this to persist the change.
func (i *Inventory) Reserve(units int) {
	i.SetReserved(i.Reserved() + units)
}

// Release frees previously reserved units (e.g. intent cancelled).
// Call app.Save(inventory) after this to persist the change.
func (i *Inventory) Release(units int) {
	newReserved := i.Reserved() - units
	if newReserved < 0 {
		newReserved = 0
	}
	i.SetReserved(newReserved)
}

// recalculateStatus updates the status enum based on current available units.
// Called automatically by SetQuantity, SetReserved, Reserve and Release.
func (i *Inventory) recalculateStatus() {
	available := i.Quantity() - i.GetInt("reserved")
	switch {
	case available <= 0:
		i.Set("status", string(InventoryStatusOut))
	case available <= LowStockThreshold:
		i.Set("status", string(InventoryStatusLow))
	default:
		i.Set("status", string(InventoryStatusInStock))
	}
}

// FindInventoryByProduct returns the inventory record for a given product ID.
func FindInventoryByProduct(app core.App, productID string) (*Inventory, error) {
	inventory := &Inventory{}
	err := app.RecordQuery("inventory").
		AndWhere(dbx.HashExp{"product_id": productID}).
		Limit(1).
		One(inventory)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}
