package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		collection := core.NewBaseCollection("categories")

		// API rules: public read, authenticated write
		collection.ListRule = types.Pointer("")
		collection.ViewRule = types.Pointer("")
		collection.CreateRule = types.Pointer("@request.auth.id != ''")
		collection.UpdateRule = types.Pointer("@request.auth.id != ''")
		collection.DeleteRule = types.Pointer("@request.auth.id != ''")

		collection.Fields.Add(
			&core.TextField{
				Name:     "name",
				Required: true,
				Max:      120,
			},
			&core.TextField{
				Name:     "slug",
				Required: true,
				Max:      120,
			},
			&core.TextField{
				Name: "description",
				Max:  500,
			},
			&core.URLField{
				Name: "image_url",
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

		collection.AddIndex("idx_categories_slug", true, "slug", "")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("categories")
		if err != nil {
			return err
		}
		return app.Delete(collection)
	})
}
