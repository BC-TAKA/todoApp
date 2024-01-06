package service

import "github.com/todoApp/api/domain/repo"

type TODOService interface {
	// 関数を宣言
	TODOFunc()
}

type todoService struct {
	todoRepo repo.TODO
}

func NewTODOService(value repo.TODO) TODOService {
	return &todoService{
		todoRepo: value,
	}
}

func (td *todoService) TODOFunc() {
	// 処理を記載
}
