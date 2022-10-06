package public

import (
	"context"

	"github.com/abyssparanoia/rapid-go/internal/infrastructure/grpc/internal/marshaller"
	public_apiv1 "github.com/abyssparanoia/rapid-go/internal/infrastructure/grpc/pb/rapid/public_api/v1"
	"github.com/abyssparanoia/rapid-go/internal/usecase/input"
)

func (h *PublicHandler) GetTenant(ctx context.Context, req *public_apiv1.GetTenantRequest) (*public_apiv1.GetTenantResponse, error) {
	got, err := h.tenantInteractor.Get(
		ctx,
		input.NewPublicGetTenant(
			req.GetTenantId(),
		),
	)
	if err != nil {
		return nil, err
	}
	return &public_apiv1.GetTenantResponse{
		Tenant: marshaller.TenantToPB(got),
	}, nil
}