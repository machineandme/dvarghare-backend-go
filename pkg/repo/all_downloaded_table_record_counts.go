package repo

//go:generate reform

// AllDownloadedTableRecordCounts represents a row in all_downloaded_table_record_counts table.
//reform:all_downloaded_table_record_counts
type AllDownloadedTableRecordCounts struct {
	Table           *string `reform:"Table"`
	NumberOfRecords *string `reform:"Number of Records"`
}
