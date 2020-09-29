package repo

//go:generate reform

// FoodNutrient represents a row in food_nutrient table.
//reform:food_nutrient
type FoodNutrient struct {
	ID              *string `reform:"id"`
	FdcID           *string `reform:"fdc_id"`
	NutrientID      *string `reform:"nutrient_id"`
	Amount          *string `reform:"amount"`
	DataPoints      *string `reform:"data_points"`
	DerivationID    *string `reform:"derivation_id"`
	Min             *string `reform:"min"`
	Max             *string `reform:"max"`
	Median          *string `reform:"median"`
	Footnote        *string `reform:"footnote"`
	MinYearAcquired *string `reform:"min_year_acquired"`
}
