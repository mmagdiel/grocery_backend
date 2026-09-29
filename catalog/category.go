package catalog

import (
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

// ensures that the Category struct satisfies the core.RecordProxy interface
var _ core.RecordProxy = (*Category)(nil)

type Category struct {
	core.BaseRecordProxy
}

func (c *Category) Name() string {
	return c.GetString("name")
}

func (c *Category) SetName(name string) {
	c.Set("name", name)
}

func (c *Category) Slug() string {
	return c.GetString("slug")
}

func (c *Category) SetSlug(slug string) {
	c.Set("slug", slug)
}

func (c *Category) Description() string {
	return c.GetString("description")
}

func (c *Category) SetDescription(description string) {
	c.Set("description", description)
}

func (c *Category) ImageURL() string {
	return c.GetString("image_url")
}

func (c *Category) SetImageURL(url string) {
	c.Set("image_url", url)
}

func (c *Category) IsActive() bool {
	return c.GetBool("is_active")
}

func (c *Category) SetIsActive(active bool) {
	c.Set("is_active", active)
}

// FindCategoryBySlug finds a category by its slug (case insensitive).
func FindCategoryBySlug(app core.App, slug string) (*Category, error) {
	category := &Category{}
	err := app.RecordQuery("categories").
		AndWhere(dbx.NewExp("LOWER(slug)={:slug}", dbx.Params{
			"slug": strings.ToLower(slug),
		})).
		Limit(1).
		One(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// FindActiveCategories returns all active categories ordered by name.
func FindActiveCategories(app core.App) ([]*Category, error) {
	categories := []*Category{}
	err := app.RecordQuery("categories").
		AndWhere(dbx.HashExp{"is_active": true}).
		OrderBy("name ASC").
		All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
