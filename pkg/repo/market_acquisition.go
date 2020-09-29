package repo

//go:generate reform

// MarketAcquisition represents a row in market_acquisition table.
//reform:market_acquisition
type MarketAcquisition struct {
	FdcID            *string `reform:"fdc_id"`
	BrandDescription *string `reform:"brand_description"`
	ExpirationDate   *string `reform:"expiration_date"`
	LabelWeight      *string `reform:"label_weight"`
	Location         *string `reform:"location"`
	AcquisitionDate  *string `reform:"acquisition_date"`
	SalesType        *string `reform:"sales_type"`
	SampleLotNbr     *string `reform:"sample_lot_nbr"`
	SellByDate       *string `reform:"sell_by_date"`
	StoreCity        *string `reform:"store_city"`
	StoreName        *string `reform:"store_name"`
	StoreState       *string `reform:"store_state"`
	UpcCode          *string `reform:"upc_code"`
}
