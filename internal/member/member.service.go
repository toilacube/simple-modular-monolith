package member

import (
	"errors"

	"gorm.io/gorm"
)

type Service struct {
	repository *MemberRepository
}

func NewMemberService(repository *MemberRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Register(dto RegisterDTO) error {
	_, err := s.repository.GetMemberByUsername(dto.Username)

	if err == nil {
		return errors.New("username already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return s.repository.CreateMember(dto)
}
