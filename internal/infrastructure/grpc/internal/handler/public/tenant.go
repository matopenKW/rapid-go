package public

import (
	"context"

	"github.com/abyssparanoia/rapid-go/internal/infrastructure/grpc/internal/marshaller"
	"github.com/abyssparanoia/rapid-go/internal/usecase/input"
	public_apiv1 "github.com/abyssparanoia/rapid-go/schema/proto/pb/rapid/public_api/v1"
)

func (h *PublicHandler) PublicGetTenant(ctx context.Context, req *public_apiv1.PublicGetTenantRequest) (*public_apiv1.PublicGetTenantResponse, error) {
	got, err := h.tenantInteractor.Get(
		ctx,
		input.NewPublicGetTenant(
			req.GetTenantId(),
		),
	)
	if err != nil {
		return nil, err
	}
	return &public_apiv1.PublicGetTenantResponse{
		Tenant: marshaller.TenantToPB(got),
	}, nil
}
