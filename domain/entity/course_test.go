package entity_test

import (
	"bangmale/domain/entity"
	"bangmale/domain/value_object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCourse(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewCourse(entity.DTOCourse{
			Name:     "Kalkulus",
			Code:     "1",
			Lecture:  &entity.Lecture{},
			SKScount: &value_object.Sks{},
			Status:   &value_object.Status{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
	t.Run("Negative Case", func(t *testing.T) {
		newVariable, err := entity.NewCourse(entity.DTOCourse{
			Name:     "",
			Code:     "",
			Lecture:  &entity.Lecture{},
			SKScount: &value_object.Sks{},
			Status:   &value_object.Status{},
		})
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})
}
