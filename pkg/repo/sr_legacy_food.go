package repo

//go:generate reform

// SrLegacyFood represents a row in sr_legacy_food table.
//reform:sr_legacy_food
type SrLegacyFood struct {
	FdcID     *string `reform:"fdc_id"`
	NDBNumber *string `reform:"NDB_number"`
}
