package entity_test

import (
	"bangmale/domain/entity"
	"bangmale/domain/value_object"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMahasiswa(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewMahasiswa(entity.DTOMahasiswa{
			NIM:         "UKRI0012",
			Name:        "Syawal",
			PhoneNumber: "008985748345",
			Address:     "Jl Kebon Jeruk",
			Bachelor:    &entity.Bachelor{},
			Faculty:     &entity.Faculty{},
			Gender:      &value_object.Gender{},
			JoinDate:    time.Time{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
	t.Run("Negative Case", func(t *testing.T) {
		newVariable, err := entity.NewMahasiswa(entity.DTOMahasiswa{
			NIM:         "",
			Name:        "",
			PhoneNumber: "008985748345",
			Address:     "Jl Kebon Jeruk",
			Bachelor:    &entity.Bachelor{},
			Faculty:     &entity.Faculty{},
			Gender:      &value_object.Gender{},
			JoinDate:    time.Time{},
		})
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})
}

// func TestMahasiswaGet(t *testing.T) {
// 	newVariable, err := entity.NewMahasiswa(entity.DTOMahasiswa{
// 		NIM:         "",
// 		Name:        "",
// 		PhoneNumber: "",
// 		Address:     "",
// 		Bachelor:    &entity.Bachelor{},
// 		Faculty:     &entity.Faculty{},
// 		Gender:      &value_object.Gender{},
// 		JoinDate:    time.Time{},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, "UKRI0012", newVariable.GetNIM())
// }
