package repo

//go:generate reform

// LabMethodCode represents a row in lab_method_code table.
//reform:lab_method_code
type LabMethodCode struct {
	ID          *string `reform:"id"`
	LabMethodID *string `reform:"lab_method_id"`
	Code        *string `reform:"code"`
}
