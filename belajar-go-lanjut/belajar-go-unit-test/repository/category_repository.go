package repository

import "github.com/alanyukeroo/belajar-go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
