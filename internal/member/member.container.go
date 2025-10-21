package member

import "gorm.io/gorm"

type MemberContainer struct {
	Handler    *MemberHandler
	Service    *Service
	Repository *MemberRepository
}

func NewMemberContainer(db *gorm.DB) *MemberContainer {
	repository := NewMemberRepository(db)
	service := NewMemberService(repository)
	handler := NewMemberHandler(service)

	return &MemberContainer{
		Handler:    handler,
		Service:    service,
		Repository: repository,
	}
}
