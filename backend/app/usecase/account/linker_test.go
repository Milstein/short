package account

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
	"testing"

	"github.com/byliuyang/app/mdtest"
)

func TestLinker_IsAccountLinked(t *testing.T) {
	testCases := []struct {
		name             string
		keys             []string
		users            map[string]entity.User
		accountRelation  []repo.Relation
		ssoUser          entity.SSOUser
		expectedIsLinked bool
	}{}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			keyGen := keygen.NewFake(testCase.keys)
			userRepo := repo.NewUserFake(testCase.users)
			accountsRelationRepo := repo.NewAccountsRelationFake(testCase.accountRelation)
			linker := NewLinker(&keyGen, userRepo, &accountsRelationRepo)
			isLinked, err := linker.IsAccountLinked(testCase.ssoUser)
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedIsLinked, isLinked)
		})
	}
}

func TestLinker_LinkAccount(t *testing.T) {
	testCases := []struct {
		name            string
		keys            []string
		users           map[string]entity.User
		accountRelation []repo.Relation
		ssoUser         entity.SSOUser
	}{}

	for _, testCase := range testCases {
		keyGen := keygen.NewFake(testCase.keys)
		fakeUserRepo := repo.NewUserFake(testCase.users)
		fakeAccountsRelationRepo := repo.NewAccountsRelationFake(testCase.accountRelation)
		linker := NewLinker(&keyGen, fakeUserRepo, &fakeAccountsRelationRepo)
		err := linker.LinkAccount(testCase.ssoUser)
		mdtest.Equal(t, nil, err)
		user := entity.User{
			ID: testCase.keys[0],
		}
		gotIsExist := fakeAccountsRelationRepo.IsRelationExist(testCase.ssoUser, user)

	}
}
