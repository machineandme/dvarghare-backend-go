package repo

//go:generate reform

// SubSampleResult represents a row in sub_sample_result table.
//reform:sub_sample_result
type SubSampleResult struct {
	FoodNutrientID *string `reform:"food_nutrient_id"`
	AdjustedAmount *string `reform:"adjusted_amount"`
	LabMethodID    *string `reform:"lab_method_id"`
	NutrientName   *string `reform:"nutrient_name"`
}
