package jobs

type Ajob struct {
	SomeVar string
}

func (Ajob) Id() string {
	return "admin.pages.ajob"
}
