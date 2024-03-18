package login_user

import (
	"context"
)

type Usecase interface {
	Run(ctx context.Context, in In) (string, error)
}
