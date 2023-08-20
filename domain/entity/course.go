package entity

import (
	"bangmale/domain/value_object"
	"errors"
)

type Course struct {
	name     string
	code     string
	lecture  *Lecture
	skscount *value_object.Sks
	status   *value_object.Status
}

type DTOCourse struct {
	Name     string
	Code     string
	Lecture  *Lecture
	SKScount *value_object.Sks
	Status   *value_object.Status
}

func NewCourse(dto DTOCourse) (*Course, error) {
	course := &Course{
		name:     dto.Name,
		code:     dto.Code,
		lecture:  dto.Lecture,
		skscount: dto.SKScount,
		status:   dto.Status,
	}

	if course.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	if course.code == "" {
		return nil, errors.New("CODE MUST BE STATED")
	}

	return course, nil
}

func (crs *Course) GetName() string {
	return crs.name
}

func (crs *Course) GetCode() string {
	return crs.code
}

func (crs *Course) GetLecture() *Lecture {
	return crs.lecture
}

func (crs *Course) GetSKSCount() *value_object.Sks {
	return crs.skscount
}

func (crs *Course) GetStatus() *value_object.Status {
	return crs.status
}

func (crs *Course) SetStatus(status bool) (*Course, error) {
	datastatus, err := value_object.NewStatus(status)
	if err != nil {
		return crs, err
	}
	crs.status = datastatus

	return crs, nil
}
