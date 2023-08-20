package entity_test

import (
	"bangmale/domain/entity"
	"bangmale/domain/value_object"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLecture(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewLecture(entity.DTOLecture{
			NIK:         "123456",
			Name:        "maleakhi",
			PhoneNumber: "085973768236",
			JoinDate:    time.Time{},
			Status:      &value_object.Status{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
	t.Run("Negative Case", func(t *testing.T) {
		newVariable, err := entity.NewLecture(entity.DTOLecture{
			NIK:         "",
			Name:        "",
			PhoneNumber: "085973768236",
			JoinDate:    time.Time{},
			Status:      &value_object.Status{},
		})
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})
}
