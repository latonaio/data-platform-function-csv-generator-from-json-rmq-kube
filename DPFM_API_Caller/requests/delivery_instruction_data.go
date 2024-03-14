package requests

type Header struct {
	DeliveryDocument                       int      `json:"DeliveryDocument"`
	DeliveryDocumentDate                   string   `json:"DeliveryDocumentDate"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverToPartyName					   string   `json:"DeliverToPartyName"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverToPlantName					   string   `json:"DeliverToPlantName"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverFromPartyName				   string   `json:"DeliverFromPartyName"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	DeliverFromPlantName				   string   `json:"DeliverFromPlantName"`
	IsExportImport                         *bool    `json:"IsExportImport"`
	OrderID                                *int     `json:"OrderID"`
	OrderItem                              *int     `json:"OrderItem"`
	Contract                               *int     `json:"Contract"`
	ContractItem                           *int     `json:"ContractItem"`
	Project                                *int     `json:"Project"`
	WBSElement                             *int     `json:"WBSElement"`
	WBSElementDescription                  *string  `json:"WBSElementDescription"`
	ProductionOrder         			   *int     `json:"ProductionOrder"`
	ProductionOrderItem     			   *int     `json:"ProductionOrderItem"`
	PlannedGoodsIssueDate                  string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                string   `json:"PlannedGoodsReceiptTime"`
	HeaderGrossWeight                      *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight                        *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit                       *string  `json:"HeaderWeightUnit"`
	Incoterms                              *string  `json:"Incoterms"`
	IncotermsText                          *string  `json:"IncotermsText"`
	Items				  				   []Items `json:"Items"`
}

type Items struct {
	DeliveryDocument                              int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                          int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory                  string   `json:"DeliveryDocumentItemCategory"`
	Product                                       string   `json:"Product"`
	ProductSpecification                          *string  `json:"ProductSpecification"`
	SizeOrDimensionText                           *string  `json:"SizeOrDimensionText"`
	DeliveryDocumentItemText                      *string  `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer               *string  `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller              *string  `json:"DeliveryDocumentItemTextBySeller"`
	PlannedGoodsIssueQuantity                     float32  `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit                float32  `json:"PlannedGoodsIssueQtyInBaseUnit"`
	BaseUnit                                      string   `json:"BaseUnit"`
	DeliveryUnit                                  string   `json:"DeliveryUnit"`
	PlannedGoodsIssueDate                         string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                         string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                       string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                       string   `json:"PlannedGoodsReceiptTime"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	OrderID                                		  *int     `json:"OrderID"`
	OrderItem                              		  *int     `json:"OrderItem"`
}
