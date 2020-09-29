package repo

//go:generate reform

// FoodUpdateLogEntry represents a row in food_update_log_entry table.
//reform:food_update_log_entry
type FoodUpdateLogEntry struct {
	ID          *string `reform:"id"`
	Description *string `reform:"description"`
	LastUpdated *string `reform:"last_updated"`
}
