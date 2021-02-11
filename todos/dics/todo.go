package modules

import (
	todos "github.com/crowdeco/skeleton/todos"
	models "github.com/crowdeco/skeleton/todos/models"
	validations "github.com/crowdeco/skeleton/todos/validations"
	"github.com/sarulabs/dingo/v4"
)

var Todo = []dingo.Def{
	{
		Name:  "module:todo:model",
		Build: (*models.Todo)(nil),
	},
	{
		Name:  "module:todo:validation",
		Build: (*validations.Todo)(nil),
	},
	{
		Name:  "module:todo",
		Build: (*todos.Module)(nil),
		Params: dingo.Params{
			"Context":       dingo.Service("bima:context:background"),
			"Elasticsearch": dingo.Service("bima:connection:elasticsearch"),
			"Handler":       dingo.Service("bima:handler:handler"),
			"Logger":        dingo.Service("bima:handler:logger"),
			"Messenger":     dingo.Service("bima:handler:messager"),
			"Validator":     dingo.Service("module:todo:validation"),
			"Cache":         dingo.Service("bima:cache:memory"),
			"Paginator":     dingo.Service("bima:pagination:paginator"),
		},
	},
	{
		Name:  "module:todo:server",
		Build: (*todos.Server)(nil),
		Params: dingo.Params{
			"Env":      dingo.Service("bima:config:env"),
			"Module":   dingo.Service("module:todo"),
			"Database": dingo.Service("bima:connection:database"),
		},
	},
}
