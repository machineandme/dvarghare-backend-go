package repo

//go:generate reform

// FoodPortion represents a row in food_portion table.
//reform:food_portion
type FoodPortion struct {
	ID                 *string `reform:"id"`
	FdcID              *string `reform:"fdc_id"`
	SeqNum             *string `reform:"seq_num"`
	Amount             *string `reform:"amount"`
	MeasureUnitID      *string `reform:"measure_unit_id"`
	PortionDescription *string `reform:"portion_description"`
	Modifier           *string `reform:"modifier"`
	GramWeight         *string `reform:"gram_weight"`
	DataPoints         *string `reform:"data_points"`
	Footnote           *string `reform:"footnote"`
	MinYearAcquired    *string `reform:"min_year_acquired"`
}
