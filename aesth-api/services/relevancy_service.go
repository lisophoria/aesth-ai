package services

import (
	"aesth-api/dto"
	"aesth-api/repositories"
)

type RelevancyService struct {
	relevancyRepo *repositories.RelevancyRepository
}

func NewRelevancyService(relevancyRepo *repositories.RelevancyRepository) *RelevancyService {
	return &RelevancyService{relevancyRepo}
}

func (s *RelevancyService) GetRelevancy(userID uint) (*dto.RelevancyDTO, error) {
	relevancy, err := s.relevancyRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.RelevancyDTO{ Smiling: relevancy.Smiling}, nil
}

func (s *RelevancyService) AdjustRelevancy(userID uint, input dto.RelevancyDTO) (*dto.RelevancyDTO, error) {
	relevancy, err := s.relevancyRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Average with existing values
	relevancy.Smiling = (relevancy.Smiling + input.Smiling) / 2.0

	if err := s.relevancyRepo.Update(relevancy); err != nil {
		return nil, err
	}

	return &dto.RelevancyDTO{ Smiling: relevancy.Smiling }, nil
}

func (s *RelevancyService) GetRelevancyPair(userID uint, diff dto.RelevancyDTO) (*dto.RelevancyPairDTO, error) {
	relevancy, err := s.relevancyRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	Left := &dto.RelevancyDTO{ Smiling: relevancy.Smiling - diff.Smiling }
	Right := &dto.RelevancyDTO{ Smiling: relevancy.Smiling + diff.Smiling }

	return &dto.RelevancyPairDTO{
		Left: *Left,
		Right: *Right,
	}, nil
}