package entity

import "time"

type SemesterDetail struct {
	mahasiswa    *Mahasiswa
	semesterType *SemesterType
	course       []*Course
	startDate    time.Time
	endDate      time.Time
}

type DTOSemesterDetail struct {
	Mahasiswa    *Mahasiswa
	SemesterType *SemesterType
	Course       []*Course
	StartDate    time.Time
	EndDate      time.Time
}

func NewSemesterDetail(dto DTOSemesterDetail) (*SemesterDetail, error) {
	semesterdetail := &SemesterDetail{
		mahasiswa:    dto.Mahasiswa,
		semesterType: dto.SemesterType,
		course:       dto.Course,
		startDate:    dto.StartDate,
		endDate:      dto.EndDate,
	}
	return semesterdetail, nil
}

func (std *SemesterDetail) GetMahasiswa() *Mahasiswa {
	return std.mahasiswa
}

func (std *SemesterDetail) SetMahasiswa(mahasiswa *Mahasiswa) {
	std.mahasiswa = mahasiswa
}

func (std *SemesterDetail) GetSemesterType() *SemesterType {
	return std.semesterType
}

func (std *SemesterDetail) SetSemesterType(semesterType *SemesterType) {
	std.semesterType = semesterType
}

func (std *SemesterDetail) GetCourse() []*Course {
	return std.course
}

func (std *SemesterDetail) SetCourse(course []*Course) {
	std.course = course
}

func (std *SemesterDetail) GetStartDate() time.Time {
	return std.startDate
}

func (std *SemesterDetail) SetStartDate(startDate time.Time) {
	std.startDate = startDate
}

func (std *SemesterDetail) GetEndDate() time.Time {
	return std.endDate
}

func (std *SemesterDetail) SetEndDate(endDate time.Time) {
	std.endDate = endDate
}
