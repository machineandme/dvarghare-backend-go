package repo

//go:generate reform

// LabMethod represents a row in lab_method table.
//reform:lab_method
type LabMethod struct {
	ID          *string `reform:"id"`
	Description *string `reform:"description"`
	Technique   *string `reform:"technique"`
}
