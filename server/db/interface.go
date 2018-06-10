package db

type RepositoryDef interface {
	GetName() string
	GetHashKey() string
}

type Repository interface {
	FindById(id string) (interface{}, error)
	FindByQuery(query string, value string) (interface{}, error)
}

type RepositoryDefMap map[string]interface{}

func (m RepositoryDefMap) GetName() string {
	if name, ok := m["name"]; ok {
		return name.(string)
	}
	return ""
}

func (m RepositoryDefMap) GetHashKey() string {
	if name, ok := m["hashKey"]; ok {
		return name.(string)
	}
	return ""
}
