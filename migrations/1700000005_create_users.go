package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		// PocketBase may auto-create a default "users" auth collection,
		// so reuse it if it already exists; otherwise create a new one.
		collection, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			collection = core.NewAuthCollection("users")
		}

		// Only the user themselves or a superuser can read/manage their record
		collection.ListRule = types.Pointer("@request.auth.id != ''")
		collection.ViewRule = types.Pointer("@request.auth.id = id")
		collection.CreateRule = nil // only superusers can create users (no self-register)
		collection.UpdateRule = types.Pointer("@request.auth.id = id")
		collection.DeleteRule = nil // only superusers can delete users

		collection.Fields.Add(
			&core.TextField{
				Name:     "display_name",
				Required: true,
				Max:      100,
			},
			&core.SelectField{
				Name:   "role",
				Values: []string{"admin", "editor"},
			},
		)

		// Disable OAuth2 and keep only password auth for the admin panel
		collection.OAuth2.Enabled = false
		collection.PasswordAuth.Enabled = true

		collection.AddIndex("idx_users_role", false, "role", "")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		return app.Delete(collection)
	})
}
