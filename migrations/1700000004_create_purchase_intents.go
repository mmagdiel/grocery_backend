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

		collection := core.NewBaseCollection("purchase_intents")

		// Public customers can create intents (no auth required).
		// Only authenticated users can list, view, update or delete them.
		collection.ListRule = types.Pointer("@request.auth.id != ''")
		collection.ViewRule = types.Pointer("@request.auth.id != ''")
		collection.CreateRule = types.Pointer("")
		collection.UpdateRule = types.Pointer("@request.auth.id != ''")
		collection.DeleteRule = types.Pointer("@request.auth.id != ''")

		collection.Fields.Add(
			&core.RelationField{
				Name:         "product_id",
				Required:     true,
				CollectionId: productsCollection.Id,
				MaxSelect:    1,
			},
			&core.TextField{
				Name:     "customer_name",
				Required: true,
				Max:      150,
			},
			&core.EmailField{
				Name:     "customer_email",
				Required: true,
			},
			&core.TextField{
				Name: "customer_phone",
				Max:  30,
			},
			&core.NumberField{
				Name:     "quantity",
				Required: true,
				Min:      types.Pointer(1.0),
			},
			&core.TextField{
				Name: "message",
				Max:  1000,
			},
			&core.SelectField{
				Name:   "status",
				Values: []string{"pending", "contacted", "closed"},
			},
			// price snapshot at the time the intent was created
			&core.NumberField{
				Name:     "price_at_intent",
				Required: true,
				Min:      types.Pointer(0.0),
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

		collection.AddIndex("idx_intents_product", false, "product_id", "")
		collection.AddIndex("idx_intents_status", false, "status", "")
		collection.AddIndex("idx_intents_email", false, "customer_email", "")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("purchase_intents")
		if err != nil {
			return err
		}
		return app.Delete(collection)
	})
}
