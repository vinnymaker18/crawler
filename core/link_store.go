package core

type LinkStore interface {
	storeItem(LinkItem) (bool, error)
}
