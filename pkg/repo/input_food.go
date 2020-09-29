package repo

//go:generate reform

// InputFood represents a row in input_food table.
//reform:input_food
type InputFood struct {
	ID                 *string `reform:"id"`
	FdcID              *string `reform:"fdc_id"`
	FdcIDOfInputFood   *string `reform:"fdc_id_of_input_food"`
	SeqNum             *string `reform:"seq_num"`
	Amount             *string `reform:"amount"`
	SrCode             *string `reform:"sr_code"`
	SrDescription      *string `reform:"sr_description"`
	Unit               *string `reform:"unit"`
	PortionCode        *string `reform:"portion_code"`
	PortionDescription *string `reform:"portion_description"`
	GramWeight         *string `reform:"gram_weight"`
	RetentionCode      *string `reform:"retention_code"`
	SurveyFlag         *string `reform:"survey_flag"`
}
