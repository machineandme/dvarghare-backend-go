package repo

//go:generate reform

// Nutrient represents a row in nutrient table.
//reform:nutrient
type Nutrient struct {
	ID          *string `reform:"id"`
	Name        *string `reform:"name"`
	UnitName    *string `reform:"unit_name"`
	NutrientNbr *string `reform:"nutrient_nbr"`
	Rank        *string `reform:"rank"`
}
