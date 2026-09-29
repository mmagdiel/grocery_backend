# pb_catalog

PocketBase backend for the product catalog application.

## Project structure

```
pb_catalog/
├── main.go                  # Entry point — registers migratecmd and starts the server
├── go.mod
├── catalog/                 # Record proxy types (typed access to PocketBase records)
│   ├── category.go
│   ├── product.go
│   ├── inventory.go
│   ├── purchase_intent.go
│   └── user.go
└── migrations/              # DB migrations — run automatically on server start
    ├── 1700000001_create_categories.go
    ├── 1700000002_create_products.go
    ├── 1700000003_create_inventory.go
    ├── 1700000004_create_purchase_intents.go
    ├── 1700000005_create_users.go
    └── 1700000006_initial_settings.go
```

## Requirements

- Go 1.22+
- Internet access on first run (to download the PocketBase dependency)

## Setup

```bash
# 1. Download dependencies
go mod tidy

# 2. Start the development server
#    Automigrate is active during "go run" — migrations apply automatically
go run . serve

# 3. Open the admin dashboard
#    http://localhost:8090/_/
```

On first launch PocketBase will ask you to create a superuser account via the dashboard.
All catalog users (admin/editor roles) are managed separately in the `users` collection.

## Useful commands

```bash
# Apply pending migrations manually
go run . migrate up

# Revert the last migration
go run . migrate down 1

# Create a new blank migration
go run . migrate create "your_migration_name"

# Generate a snapshot of the current collections state
go run . migrate collections

# After squashing intermediate dev migrations, sync the history table
go run . migrate history-sync
```

## Collections & API rules

| Collection       | List / View       | Create            | Update / Delete   |
|------------------|-------------------|-------------------|-------------------|
| categories       | public            | authenticated     | authenticated     |
| products         | public            | authenticated     | authenticated     |
| inventory        | authenticated     | authenticated     | authenticated     |
| purchase_intents | authenticated     | **public**        | authenticated     |
| users            | authenticated     | superuser only    | own record        |

`purchase_intents.CreateRule` is intentionally public so customers can submit
a purchase intent without registering.

## Inventory logic

The `inventory` proxy exposes `Reserve(n)` and `Release(n)` helpers that adjust
the `reserved` counter and keep the `status` enum (`in_stock` / `low` / `out`)
in sync automatically. Always call `app.Save(inventory)` after mutating:

```go
inv, err := catalog.FindInventoryByProduct(app, productID)
if err != nil { ... }

inv.Reserve(intentQuantity)

if err := app.Save(inv); err != nil { ... }
```
