package restclient

type Algorithm interface {
	Prefix() string
	Name() string
	Exec(interface{}) interface{}
}
