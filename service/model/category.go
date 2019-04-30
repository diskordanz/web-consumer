package model

import (
	pkgCategoryModel "github.com/diskordanz/consumer/pkg/category/model"
)

type Category struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func MapToCategory(originCategory pkgCategoryModel.Category) Category {
	return Category{ID: originCategory.ID, Name: originCategory.Name}
}

func MapToCategoryList(originList []pkgCategoryModel.Category) []Category {
	resultList := make([]Category, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToCategory(v)
	}
	return resultList
}
