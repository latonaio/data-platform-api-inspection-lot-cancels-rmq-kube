package requests

type Header struct {
	InspectionLot		int     `json:"InspectionLot"`
	IsCancelled			*bool   `json:"IsCancelled"`
}
