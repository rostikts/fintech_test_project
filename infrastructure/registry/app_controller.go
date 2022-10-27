package registry

import "github.com/rostikts/fintech_test_project/internal/transaction"

type AppController struct {
	Transaction interface{ transaction.Handler }
}
