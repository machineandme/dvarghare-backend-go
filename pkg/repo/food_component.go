package repo

//go:generate reform

// FoodComponent represents a row in food_component table.
//reform:food_component
type FoodComponent struct {
	ID              *string `reform:"id"`
	FdcID           *string `reform:"fdc_id"`
	Name            *string `reform:"name"`
	PctWeight       *string `reform:"pct_weight"`
	IsRefuse        *string `reform:"is_refuse"`
	GramWeight      *string `reform:"gram_weight"`
	DataPoints      *string `reform:"data_points"`
	MinYearAcquired *string `reform:"min_year_acquired"`
}
