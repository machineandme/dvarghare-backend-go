package repo

//go:generate reform

// LabMethodNutrient represents a row in lab_method_nutrient table.
//reform:lab_method_nutrient
type LabMethodNutrient struct {
	ID          *string `reform:"id"`
	LabMethodID *string `reform:"lab_method_id"`
	NutrientID  *string `reform:"nutrient_id"`
}
