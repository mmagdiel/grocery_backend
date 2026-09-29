package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		settings := app.Settings()

		settings.Meta.AppName = "Product Catalog"
		settings.Meta.AppURL = "http://localhost:8090"
		settings.Logs.MaxDays = 7
		settings.Logs.LogAuthId = true
		settings.Logs.LogIP = false

		return app.Save(settings)
	}, func(app core.App) error {
		// settings reset is not needed on down
		return nil
	})
}
