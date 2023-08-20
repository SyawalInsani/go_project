package entity_test

import (
	"bangmale/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSemesterDetail(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewSemesterDetail(entity.DTOSemesterDetail{
			Mahasiswa:    &entity.Mahasiswa{},
			SemesterType: &entity.SemesterType{},
			Course:       []*entity.Course{},
			StartDate:    time.Time{},
			EndDate:      time.Time{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
}
