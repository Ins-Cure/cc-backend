package service

import (
	"mods/dto"
	"mods/entity"
	"mods/repository"
	"mods/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type predictionService struct {
	predictionRepository repository.PredictionRepository
}

type PredictionService interface {
	CreatePrediction(ctx *gin.Context, predictionDTO dto.PredictImageDTO, userID string) (entity.Prediction, error)
}

func NewPredictionService(pr repository.PredictionRepository) PredictionService {
	return &predictionService{
		predictionRepository: pr,
	}
}

func (ps *predictionService) CreatePrediction(ctx *gin.Context, predictionDTO dto.PredictImageDTO, userID string) (entity.Prediction, error) {
	id := uuid.NewString()
	imageFile := predictionDTO.File

	img_uuid, err := utils.UploadToBucket(imageFile)
	if err != nil {
		return entity.Prediction{}, err
	}

	result, err := utils.PredictionAPI(img_uuid)
	if err != nil {
		return entity.Prediction{}, err
	}

	newPrediction := entity.Prediction{
		Pr_ID:          id,
		Gambar:         img_uuid,
		Hasil_Prediksi: result,
		Tgl:            time.Now(),
		UserID:         userID,
	}

	return ps.predictionRepository.AddPrediction(ctx, newPrediction)
}
