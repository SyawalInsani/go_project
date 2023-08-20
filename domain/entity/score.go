package entity

import (
	"bangmale/domain/value_object"
	"time"
)

type Score struct {
	value     *value_object.Grade
	mahasiswa *Mahasiswa
	course    *Course
	lecture   *Lecture
	date      time.Time
	submit    time.Time
}

type DTOScore struct {
	Value     *value_object.Grade
	Mahasiswa *Mahasiswa
	Course    *Course
	Lecture   *Lecture
	Date      time.Time
	Submit    time.Time
}

func NewScore(dto DTOScore) (*Score, error) {
	score := &Score{
		value:     dto.Value,
		mahasiswa: dto.Mahasiswa,
		course:    dto.Course,
		lecture:   dto.Lecture,
		date:      dto.Date,
		submit:    dto.Submit,
	}

	return score, nil
}

func (scr *Score) GetValue() *value_object.Grade {
	return scr.value
}

func (scr *Score) SetValue(grade string) (*Score, error) {
	datagrade, err := value_object.NewGrade(grade)
	if err != nil {
		return scr, err
	}
	scr.value = datagrade
	return scr, nil
}

func (scr *Score) GetMahasiswa() *Mahasiswa {
	return scr.mahasiswa
}

func (scr *Score) SetMahasiswa(mahasiswa *Mahasiswa) {
	scr.mahasiswa = mahasiswa
}

func (scr *Score) GetCourse() *Course {
	return scr.course
}

func (scr *Score) SetCourse(course *Course) {
	scr.course = course
}

func (scr *Score) GetLecture() *Lecture {
	return scr.lecture
}

func (scr *Score) SetLecture(lecture *Lecture) {
	scr.lecture = lecture
}

func (scr *Score) GetDate() time.Time {
	return scr.date
}

func (scr *Score) GetSubmit() time.Time {
	return scr.submit
}
