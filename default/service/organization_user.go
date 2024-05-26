package service

import (
	"strconv"
	"strings"

	"github.com/gophab/gophrame/core/inject"
	"github.com/gophab/gophrame/core/query"
	"github.com/gophab/gophrame/service"

	"github.com/gophab/gophrame/default/domain"
	"github.com/gophab/gophrame/default/repository"
)

type OrganizationUserService struct {
	service.BaseService
	OrganizationUserRepository *repository.OrganizationUserRepository `inject:"organizationUserRepository"`
	OrganizationRepository     *repository.OrganizationRepository     `inject:"organizationRepository"`
}

var organizationUserService *OrganizationUserService = &OrganizationUserService{}

func init() {
	inject.InjectValue("organizationUserService", organizationUserService)
}

func (u *OrganizationUserService) ListMembers(organizationId int64, userName string, pageable query.Pageable) (int64, []domain.OrganizationMember) {
	return u.OrganizationUserRepository.ListMembers(organizationId, userName, pageable)
}

// 根据用户id查询所有可能的岗位节点id
func (u *OrganizationUserService) GetUserOrganizationIds(userId string) []int {
	//获取用户的所有岗位id
	organizationUsers := u.OrganizationUserRepository.GetByUserId(userId)

	memberIds := []int64{}
	for _, v := range organizationUsers {
		memberIds = append(memberIds, v.OrganizationId)
	}

	//根据岗位ID获取所有的岗位ID,父子级(需要去重)
	organization := u.OrganizationRepository.GetByIds(memberIds)
	organizationIdArr := []int{}
	for _, v := range organization {
		idArr := strings.Split(v.PathInfo, ",")
		for _, vv := range idArr {
			id, _ := strconv.Atoi(vv)
			if id > 0 {
				organizationIdArr = append(organizationIdArr, id)
			}
		}
	}
	return organizationIdArr
}
