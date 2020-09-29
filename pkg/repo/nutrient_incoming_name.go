package repo

//go:generate reform

// NutrientIncomingName represents a row in nutrient_incoming_name table.
//reform:nutrient_incoming_name
type NutrientIncomingName struct {
	ID         *string `reform:"id"`
	Name       *string `reform:"name"`
	NutrientID *string `reform:"nutrient_id"`
}
