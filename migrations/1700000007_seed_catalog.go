package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

type seedCategory struct {
	Name        string
	Slug        string
	Description string
}

type seedProduct struct {
	Name        string
	Slug        string
	Category    string // category slug
	Price       float64
	Description string
	Image       string
	IsActive    bool
	IsFeatured  bool
	Unit        string
}

func init() {
	m.Register(func(app core.App) error {
		// ── Categories ──────────────────────────────────────────────
		categories := []seedCategory{
			{"Frutas", "frutas", "Frutas frescas de la mejor calidad"},
			{"Verduras", "verduras", "Verduras frescas y saludables"},
			{"Carnes", "carnes", "Carnes frescas y de primera"},
			{"Lácteos", "lacteos", "Productos lácteos frescos"},
			{"Panadería", "panaderia", "Pan y productos de panadería"},
			{"Despensa", "despensa", "Productos básicos de despensa"},
		}

		catCol, err := app.FindCollectionByNameOrId("categories")
		if err != nil {
			return err
		}

		// slug → record id
		catMap := make(map[string]string, len(categories))

		for _, c := range categories {
			rec := core.NewRecord(catCol)
			rec.Set("name", c.Name)
			rec.Set("slug", c.Slug)
			rec.Set("description", c.Description)
			rec.Set("is_active", true)
			if err := app.Save(rec); err != nil {
				return err
			}
			catMap[c.Slug] = rec.Id
		}

		// ── Products ────────────────────────────────────────────────
		products := []seedProduct{
			// Frutas
			{"Mango Tommy", "mango-tommy", "frutas", 3500, "Mangos frescos y dulces, perfectos para jugos o comer directamente.", "https://images.unsplash.com/photo-1553279768-865429fa0078?w=400&q=80", true, true, "kg"},
			{"Banano", "banano", "frutas", 2500, "Bananos maduros de la mejor calidad.", "https://images.unsplash.com/photo-1571771894821-ce9b6c11b08e?w=400&q=80", true, false, "kg"},
			{"Papaya", "papaya", "frutas", 4000, "Papaya fresca y dulce, rica en vitaminas.", "https://images.unsplash.com/photo-1617112848923-cc2234396a8d?w=400&q=80", true, false, "kg"},
			{"Piña", "pina", "frutas", 5000, "Piña madura y jugosa.", "https://images.unsplash.com/photo-1550258987-190a2d41a8ba?w=400&q=80", true, false, "unidad"},
			{"Guayaba", "guayaba", "frutas", 3000, "Guayabas frescas para jugos o postres.", "https://images.unsplash.com/photo-1536511132770-e5058c7e8c46?w=400&q=80", true, false, "kg"},
			// Verduras
			{"Tomate", "tomate", "verduras", 2800, "Tomates frescos y rojos, ideales para ensaladas.", "https://serreslavoie.com/cdn/shop/articles/shutterstock_267812054.jpg?crop=center&height=1023&v=1652289544&width=1540", true, true, "kg"},
			{"Cebolla Cabezona", "cebolla-cabezona", "verduras", 2200, "Cebolla fresca para tus preparaciones.", "https://images.unsplash.com/photo-1618512496248-a07fe83aa8cb?w=400&q=80", true, false, "kg"},
			{"Zanahoria", "zanahoria", "verduras", 2500, "Zanahorias frescas y crujientes.", "https://images.unsplash.com/photo-1598170845058-32b9d6a5da37?w=400&q=80", true, false, "kg"},
			{"Lechuga", "lechuga", "verduras", 2000, "Lechuga fresca para ensaladas.", "https://images.unsplash.com/photo-1622206151226-18ca2c9ab4a1?w=400&q=80", true, false, "unidad"},
			{"Papa Criolla", "papa-criolla", "verduras", 3500, "Papa criolla colombiana, perfecta para freír.", "https://images.unsplash.com/photo-1518977676601-b53f82aba655?w=400&q=80", true, false, "kg"},
			// Carnes
			{"Pechuga de Pollo", "pechuga-de-pollo", "carnes", 12000, "Pechuga de pollo fresca y sin hueso.", "https://images.unsplash.com/photo-1604503468506-a8da13d82791?w=400&q=80", true, true, "kg"},
			{"Carne Molida", "carne-molida", "carnes", 15000, "Carne molida de res, ideal para hamburguesas.", "https://images.unsplash.com/photo-1603048588665-791ca8aea617?w=400&q=80", true, false, "kg"},
			{"Costilla de Cerdo", "costilla-de-cerdo", "carnes", 18000, "Costillas de cerdo frescas para asar.", "https://images.unsplash.com/photo-1602470520998-f4a52199a3d6?w=400&q=80", true, false, "kg"},
			{"Chuleta de Cerdo", "chuleta-de-cerdo", "carnes", 16000, "Chuletas de cerdo jugosas.", "https://images.unsplash.com/photo-1607623814075-e51df1bdc82f?w=400&q=80", true, false, "kg"},
			{"Chorizo", "chorizo", "carnes", 14000, "Chorizo artesanal colombiano.", "https://images.unsplash.com/photo-1613564834361-9436948817d1?w=400&q=80", true, false, "kg"},
			// Lácteos
			{"Leche Entera", "leche-entera", "lacteos", 3500, "Leche fresca entera.", "https://images.unsplash.com/photo-1563636619-e9143da7973b?w=400&q=80", true, false, "litro"},
			{"Queso Campesino", "queso-campesino", "lacteos", 8000, "Queso campesino fresco y artesanal.", "https://images.unsplash.com/photo-1486297678162-eb2a19b0a32d?w=400&q=80", true, true, "500g"},
			{"Yogurt Natural", "yogurt-natural", "lacteos", 4500, "Yogurt natural sin azúcar.", "https://images.unsplash.com/photo-1488477181946-6428a0291777?w=400&q=80", true, false, "litro"},
			{"Mantequilla", "mantequilla", "lacteos", 6000, "Mantequilla cremosa.", "https://images.unsplash.com/photo-1589985270826-4b7bb135bc9d?w=400&q=80", true, false, "500g"},
			{"Cuajada", "cuajada", "lacteos", 7000, "Cuajada fresca colombiana.", "https://images.unsplash.com/photo-1628088062854-d1870b4553da?w=400&q=80", true, false, "500g"},
			// Panadería
			{"Pan Francés", "pan-frances", "panaderia", 500, "Pan francés recién horneado.", "https://images.unsplash.com/photo-1509440159596-0249088772ff?w=400&q=80", true, false, "unidad"},
			{"Pandebono", "pandebono", "panaderia", 1500, "Pandebono tradicional colombiano.", "https://images.unsplash.com/photo-1608198093002-ad4e005484ec?w=400&q=80", true, true, "unidad"},
			{"Almojábana", "almojabana", "panaderia", 1800, "Almojábana fresca de queso.", "https://images.unsplash.com/photo-1555507036-ab1f4038808a?w=400&q=80", true, false, "unidad"},
			{"Buñuelo", "bunuelo", "panaderia", 1200, "Buñuelos colombianos crujientes.", "https://media.istockphoto.com/id/1351588746/es/foto/bu%C3%B1uelos-cocina-navide%C3%B1a-colombiana-pan-de-queso-frito.jpg?s=612x612&w=0&k=20&c=pjkr5LdTuty41J4tKOgx2t1TWgoWzcZyeY5Tq13q0Bk=", true, false, "unidad"},
			{"Pan Integral", "pan-integral", "panaderia", 8000, "Pan integral artesanal.", "https://images.unsplash.com/photo-1586444248902-2f64eddc13df?w=400&q=80", true, false, "unidad"},
			// Despensa
			{"Arroz", "arroz", "despensa", 4500, "Arroz blanco de primera calidad.", "https://images.unsplash.com/photo-1586201375761-83865001e31c?w=400&q=80", true, false, "kg"},
			{"Frijol", "frijol", "despensa", 5000, "Frijol rojo colombiano.", "https://images.unsplash.com/photo-1589367920969-ab8e050bbb04?w=400&q=80", true, false, "kg"},
			{"Aceite de Cocina", "aceite-de-cocina", "despensa", 12000, "Aceite vegetal para cocinar.", "https://images.unsplash.com/photo-1474979266404-7eaacbcd87c5?w=400&q=80", true, false, "litro"},
			{"Panela", "panela", "despensa", 3000, "Panela colombiana pura.", "https://images.unsplash.com/photo-1587735243615-c03f25aaff15?w=400&q=80", true, false, "500g"},
			{"Café Colombiano", "cafe-colombiano", "despensa", 15000, "Café colombiano premium molido.", "https://images.unsplash.com/photo-1447933601403-0c6688de566e?w=400&q=80", true, true, "500g"},
		}

		prodCol, err := app.FindCollectionByNameOrId("products")
		if err != nil {
			return err
		}
		invCol, err := app.FindCollectionByNameOrId("inventory")
		if err != nil {
			return err
		}

		for _, p := range products {
			catID, ok := catMap[p.Category]
			if !ok {
				continue
			}

			rec := core.NewRecord(prodCol)
			rec.Set("category_id", catID)
			rec.Set("name", p.Name)
			rec.Set("slug", p.Slug)
			rec.Set("description", p.Description)
			rec.Set("price", p.Price)
			rec.Set("is_active", p.IsActive)

			if err := app.Save(rec); err != nil {
				return err
			}

			// Create inventory record (default 50 units, in_stock)
			inv := core.NewRecord(invCol)
			inv.Set("product_id", rec.Id)
			inv.Set("quantity", 50)
			inv.Set("reserved", 0)
			inv.Set("status", "in_stock")
			if err := app.Save(inv); err != nil {
				return err
			}
		}

		return nil
	}, func(app core.App) error {
		// Down: remove seeded data in reverse order (inventory → products → categories)
		for _, col := range []string{"inventory", "products", "categories"} {
			records, err := app.FindAllRecords(col)
			if err != nil {
				continue
			}
			for _, r := range records {
				_ = app.Delete(r)
			}
		}
		return nil
	})
}
