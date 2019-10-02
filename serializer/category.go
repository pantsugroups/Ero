package serializer

import (
	"eroauz/models"
)

type Category struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Count int    `json:"count"`
}

type CategoryResponse struct {
	Response
	Data Category `json:"data"`
}

type CategoryListResponse struct {
	Response
	Count int        `json:"count"`
	Data  []Category `json:"data"`
	Next  bool       `json:"have_next"`
	Last  bool       `json:"have_last"`
	Pages int        `json:"pages"`
}

func BuildCategory(category models.Category) Category {
	return Category{
		ID:    category.ID,
		Title: category.Title,
		Count: category.Count,
	}
}

func BuildCategoryList(categorys []models.Category) []Category {
	var categoryList []Category
	for _, a := range categorys {
		i := BuildCategory(a)
		categoryList = append(categoryList, i)
	}
	return categoryList
}

func BuildCategoryResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		Data: BuildCategory(category),
	}
}

func BuildCategoryListResponse(categorys []models.Category, count int, next bool, last bool, pages int) CategoryListResponse {
	return CategoryListResponse{
		Count: count,
		Data:  BuildCategoryList(categorys),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
