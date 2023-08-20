package entity

import (
	"bangmale/domain/value_object"
	"errors"
	"time"
)

type Lecture struct {
	nik         string
	name        string
	phonenumber string
	joinDate    time.Time
	status      *value_object.Status
}

type DTOLecture struct {
	NIK         string
	Name        string
	PhoneNumber string
	JoinDate    time.Time
	Status      *value_object.Status
}

func NewLecture(dto DTOLecture) (*Lecture, error) {
	lecture := &Lecture{
		nik:         dto.NIK,
		name:        dto.Name,
		phonenumber: dto.PhoneNumber,
		joinDate:    dto.JoinDate,
		status:      dto.Status,
	}

	if lecture.nik == "" {
		return nil, errors.New("NIK MUST BE STATED")
	}
	if lecture.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	return lecture, nil
}

func (lct *Lecture) GetNik() string {
	return lct.nik
}

func (lct *Lecture) GetName() string {
	return lct.name
}

func (lct *Lecture) GetPhonenumber() string {
	return lct.phonenumber
}

func (lct *Lecture) GetJoinDate() time.Time {
	return lct.joinDate
}

func (lct *Lecture) SetJoinDate(joinDate time.Time) {
	lct.joinDate = joinDate
}

func (lct *Lecture) GetStatus() *value_object.Status {
	return lct.status
}

func (lct *Lecture) SetStatus(status bool) (*Lecture, error) {
	datastatus, err := value_object.NewStatus(status)
	if err != nil {
		return lct, err
	}
	lct.status = datastatus

	return lct, nil
}
