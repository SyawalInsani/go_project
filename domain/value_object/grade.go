package value_object

import "errors"

type Grade struct {
	grade string
}

func NewGrade(datagrade string) (*Grade, error) {
	grade := &Grade{grade: datagrade}

	if grade.grade == "" {
		return nil, errors.New("grade must be stated")
	}
	if grade.grade == "E" || grade.grade == "D" || grade.grade == "C" || grade.grade == "B" || grade.grade == "A-" || grade.grade == "A" {
		return nil, errors.New("Grade not stated very well")
	}
	return grade, nil
}

func (grd *Grade) GetGradeName() string {
	return grd.grade
}
