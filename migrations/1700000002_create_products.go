package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		categoriesCollection, err := app.FindCollectionByNameOrId("categories")
		if err != nil {
			return err
		}

		collection := core.NewBaseCollection("products")

		// API rules: public read, authenticated write
		collection.ListRule = types.Pointer("")
		collection.ViewRule = types.Pointer("")
		collection.CreateRule = types.Pointer("@request.auth.id != ''")
		collection.UpdateRule = types.Pointer("@request.auth.id != ''")
		collection.DeleteRule = types.Pointer("@request.auth.id != ''")

		collection.Fields.Add(
			&core.RelationField{
				Name:         "category_id",
				Required:     true,
				CollectionId: categoriesCollection.Id,
				// one product belongs to one category
				MaxSelect: 1,
			},
			&core.TextField{
				Name:     "name",
				Required: true,
				Max:      200,
			},
			&core.TextField{
				Name:     "slug",
				Required: true,
				Max:      200,
			},
			&core.EditorField{
				Name: "description",
			},
			&core.NumberField{
				Name:     "price",
				Required: true,
				Min:      types.Pointer(0.0),
			},
			&core.FileField{
				Name:      "images",
				MaxSelect: 10,
				MaxSize:   5 * 1024 * 1024, // 5MB per file
				MimeTypes: []string{
					"image/jpeg",
					"image/png",
					"image/webp",
					"image/gif",
				},
			},
			&core.BoolField{
				Name: "is_active",
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

		collection.AddIndex("idx_products_slug", true, "slug", "")
		collection.AddIndex("idx_products_category", false, "category_id", "")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("products")
		if err != nil {
			return err
		}
		return app.Delete(collection)
	})
}
