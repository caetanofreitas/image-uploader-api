package repository

import (
	"os"
	"testing"
	"uploader/adapter/repository/fixture"
	"uploader/entity"
	"uploader/utils/formats"
	"uploader/utils/upload_status"

	"github.com/stretchr/testify/assert"
)

func TestImageRepositoryDb_Insert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewImageRepository(db)
	err := repository.Insert("1", "Image Name", 1.00, formats.JPG, upload_status.PENDING, "")

	assert.Nil(t, err)
}

func TestImageRepositoryDb_Update(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewImageRepository(db)
	err := repository.Update("1", "Image Name", 1.00, formats.JPG, upload_status.FAILED, "error")

	assert.Nil(t, err)
}

func TestImageRepositoryDb_Get(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	image := entity.NewImage()
	image.ID = "1"
	image.Name = "Image Name"
	image.Size = 1.00
	image.Extension = formats.JPG
	image.Status = upload_status.PENDING
	repository := NewImageRepository(db)
	repository.Insert(image.ID, image.Name, image.Size, image.Extension, upload_status.PENDING, "")

	result, err := repository.Get()
	var expectedOutput []entity.Image
	expectedOutput = append(expectedOutput, *image)

	assert.Nil(t, err)
	assert.Equal(t, result, expectedOutput)
	assert.Equal(t, len(result), 1)
}

func TestImageRepositoryDb_GetDetail(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	expectedOutput := entity.NewImage()
	expectedOutput.ID = "1"
	expectedOutput.Name = "Image Name"
	expectedOutput.Size = 1.00
	expectedOutput.Extension = formats.JPG
	expectedOutput.Status = upload_status.PENDING
	repository := NewImageRepository(db)
	repository.Insert(expectedOutput.ID, expectedOutput.Name, expectedOutput.Size, expectedOutput.Extension, upload_status.PENDING, "")

	result, err := repository.GetDetail("1")

	assert.Nil(t, err)
	assert.Equal(t, result, *expectedOutput)
}
