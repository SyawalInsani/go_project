package value_object

type Status struct {
	value       bool
	detailValue string
}

func NewStatus(datastatus bool) (*Status, error) {
	status := &Status{value: datastatus}

	if status.value == false {
		status.detailValue = "InActive"
	} else {
		status.detailValue = "Active"
	}

	return status, nil
}

func (sts *Status) GetValue() bool {
	return sts.value
}

func (sts *Status) GetDetailValue() string {
	return sts.detailValue
}
