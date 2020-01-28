package model

import (
	"context"

	"github.com/abyssparanoia/rapid-go/internal/pkg/log"
	"github.com/abyssparanoia/rapid-go/internal/pkg/util"
)

// Token ... token model
type Token struct {
	ID        string
	Platform  Platform
	AppID     string
	UserID    string
	DeviceID  string
	Value     string
	CreatedAt int64
}

// NewToken ... new token model
func NewToken(platform Platform,
	appID, deviceID, value string,
) *Token {
	return &Token{
		Platform:  platform,
		AppID:     appID,
		DeviceID:  deviceID,
		Value:     value,
		CreatedAt: util.TimeNowUnixMill(),
	}
}

// Platform ... platform type
type Platform string

const (
	// PlatformIOS ... ios
	PlatformIOS Platform = "ios"
	// PlatformAndroid ... android
	PlatformAndroid Platform = "android"
	// PlatformWeb ... web
	PlatformWeb Platform = "web"
)

// NewPlatform ... new platform
func NewPlatform(ctx context.Context, platform string) (Platform, error) {

	switch Platform(platform) {
	case PlatformIOS, PlatformAndroid, PlatformWeb:
		return Platform(platform), nil
	default:
	}

	return "", log.Errore(ctx, "no match case")
}

// MustPlatform ... must platform
func MustPlatform(platform string) Platform {

	switch Platform(platform) {
	case PlatformIOS, PlatformAndroid, PlatformWeb:
		return Platform(platform)
	default:
	}

	panic("no match case")
}

func (v Platform) String() string {
	return string(v)
}
