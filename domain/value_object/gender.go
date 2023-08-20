package value_object

type Gender struct {
	values bool
	detail string
}

func NewGender(datagender bool) (*Gender, error) {
	gender := &Gender{values: datagender}

	if gender.values == true {
		gender.detail = "Man"
	} else {
		gender.detail = "Woman"
	}
	return gender, nil
}

func (gdr *Gender) Getvalues() bool {
	return gdr.values
}

func (gdr *Gender) GetDetail() string {
	return gdr.detail
}
