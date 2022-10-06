package model

import "github.com/volatiletech/null/v8"

type Claims struct {
	AuthUID  string
	TenantID null.String
	UserID   null.String
	UserRole UserRole
}

func NewClaims(
	authUID string,
) *Claims {
	return &Claims{
		AuthUID: authUID,
	}
}

func (m *Claims) SetTenantID(tenantID string) {
	m.TenantID = null.StringFrom(tenantID)
}

func (m *Claims) SetUserID(userID string) {
	m.UserID = null.StringFrom(userID)
}

func (m *Claims) SetUserRole(userRole UserRole) {
	m.UserRole = userRole
}