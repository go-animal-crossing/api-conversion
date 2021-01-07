package pages

// CopyFromFirstItem sets title, h1 and similar from the `.Item` property
func (pg *Page) CopyFromFirstItem() {
	pg.ID = pg.Items[0].ID
	pg.Title = pg.Items[0].Title
	pg.H1 = pg.Items[0].Title
	pg.Template = pg.Items[0].Template
	pg.URI = pg.Items[0].URI
	pg.Has = pg.Items[0].Has
	pg.Type = pg.Items[0].Type

}

// GetURL returns url
func (pg *Page) GetURL() string {
	if pg.IsList {
		return pg.URI
	}
	return pg.Items[0].URI
}
