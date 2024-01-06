package infra

import "github.com/todoApp/api/domain/repo"

type todoInfra struct{}

func NewTODO() repo.TODO {
	return &todoInfra{}
}

func (td *todoInfra) TODOFunc() {
	// 処理を宣言
}
