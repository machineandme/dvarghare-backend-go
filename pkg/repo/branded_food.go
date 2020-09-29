package repo

//go:generate reform

// BrandedFood represents a row in branded_food table.
//reform:branded_food
type BrandedFood struct {
	FdcID                    *string `reform:"fdc_id"`
	BrandOwner               *string `reform:"brand_owner"`
	GtinUpc                  *string `reform:"gtin_upc"`
	Ingredients              *string `reform:"ingredients"`
	ServingSize              *string `reform:"serving_size"`
	ServingSizeUnit          *string `reform:"serving_size_unit"`
	HouseholdServingFulltext *string `reform:"household_serving_fulltext"`
	BrandedFoodCategory      *string `reform:"branded_food_category"`
	DataSource               *string `reform:"data_source"`
	ModifiedDate             *string `reform:"modified_date"`
	AvailableDate            *string `reform:"available_date"`
}
