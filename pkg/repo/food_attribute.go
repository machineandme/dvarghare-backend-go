package repo

//go:generate reform

// FoodAttribute represents a row in food_attribute table.
//reform:food_attribute
type FoodAttribute struct {
	ID                  *string `reform:"id"`
	FdcID               *string `reform:"fdc_id"`
	SeqNum              *string `reform:"seq_num"`
	FoodAttributeTypeID *string `reform:"food_attribute_type_id"`
	Name                *string `reform:"name"`
	Value               *string `reform:"value"`
}
