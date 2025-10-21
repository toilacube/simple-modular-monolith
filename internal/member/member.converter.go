package member

import (
	"tutorial/pkg/model"
	"tutorial/pkg/utils"
)

func ConvertRegisterDTOToMember(dto RegisterDTO) (model.Member, error) {

	hashed_pwd, err := utils.HashPassword(dto.Password)
	if err != nil {
		return model.Member{}, err
	}
	id := utils.GetUUID()
	return model.Member{
		ID:       id,
		Username: dto.Username,
		Password: hashed_pwd,
	}, nil
}

func ConvertMemberToDTO(member model.Member) MemberDTO {
	return MemberDTO{
		ID:        member.ID,
		Username:  member.Username,
		CreatedAt: member.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
