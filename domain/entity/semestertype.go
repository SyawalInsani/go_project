package entity

import "errors"

type SemesterType struct {
	name         string
	typesemester string
}

type DTOSemesterType struct {
	Name         string
	TypeSemester string
}

func NewSemeterType(dto *DTOSemesterType) (*SemesterType, error) {
	semestertype := &SemesterType{
		name:         dto.Name,
		typesemester: dto.TypeSemester,
	}

	if semestertype.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	if semestertype.typesemester == "" {
		return nil, errors.New("SEMESTER MUST BE STATED")
	}
	return semestertype, nil
}

func (st *SemesterType) GetName() string {
	return st.name
}

func (st *SemesterType) GetSemesterType() string {
	return st.typesemester
}
