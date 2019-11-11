//+build wireinject

package server

import (
	"backend_api/application/controller"
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB"
	"backend_api/pkg/server/middleware"
	"github.com/google/wire"
)

func InitializeAuthenticator() *middleware.AuthenticateCtrl {
	wire.Build(middleware.NewAuthenticateCtl, mysqlDB.NewUserRepoImpl, repository.InitDB)
	return &middleware.AuthenticateCtrl{}
}

func InitializeAuthHandler() *controller.AuthCtrl {
	wire.Build(controller.NewAuthCtl, mysqlDB.NewUserRepoImpl, repository.InitDB)
	return &controller.AuthCtrl{}
}

func InitializeUserHandler() *controller.UserCtrl {
	wire.Build(controller.NewUserCtl, mysqlDB.NewUserRepoImpl, repository.InitDB)
	return &controller.UserCtrl{}
}

func InitializeLoginHandler() *controller.LoginCtrl {
	wire.Build(controller.NewLoginCtl, mysqlDB.NewUserRepoImpl, repository.InitDB)
	return &controller.LoginCtrl{}
}

func InitializeProjectHandler() *controller.ProjectCtrl {
	wire.Build(
		controller.NewProjectCtl,
		mysqlDB.NewProjectRepoImpl,
		mysqlDB.NewProjectPositionRepoImpl,
		mysqlDB.NewProjectSkillRepoImpl,
		mysqlDB.NewProjectMentorRepoImpl,
		mysqlDB.NewProjectUserRepoImpl,
		mysqlDB.NewPositionRepoImpl,
		mysqlDB.NewSkillRepoImpl,
		mysqlDB.NewMentorRepoImpl,
		mysqlDB.NewUserRepoImpl,
		repository.InitDB,
	)
	return &controller.ProjectCtrl{}
}
