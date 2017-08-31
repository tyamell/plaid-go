package plaid

type mfaResponse struct {
	Type string

	Device     mfaDevice
	List       []mfaList
	Questions  []mfaQuestion
	Selections []mfaSelection
}

type mfaDevice struct {
	Message string
}
type mfaList struct {
	DeviceID string `json:"device_id"`
	Mask     string
	Type     string
}
type mfaQuestion struct {
	Question string
}
type mfaSelection struct {
	Answers  []string
	Question string
}
