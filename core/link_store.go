package core

type LinkStore interface {
	StoreItems([]LinkItem) (bool, error)
}
