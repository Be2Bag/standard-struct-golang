package module

import (
	"standard-struct-golang/app"
	"standard-struct-golang/modules/frontweb"
)

func CreateModule(app *app.App) error {
	l := app.NewLogger().WithField("module", "generic")

	//Create Module
	if err := frontweb.Create(app); err != nil {
		l.Errorln("[x] Create FrontWeb module error -:", err)
		return err
	}
	return nil
}
