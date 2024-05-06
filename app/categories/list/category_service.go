package main

type CategoryService struct {
	Store CategoryStore
}

func NewCategoryService(s CategoryStore) *CategoryService {
	return &CategoryService{Store: s}
}

func (cs *CategoryService) List() (string, error) {
	return "[  {\"id\": 1, \"name\": \"Groceries\"},  {\"id\": 2, \"name\": \"Transportation\"},  {\"id\": 3, \"name\": \"Utilities\"},  {\"id\": 4, \"name\": \"Entertainment\"},  {\"id\": 5, \"name\": \"Dining Out\"},  {\"id\": 6, \"name\": \"Healthcare\"},  {\"id\": 7, \"name\": \"Clothing\"},  {\"id\": 8, \"name\": \"Home Improvement\"},  {\"id\": 9, \"name\": \"Education\"},  {\"id\": 10, \"name\": \"Travel\"}]", nil
}
