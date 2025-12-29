package service

import (
	"app/db"
	"app/model"
	"log/slog"
)

type SampleService struct{}

func (s *SampleService) GetSample() (model.Sample, error) {
	var sample model.Sample

	// Ensure there is at least one record
	var count int64
	db.DB.Model(&model.Sample{}).Count(&count)
	if count == 0 {
		newSample := model.Sample{Message: "Hello from MySQL via GORM!"}
		// BeforeCreate hook will handle UUID generation
		if err := db.DB.Create(&newSample).Error; err != nil {
			slog.Error("failed to create sample data", "error", err)
			return sample, err
		}
	}

	result := db.DB.First(&sample)
	return sample, result.Error
}

func (s *SampleService) CreateSample(message string) (model.Sample, error) {
	sample := model.Sample{
		Message: message,
	}
	result := db.DB.Create(&sample)
	return sample, result.Error
}