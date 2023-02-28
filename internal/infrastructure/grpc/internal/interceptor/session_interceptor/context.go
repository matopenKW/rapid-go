package session_interceptor

import (
	"context"

	"github.com/abyssparanoia/rapid-go/internal/domain/model"
	"github.com/abyssparanoia/rapid-go/internal/pkg/errors"
	"github.com/abyssparanoia/rapid-go/internal/pkg/nullable"
	"github.com/volatiletech/null/v8"
)

type SessionContext struct {
	AuthUID   string
	TenantID  null.String
	StaffID   null.String
	StaffRole nullable.Type[model.StaffRole]
}

type contextKey struct{}

func newSessionContext(claims *model.Claims) *SessionContext {
	return &SessionContext{
		AuthUID:   claims.AuthUID,
		TenantID:  claims.TenantID,
		StaffID:   claims.StaffID,
		StaffRole: claims.StaffRole,
	}
}

var (
	sessionContextKey contextKey = contextKey{}
)

func SaveSessionContext(
	ctx context.Context,
	sessionContext *SessionContext,
) context.Context {
	return context.WithValue(ctx, sessionContextKey, *sessionContext)
}

func RequireSessionContext(ctx context.Context) (*SessionContext, error) {
	sctx, ok := GetSessionContext(ctx)
	if !ok {
		return nil, errors.UnauthorizedErr.New()
	}
	return sctx, nil
}

func GetSessionContext(ctx context.Context) (*SessionContext, bool) {
	sessionContext, ok := ctx.Value(sessionContextKey).(SessionContext)
	return &sessionContext, ok
}
