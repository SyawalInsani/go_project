package entity

import (
	"bangmale/domain/value_object"
	"errors"
	"time"
)

type Mahasiswa struct {
	nim         string
	name        string
	phonenumber string
	address     string
	bachelor    *Bachelor
	faculty     *Faculty
	gender      *value_object.Gender
	joinDate    time.Time
}

type DTOMahasiswa struct {
	NIM         string
	Name        string
	PhoneNumber string
	Address     string
	Bachelor    *Bachelor
	Faculty     *Faculty
	Gender      *value_object.Gender
	JoinDate    time.Time
}

func NewMahasiswa(dto DTOMahasiswa) (*Mahasiswa, error) {
	mahasiwa := &Mahasiswa{
		nim:         dto.NIM,
		name:        dto.Name,
		phonenumber: dto.PhoneNumber,
		address:     dto.Address,
		gender:      dto.Gender,
		faculty:     dto.Faculty,
		bachelor:    dto.Bachelor,
		joinDate:    dto.JoinDate,
	}

	if mahasiwa.nim == "" {
		return nil, errors.New("NIM MUST BE STATED")
	}

	if mahasiwa.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	return mahasiwa, nil
}

func (mhs *Mahasiswa) GetNIM() string {
	return mhs.nim
}

func (mhs *Mahasiswa) GetName() string {
	return mhs.name
}

func (mhs *Mahasiswa) GetPhonenumber() string {
	return mhs.phonenumber
}

func (mhs *Mahasiswa) GetAddress() string {
	return mhs.address
}

func (mhs *Mahasiswa) GetGender() *value_object.Gender {
	return mhs.gender
}

func (mhs *Mahasiswa) SetGender(gender bool) (*Mahasiswa, error) {
	datagender, err := value_object.NewGender(gender)
	if err != nil {
		return mhs, err
	}
	mhs.gender = datagender

	return mhs, nil
}

func (mhs *Mahasiswa) GetFaculty() *Faculty {
	return mhs.faculty
}

func (mhs *Mahasiswa) GetTypeBachelor() *Bachelor {
	return mhs.bachelor
}

func (mhs *Mahasiswa) SetTypeBachelor(bachelor *Bachelor) {
	mhs.bachelor = bachelor
}

func (mhs *Mahasiswa) GetJoinDate() time.Time {
	return mhs.joinDate
}

func (mhs *Mahasiswa) SetJoinDate(joinDate time.Time) {
	mhs.joinDate = joinDate
}
