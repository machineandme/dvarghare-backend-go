package repo

//go:generate reform

// SubSampleFood represents a row in sub_sample_food table.
//reform:sub_sample_food
type SubSampleFood struct {
	FdcID             *string `reform:"fdc_id"`
	FdcIDOfSampleFood *string `reform:"fdc_id_of_sample_food"`
}
