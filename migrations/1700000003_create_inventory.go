package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		productsCollection, err := app.FindCollectionByNameOrId("products")
		if err != nil {
			return err
		}

		collection := core.NewBaseCollection("inventory")

		// Only authenticated users (admins/editors) can read or write inventory
		collection.ListRule = types.Pointer("@request.auth.id != ''")
		collection.ViewRule = types.Pointer("@request.auth.id != ''")
		collection.CreateRule = types.Pointer("@request.auth.id != ''")
		collection.UpdateRule = types.Pointer("@request.auth.id != ''")
		collection.DeleteRule = types.Pointer("@request.auth.id != ''")

		collection.Fields.Add(
			&core.RelationField{
				Name:          "product_id",
				Required:      true,
				CollectionId:  productsCollection.Id,
				MaxSelect:     1,
				CascadeDelete: true, // delete inventory when the product is deleted
			},
			&core.NumberField{
				Name:     "quantity",
				Required: true,
				Min:      types.Pointer(0.0),
			},
			&core.NumberField{
				Name: "reserved",
				Min:  types.Pointer(0.0),
			},
			// status is kept in sync programmatically via the Inventory proxy methods
			&core.SelectField{
				Name:   "status",
				Values: []string{"in_stock", "low", "out"},
			},
			&core.AutodateField{
				Name:     "created",
				OnCreate: true,
			},
			&core.AutodateField{
				Name:     "updated",
				OnCreate: true,
				OnUpdate: true,
			},
		)

		// one inventory record per product
		collection.AddIndex("idx_inventory_product", true, "product_id", "")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("inventory")
		if err != nil {
			return err
		}
		return app.Delete(collection)
	})
}
