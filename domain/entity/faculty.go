package entity

import (
	"bangmale/domain/value_object"
	"errors"
)

type Faculty struct {
	name   string
	code   string
	status *value_object.Status
	course []*Course
}

type DTOFaculty struct {
	Name   string
	Code   string
	Status *value_object.Status
	Course []*Course
}

func NewFaculty(dto DTOFaculty) (*Faculty, error) {
	faculty := &Faculty{
		name:   dto.Name,
		code:   dto.Code,
		status: dto.Status,
		course: dto.Course,
	}

	if faculty.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	if faculty.code == "" {
		return nil, errors.New("CODE MUST BE STATED")
	}

	return faculty, nil
}

func (fcl *Faculty) GetName() string {
	return fcl.name
}

func (fcl *Faculty) GetCode() string {
	return fcl.code
}

func (fcl *Faculty) GetStatus() *value_object.Status {
	return fcl.status
}

func (fcl *Faculty) SetStatus(status bool) (*Faculty, error) {
	datastatus, err := value_object.NewStatus(status)
	if err != nil {
		return fcl, err
	}
	fcl.status = datastatus

	return fcl, nil
}

func (fcl *Faculty) GetCourses() []*Course {
	return fcl.course
}
