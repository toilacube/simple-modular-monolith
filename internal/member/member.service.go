package member

import (
	"errors"
	"tutorial/pkg/utils"

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

func (s *Service) Login(dto LoginDTO) (string, error) {
	member, err := s.repository.GetMemberByUsername(dto.Username)

	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(dto.Password, member.Password) {
		return "", errors.New("invalid credentials")
	}

	// generate jwt token
	token, err := utils.GenerateJWTToken(member.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
