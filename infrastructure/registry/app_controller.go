package registry

import "github.com/rostikts/fintech_test_project/internal/loader"

type AppController struct {
	Transaction interface{ loader.Handler }
}
