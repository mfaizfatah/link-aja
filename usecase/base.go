package usecase

import (
	"context"

	"github.com/mfaizfatah/link-aja/model"
	"github.com/mfaizfatah/link-aja/repository"
)

// uc struct with value interface Repository
type uc struct {
	query repository.Repo
}

// Usecases represent the Usecases contract
type Usecases interface {
	CheckSaldo(ctx context.Context, accNumber string) (context.Context, interface{}, string, int, error)
	Transfer(ctx context.Context, transfer model.Transfer) (context.Context, interface{}, string, int, error)
}

/*NewUC will create an object that represent the Usecases interface (Usecases)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Usecases
 *
 * @return
 * uc struct with value interface Repository
 */
func NewUC(r repository.Repo) Usecases {
	return &uc{query: r}
}
