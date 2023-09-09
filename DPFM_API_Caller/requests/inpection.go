package requests

type Inspection struct {
	InspectionLot		int     `json:"InspectionLot"`
	Inspection			int     `json:"Inspection"`
	IsCancelled			*bool   `json:"IsCancelled"`
}
