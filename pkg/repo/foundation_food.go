package repo

//go:generate reform

// FoundationFood represents a row in foundation_food table.
//reform:foundation_food
type FoundationFood struct {
	FdcID     *string `reform:"fdc_id"`
	NDBNumber *string `reform:"NDB_number"`
	Footnote  *string `reform:"footnote"`
}
