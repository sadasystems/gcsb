// selector package provides a mechanism for selecting between things
package selector

// Selector provides an interface for selecting choices in various ways
type Selector interface {
	Select() Choice
}
