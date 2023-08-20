package entity_test

import (
	"bangmale/domain/entity"
	"bangmale/domain/value_object"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewScore(entity.DTOScore{
			Value:     &value_object.Grade{},
			Mahasiswa: &entity.Mahasiswa{},
			Course:    &entity.Course{},
			Lecture:   &entity.Lecture{},
			Date:      time.Time{},
			Submit:    time.Time{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
}
