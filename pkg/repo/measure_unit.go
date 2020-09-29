package repo

//go:generate reform

// MeasureUnit represents a row in measure_unit table.
//reform:measure_unit
type MeasureUnit struct {
	ID   *string `reform:"id"`
	Name *string `reform:"name"`
}
