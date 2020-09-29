package repo

//go:generate reform

// Food represents a row in food table.
//reform:food
type Food struct {
	FdcID           *string `reform:"fdc_id"`
	DataType        *string `reform:"data_type"`
	Description     *string `reform:"description"`
	FoodCategoryID  *string `reform:"food_category_id"`
	PublicationDate *string `reform:"publication_date"`
}
