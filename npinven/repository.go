package NPINven

type Repository interface {
	GenDocNoInven(Type string, Search string, Branch string) (interface{}, error)
}
