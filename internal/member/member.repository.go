package member

import (
	"tutorial/pkg/model"

	"gorm.io/gorm"
)

type MemberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

func (r *MemberRepository) CreateMember(dto RegisterDTO) error {
	member, err := ConvertRegisterDTOToMember(dto)
	if err != nil {
		return err
	}
	return r.db.Create(&member).Error
}

func (r *MemberRepository) GetMemberByUsername(username string) (model.Member, error) {
	var member model.Member
	err := r.db.Where("username = ?", username).First(&member).Error
	if err != nil {
		return model.Member{}, err
	}
	return member, nil
}
