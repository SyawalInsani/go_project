package entity

import (
	"bangmale/domain/value_object"
	"errors"
)

type Bachelor struct {
	code   string
	name   string
	status *value_object.Status
}

type DTOBachelor struct {
	Code   string
	Name   string
	Status *value_object.Status
}

func NewBachelor(dto DTOBachelor) (*Bachelor, error) {
	bachelor := &Bachelor{
		code:   dto.Code,
		name:   dto.Name,
		status: dto.Status,
	}

	if bachelor.code == "" {
		return nil, errors.New("CODE MUST BE STATED")
	}

	if bachelor.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	return bachelor, nil
}

func (bcl *Bachelor) GetCode() string {
	return bcl.code
}

func (bcl *Bachelor) GetName() string {
	return bcl.name
}

func (bcl *Bachelor) GetStatus() *value_object.Status {
	return bcl.status
}

func (bcl *Bachelor) SetStatus(status bool) (*Bachelor, error) {
	datastatus, err := value_object.NewStatus(status)
	if err != nil {
		return bcl, err
	}
	bcl.status = datastatus

	return bcl, nil
}
