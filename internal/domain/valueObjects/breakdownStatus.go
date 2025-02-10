package valueobjects

type BreakdownStatus int

const (
	BreakdownStatusUnknown = iota
	BreakdownStatusProcessing
	BreakdownStatusError
	BreakdownStatusSuccess
)

var BreakdownStatusNames = map[BreakdownStatus]string{
	BreakdownStatusProcessing: "Processing",
	BreakdownStatusError:      "Error",
	BreakdownStatusSuccess:    "Success",
}
