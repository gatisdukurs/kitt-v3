package pages

type Page struct {
	ID      int64  `db:"id, pk, auto"`
	Title   string `db:"title, notnull"`
	Content string `db:"content, notnull"`
}
