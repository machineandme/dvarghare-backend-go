package repo

//go:generate reform

// RetentionFactor represents a row in retention_factor table.
//reform:retention_factor
type RetentionFactor struct {
	ID          *string `reform:"id"`
	Code        *string `reform:"code"`
	FoodGroupID *string `reform:"food_group_id"`
	Description *string `reform:"description"`
}
