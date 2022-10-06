package repository

import "context"

//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_repository
type Transactable interface {
	RWTx(ctx context.Context, fn func(ctx context.Context) error) error
}