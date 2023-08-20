package entity_test

import (
	"bangmale/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaculty(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewFaculty(entity.DTOFaculty{
			Name: "IT",
			Code: "1",
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
	t.Run("Negative Case", func(t *testing.T) {
		newVariable, err := entity.NewFaculty(entity.DTOFaculty{
			Name: "",
			Code: "",
		})
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})
}
