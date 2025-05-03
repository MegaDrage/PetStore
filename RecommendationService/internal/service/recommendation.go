package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"MegaDrage/PetStore/RecommendationService/internal/repository"
)

type RecommendationService struct {
	mongoRepo     *repository.MongoRepository
	kafkaProducer *repository.KafkaProducer
	httpClient    *http.Client
	profileURL    string
	labURL        string
}

func NewRecommendationService(mongoRepo *repository.MongoRepository, kafkaProducer *repository.KafkaProducer, cfg config.Config) *RecommendationService {
	return &RecommendationService{
		mongoRepo:     mongoRepo,
		kafkaProducer: kafkaProducer,
		httpClient:    &http.Client{Timeout: 10 * time.Second},
		profileURL:    cfg.ProfileURL,
		labURL:        cfg.LabURL,
	}
}

func (s *RecommendationService) getPetProfile(ctx context.Context, petID string) (models.PetProfile, error) {
	var profile models.PetProfile
	url := fmt.Sprintf("%s/profiles/%s", s.profileURL, petID)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return profile, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return profile, err
	}
	return profile, nil
}

func (s *RecommendationService) getLatestLabResults(ctx context.Context, petID string) (models.LabResult, error) {
	var result models.LabResult
	url := fmt.Sprintf("%s/results/%s/latest", s.labURL, petID)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

func (s *RecommendationService) generateFeedRecommendation(profile models.PetProfile) models.Recommendation {
	content := "Recommended hypoallergenic feed"
	if len(profile.Allergies) > 0 {
		content = fmt.Sprintf("Recommended feed avoiding allergens: %v", profile.Allergies)
	}
	return models.Recommendation{
		PetID:     profile.PetID,
		Type:      "feed",
		Content:   content,
		Priority:  2,
		CreatedAt: time.Now(),
	}
}

func (s *RecommendationService) generateVaccinationReminder(profile models.PetProfile) models.Recommendation {
	schedule := time.Now().AddDate(0, 6, 0)
	return models.Recommendation{
		PetID:       profile.PetID,
		Type:        "vaccination",
		Content:     fmt.Sprintf("Schedule %s's next vaccination", profile.Species),
		Priority:    1,
		CreatedAt:   time.Now(),
		ScheduledAt: schedule,
	}
}

func (s *RecommendationService) generateDewormingReminder(profile models.PetProfile) models.Recommendation {
	schedule := time.Now().AddDate(0, 3, 0)
	return models.Recommendation{
		PetID:       profile.PetID,
		Type:        "deworming",
		Content:     fmt.Sprintf("Schedule %s's next deworming", profile.Species),
		Priority:    1,
		CreatedAt:   time.Now(),
		ScheduledAt: schedule,
	}
}

func (s *RecommendationService) GenerateRecommendations(ctx context.Context, petID string) ([]models.Recommendation, error) {
	profile, err := s.getPetProfile(ctx, petID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch profile: %v", err)
	}

	recommendations := []models.Recommendation{
		s.generateFeedRecommendation(profile),
		s.generateVaccinationReminder(profile),
		s.generateDewormingReminder(profile),
	}

	for _, rec := range recommendations {
		if err := s.mongoRepo.SaveRecommendation(ctx, rec); err != nil {
			// Log error but continue processing
			fmt.Printf("Failed to save recommendation: %v\n", err)
			continue
		}
		if err := s.kafkaProducer.PublishNotification(rec); err != nil {
			fmt.Printf("Failed to publish notification: %v\n", err)
		}
	}

	return recommendations, nil
}