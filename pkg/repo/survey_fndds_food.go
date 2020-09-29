package repo

//go:generate reform

// SurveyFnddsFood represents a row in survey_fndds_food table.
//reform:survey_fndds_food
type SurveyFnddsFood struct {
	FdcID             *string `reform:"fdc_id"`
	FoodCode          *string `reform:"food_code"`
	WweiaCategoryCode *string `reform:"wweia_category_code"`
	StartDate         *string `reform:"start_date"`
	EndDate           *string `reform:"end_date"`
}
