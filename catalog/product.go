package catalog

import (
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

// ensures that the Product struct satisfies the core.RecordProxy interface
var _ core.RecordProxy = (*Product)(nil)

type Product struct {
	core.BaseRecordProxy
}

func (p *Product) CategoryID() string {
	return p.GetString("category_id")
}

func (p *Product) SetCategoryID(id string) {
	p.Set("category_id", id)
}

func (p *Product) Name() string {
	return p.GetString("name")
}

func (p *Product) SetName(name string) {
	p.Set("name", name)
}

func (p *Product) Slug() string {
	return p.GetString("slug")
}

func (p *Product) SetSlug(slug string) {
	p.Set("slug", slug)
}

func (p *Product) Description() string {
	return p.GetString("description")
}

func (p *Product) SetDescription(description string) {
	p.Set("description", description)
}

func (p *Product) Price() float64 {
	return p.GetFloat("price")
}

func (p *Product) SetPrice(price float64) {
	p.Set("price", price)
}

// Images returns the list of uploaded image filenames.
func (p *Product) Images() []string {
	return p.GetStringSlice("images")
}

func (p *Product) SetImages(images interface{}) {
	p.Set("images", images)
}

func (p *Product) IsActive() bool {
	return p.GetBool("is_active")
}

func (p *Product) SetIsActive(active bool) {
	p.Set("is_active", active)
}

func (p *Product) CreatedAt() types.DateTime {
	return p.GetDateTime("created")
}

// FindProductBySlug finds a product by its slug (case insensitive).
func FindProductBySlug(app core.App, slug string) (*Product, error) {
	product := &Product{}
	err := app.RecordQuery("products").
		AndWhere(dbx.NewExp("LOWER(slug)={:slug}", dbx.Params{
			"slug": strings.ToLower(slug),
		})).
		Limit(1).
		One(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// FindProductsByCategory returns all active products for a given category ID.
func FindProductsByCategory(app core.App, categoryID string) ([]*Product, error) {
	products := []*Product{}
	err := app.RecordQuery("products").
		AndWhere(dbx.HashExp{
			"category_id": categoryID,
			"is_active":   true,
		}).
		OrderBy("name ASC").
		All(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
