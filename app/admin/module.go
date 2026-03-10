package admin

import (
	"kitt/app/admin/internal/dashboard"
	"kitt/app/admin/internal/pages"
)

type Module struct{}

func (m Module) Boot() {
	dashboard.Controller{}.Boot()
	pages := &pages.Controller{}
	pages.Boot()

	// Migrate
	m.migrate()
}

func (Module) migrate() {
	// _, err := kitt.SQL().Exec(context.Background(), `
	// 	CREATE TABLE IF NOT EXISTS pages (
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		title TEXT NOT NULL,
	// 		content TEXT NOT NULL
	// 	);
	// `)

	// if err != nil {
	// 	fmt.Println(err)
	// }
}
