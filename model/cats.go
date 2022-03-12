package model

//Category struct
type Category struct {
	ID   int
	Name string
}

//WriteCategories inserts new categories into database, or replaces the exist one
func (c *Category) WriteCategories() error {
	row := Db.QueryRow(`INSERT OR REPLACE INTO categories 
	(id, name) VALUES(
		(SELECT Id FROM categories WHERE categories.name = ?), ?) RETURNING id
`, c.Name, c.Name)

	err := row.Scan(&c.ID)
	if err != nil {
		return err
	}

	return nil
}

//WriteCatPost writes relation between categories and posts into database
func (c *Category) WriteCatPost(id int) error {
	_, err := Db.Exec(`INSERT INTO cat_posts (post_id, cat_id) VALUES(?, ?)`, id, c.ID)

	if err != nil {
		return err
	}

	return nil
}

//WriteCats writes categories for the post in database
func (f *Forum) WriteCats(cats []Category) error {
	for _, c := range cats {
		if err := c.WriteCategories(); err != nil {
			return err
		}
		if err := c.WriteCatPost(f.Post.ID); err != nil {
			return err
		}
	}
	return nil
}

//EditCats updates categories for the post
func (f *Forum) EditCats(cats []Category) error {
	if err := DeleteCatPost(f.Post.ID); err != nil {
		return err
	}

	for _, c := range cats {
		if err := c.EditCategories(); err != nil {
			return err
		}

		if err := c.EditCatPost(f.Post.ID); err != nil {
			return err
		}

	}
	return nil
}

//DeleteCatPost deletes cat_posts from db
func DeleteCatPost(id int) error {
	_, err := Db.Exec(`DELETE FROM cat_posts WHERE post_id = ?`, id)

	if err != nil {
		return err
	}

	return nil
}

//EditCategories updates new categories into database, or replaces the exist one
func (c *Category) EditCategories() error {
	row := Db.QueryRow(`INSERT OR REPLACE INTO categories 
	(id, name) VALUES(
		(SELECT id FROM categories WHERE categories.name = ?), ?) RETURNING id
`, c.Name, c.Name)

	err := row.Scan(&c.ID)
	if err != nil {
		return err
	}

	return nil
}

//EditCatPost updates relation between categories and posts into database
func (c *Category) EditCatPost(id int) error {
	_, err := Db.Exec(`INSERT INTO cat_posts (post_id, cat_id) VALUES(?, ?)`, id, c.ID)

	if err != nil {
		return err
	}

	return nil
}
