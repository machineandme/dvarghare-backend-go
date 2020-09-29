package repo

//go:generate reform

// AcquisitionSample represents a row in acquisition_sample table.
//reform:acquisition_sample
type AcquisitionSample struct {
	FdcIDOfSampleFood      *string `reform:"fdc_id_of_sample_food"`
	FdcIDOfAcquisitionFood *string `reform:"fdc_id_of_acquisition_food"`
}
