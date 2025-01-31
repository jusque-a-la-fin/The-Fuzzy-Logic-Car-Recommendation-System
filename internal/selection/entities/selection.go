package entities

type Selection struct {
	// Priorities - приоритеты или нечеткие множества, например,
	// "Экономичность", "Комфорт", "Управляемость", "Динамика", "Безопасность"
	Priorities []string `json:"priorities"`
	// MinPrice - нижний предел цены
	MinPrice string `json:"minPrice"`
	// MaxPrice - верхний предел цены
	MaxPrice string `json:"maxPrice"`
	// Manufacturers - страны-производители
	Manufacturers []string `json:"manufacturers"`
}
