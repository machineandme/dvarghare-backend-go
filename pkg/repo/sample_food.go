package repo

//go:generate reform

// SampleFood represents a row in sample_food table.
//reform:sample_food
type SampleFood struct {
	FdcID *string `reform:"fdc_id"`
}
