package entity_test

import (
	"bangmale/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemesterTypes(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewSemeterType(&entity.DTOSemesterType{
			Name:         "a",
			TypeSemester: "a",
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
	t.Run("Negative Case", func(t *testing.T) {
		newVariable, err := entity.NewSemeterType(&entity.DTOSemesterType{
			Name:         "",
			TypeSemester: "a",
		})
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})
}
