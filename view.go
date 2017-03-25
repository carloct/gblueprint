package main

import (
	"github.com/CloudyKit/jet"
)

// Template ...
type Template struct {
	Engine *jet.Set
	Env    bool
}

// Configure ...
func (v *Template) Configure(env bool) {
	v.Engine = jet.NewHTMLSet("./views")
	v.Engine.SetDevelopmentMode(env)
}

// Render ...
/*func (v *Template) Render(w http.ResponseWriter, template string, vars jet.VarMap, ctx interface{}) {
	t, err := v.Engine.GetTemplate(template)
	if err != nil {
		//exitWithError("cannot find the page", err)
	}

	if err = t.Execute(w, vars, ctx); err != nil {

	}

}*/
