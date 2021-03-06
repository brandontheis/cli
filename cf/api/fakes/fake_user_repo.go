package fakes

import (
	"github.com/cloudfoundry/cli/cf/errors"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeUserRepository struct {
	StubbedError func(string, string) error

	FindByUsernameUsername   string
	FindByUsernameUserFields models.UserFields
	FindByUsernameNotFound   bool

	ListUsersOrganizationGuid string
	ListUsersSpaceGuid        string
	ListUsersByRole           map[string][]models.UserFields

	CreateUserUsername         string
	CreateUserPassword         string
	CreateUserExists           bool
	CreateUserReturnsHttpError bool

	DeleteUserGuid string

	SetOrgRoleUserGuid         string
	SetOrgRoleOrganizationGuid string
	SetOrgRoleRole             string
	SetOrgRoleCalled           bool
	SetOrgRoleUsername         string
	SetOrgRoleByUsernameError  error

	UnsetOrgRoleUserGuid         string
	UnsetOrgRoleOrganizationGuid string
	UnsetOrgRoleRole             string

	SetSpaceRoleUserGuid         string
	SetSpaceRoleOrgGuid          string
	SetSpaceRoleSpaceGuid        string
	SetSpaceRoleRole             string
	SetSpaceRoleCalled           bool
	SetSpaceRoleByUsernameCalled bool
	SetSpaceRoleByUsernameError  error

	UnsetSpaceRoleUserGuid  string
	UnsetSpaceRoleSpaceGuid string
	UnsetSpaceRoleRole      string

	ListUsersInOrgForRoleWithNoUAA_CallCount   int
	ListUsersInOrgForRole_CallCount            int
	ListUsersInSpaceForRoleWithNoUAA_CallCount int
	ListUsersInSpaceForRole_CallCount          int
}

func (repo *FakeUserRepository) FindByUsername(username string) (user models.UserFields, apiErr error) {
	repo.FindByUsernameUsername = username
	user = repo.FindByUsernameUserFields

	if repo.FindByUsernameNotFound {
		apiErr = errors.NewModelNotFoundError("User", "")
	}

	return
}

func (repo *FakeUserRepository) ListUsersInOrgForRoleWithNoUAA(orgGuid string, roleName string) ([]models.UserFields, error) {
	repo.ListUsersOrganizationGuid = orgGuid
	repo.ListUsersInOrgForRoleWithNoUAA_CallCount++
	return repo.ListUsersByRole[roleName], repo.invokeStubbedError(orgGuid, roleName)
}

func (repo *FakeUserRepository) ListUsersInOrgForRole(orgGuid string, roleName string) ([]models.UserFields, error) {
	repo.ListUsersOrganizationGuid = orgGuid
	repo.ListUsersInOrgForRole_CallCount++
	return repo.ListUsersByRole[roleName], repo.invokeStubbedError(orgGuid, roleName)
}

func (repo *FakeUserRepository) ListUsersInSpaceForRole(spaceGuid string, roleName string) ([]models.UserFields, error) {
	repo.ListUsersSpaceGuid = spaceGuid
	repo.ListUsersInSpaceForRole_CallCount++
	return repo.ListUsersByRole[roleName], repo.invokeStubbedError(spaceGuid, roleName)
}

func (repo *FakeUserRepository) ListUsersInSpaceForRoleWithNoUAA(spaceGuid string, roleName string) ([]models.UserFields, error) {
	repo.ListUsersSpaceGuid = spaceGuid
	repo.ListUsersInSpaceForRoleWithNoUAA_CallCount++
	return repo.ListUsersByRole[roleName], repo.invokeStubbedError(spaceGuid, roleName)
}

func (repo *FakeUserRepository) Create(username, password string) (apiErr error) {
	repo.CreateUserUsername = username
	repo.CreateUserPassword = password

	if repo.CreateUserReturnsHttpError {
		apiErr = errors.NewHttpError(403, "403", "Forbidden")
	}
	if repo.CreateUserExists {
		apiErr = errors.NewModelAlreadyExistsError("User", username)
	}

	return
}

func (repo *FakeUserRepository) Delete(userGuid string) (apiErr error) {
	repo.DeleteUserGuid = userGuid
	return
}

func (repo *FakeUserRepository) SetOrgRole(userGuid, orgGuid, role string) (apiErr error) {
	repo.SetOrgRoleUserGuid = userGuid
	repo.SetOrgRoleOrganizationGuid = orgGuid
	repo.SetOrgRoleRole = role
	repo.SetOrgRoleCalled = true
	return
}

func (repo *FakeUserRepository) SetOrgRoleByUsername(username, orgGuid, role string) (apiErr error) {
	if repo.SetOrgRoleByUsernameError != nil {
		return repo.SetOrgRoleByUsernameError
	} else {
		repo.SetOrgRoleUsername = username
		return
	}
}

func (repo *FakeUserRepository) UnsetOrgRoleByGuid(userGuid, orgGuid, role string) (apiErr error) {
	repo.UnsetOrgRoleUserGuid = userGuid
	repo.UnsetOrgRoleOrganizationGuid = orgGuid
	repo.UnsetOrgRoleRole = role
	return
}

func (repo *FakeUserRepository) SetSpaceRoleByGuid(userGuid, spaceGuid, orgGuid, role string) (apiErr error) {
	repo.SetSpaceRoleUserGuid = userGuid
	repo.SetSpaceRoleOrgGuid = orgGuid
	repo.SetSpaceRoleSpaceGuid = spaceGuid
	repo.SetSpaceRoleRole = role
	repo.SetSpaceRoleCalled = true
	return
}

func (repo *FakeUserRepository) SetSpaceRoleByUsername(username, spaceGuid, orgGuid, role string) (apiErr error) {
	repo.SetSpaceRoleByUsernameCalled = true
	return repo.SetSpaceRoleByUsernameError
}

func (repo *FakeUserRepository) UnsetSpaceRole(userGuid, spaceGuid, role string) (apiErr error) {
	repo.UnsetSpaceRoleUserGuid = userGuid
	repo.UnsetSpaceRoleSpaceGuid = spaceGuid
	repo.UnsetSpaceRoleRole = role
	return
}

func (repo *FakeUserRepository) invokeStubbedError(guid string, role string) error {
	if repo.StubbedError != nil {
		return repo.StubbedError(guid, role)
	}
	return nil
}
