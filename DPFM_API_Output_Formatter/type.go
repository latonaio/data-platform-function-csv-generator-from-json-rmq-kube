package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey          string      `json:"connection_key"`
	Result                 bool        `json:"result"`
	RedisKey               string      `json:"redis_key"`
	Filepath               string      `json:"filepath"`
	APIStatusCode          int         `json:"api_status_code"`
	RuntimeSessionID       string      `json:"runtime_session_id"`
	BusinessPartnerID      *int        `json:"business_partner"`
	ServiceLabel           string      `json:"service_label"`
	APIType                string      `json:"api_type"`
	Message                interface{} `json:"message"`
	APISchema              string      `json:"api_schema"`
	Accepter               []string    `json:"accepter"`
	Deleted                bool        `json:"deleted"`
	APIProcessingResult    *bool       `json:"api_processing_result"`
	APIProcessingError     string      `json:"api_processing_error"`
	DeliveryInstructionPdf string      `json:"delivery_instruction_pdf"`
	MountPath              *string     `json:"mount_path"`
}

type Message struct {
	Header  *[]Header `json:"Header"`
	Records *string   `json:"Records"`
}

type Header struct {
	Header string `json:"Header"`
}
