package entity_test

import (
	"bangmale/domain/entity"
	"bangmale/domain/value_object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBachelor(t *testing.T) {
	t.Run("Positive case", func(t *testing.T) {
		newVariable, err := entity.NewBachelor(entity.DTOBachelor{
			Code:   "1",
			Name:   "Bachelor",
			Status: &value_object.Status{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
	t.Run("Negative case", func(t *testing.T) {
		newVariable, err := entity.NewBachelor(entity.DTOBachelor{
			Code:   "",
			Name:   "",
			Status: &value_object.Status{},
		})
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})
}
