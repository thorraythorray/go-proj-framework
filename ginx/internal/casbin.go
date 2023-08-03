package internal

import (
	"github.com/thorraythorray/go-proj/pkg/rbac"
)

var DefaultCasbinRules = []rbac.CasbinRules{
	{
		Role: "admin",
		CasbinInfos: []rbac.CasbinSubRule{
			{Path: "/login", Method: "GET"},
			{Path: "/register", Method: "GET"},
		},
	},
}
