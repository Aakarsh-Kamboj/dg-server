package domain

type ControlStatus string

const (
	StatusCompliant     ControlStatus = "Compliant"
	StatusNonCompliant  ControlStatus = "Non Compliant"
	StatusNotApplicable ControlStatus = "Not Applicable"
)

func IsValidStatus(s ControlStatus) bool {
	switch s {
	case StatusCompliant, StatusNonCompliant, StatusNotApplicable:
		return true
	default:
		return false
	}
}
