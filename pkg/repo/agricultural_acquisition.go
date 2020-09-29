package repo

//go:generate reform

// AgriculturalAcquisition represents a row in agricultural_acquisition table.
//reform:agricultural_acquisition
type AgriculturalAcquisition struct {
	FdcID           *string `reform:"fdc_id"`
	AcquisitionDate *string `reform:"acquisition_date"`
	MarketClass     *string `reform:"market_class"`
	Treatment       *string `reform:"treatment"`
	State           *string `reform:"state"`
}
