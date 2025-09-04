package frontweb

import (
	"standard-struct-golang/app"

	auth_repo "standard-struct-golang/modules/frontweb/modules/auth/repositories"
	auth_services "standard-struct-golang/modules/frontweb/modules/auth/services"
	example_repo "standard-struct-golang/modules/frontweb/modules/example/repositories"
	example_services "standard-struct-golang/modules/frontweb/modules/example/services"

	auth_handler "standard-struct-golang/modules/frontweb/modules/auth/handler"
	example_handler "standard-struct-golang/modules/frontweb/modules/example/handler"
	repositories "standard-struct-golang/modules/frontweb/repo"
)

func Create(app *app.App) error {
	repo, err := repositories.NewRepo(app)
	if err != nil {
		return err
	}

	// Repository
	exampleRepo := example_repo.NewExampleRepo(repo)
	authRepo := auth_repo.NewAuthRepo(repo)

	// Service
	exampleSvc := example_services.NewExampleService(exampleRepo)
	authSvc := auth_services.NewAuthService(authRepo)

	// Handler
	exampleHdlr := example_handler.NewExampleHandler(exampleSvc)
	authHdlr := auth_handler.NewAuthHandler(authSvc)

	// Route
	prefix := app.Router.Group(app.Config.AppConfig.PrefixPath)

	// Register
	exampleHdlr.AddExampleRouter(prefix)
	authHdlr.AddAuthRouter(prefix)

	return nil
}
