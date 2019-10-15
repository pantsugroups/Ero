package serializer

import (
	"eroauz/models"
)

type Category struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Type  int    `json:"type"`
	Count int    `json:"count"`
}

type CategoryResponse struct {
	Response
	Data Category `json:"data"`
}

type CategoryListResponse struct {
	Response
	Count int        `json:"count"`
	All   int        `json:"all"`
	Data  []Category `json:"data"`
	Next  bool       `json:"have_next"`
	Last  bool       `json:"have_last"`
	Pages int        `json:"pages"`
}

func BuildCategory(category models.Category) Category {
	return Category{
		ID:    category.ID,
		Title: category.Title,
		Type:  category.Type,
		Count: category.Count,
	}
}

func BuildCategoryList(categories []models.Category) []Category {
	var categoryList []Category
	for _, a := range categories {
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

func BuildCategoryListResponse(categories []models.Category, all int, count int, next bool, last bool, pages int) CategoryListResponse {
	return CategoryListResponse{
		Count: count,
		All:   all,
		Data:  BuildCategoryList(categories),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
