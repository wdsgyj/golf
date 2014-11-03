// hash
package hashmap

type Hash interface {
	HashCode() uint32
	Equals(Hash) bool
}

type MapEntry interface {
	Key() Hash
	Value() interface{}
}

type Map interface {
	Put(h Hash, v interface{}) (interface{}, error)

	Get(h Hash) (interface{}, error)

	Remove(h Hash) interface{}

	Contains(h Hash) bool

	Size() int
	IsEmpty() bool

	Clear()

	Travel(func(MapEntry) bool)
}