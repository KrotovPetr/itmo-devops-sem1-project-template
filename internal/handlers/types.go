package handlers

type PriceStatsStruct struct {
	TotalCount      int `json:"total_count"`
	DuplicateCount  int `json:"duplicates_count"`
	TotalItems      int `json:"total_items"`
	TotalCategories int `json:"total_categories"`
	TotalPrice      int `json:"total_price"`
}
