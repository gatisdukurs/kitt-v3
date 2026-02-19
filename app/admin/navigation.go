package admin

type NavigationItem struct {
	Label  string
	Path   string
	Active bool
}

type Navigation struct {
	Items []NavigationItem
}

func (n Navigation) WithActive(path string) Navigation {
	n.Items = append([]NavigationItem(nil), n.Items...)

	for i := range n.Items {
		n.Items[i].Active = false
		if n.Items[i].Path == path {
			n.Items[i].Active = true
		}
	}

	return n
}
