package oidc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/token/hmac"
	"github.com/ory/x/errorsx"
)

// HMACCoreStrategy implements oauth2.CoreStrategy. It's a copy of the oauth2.HMACSHAStrategy.
type HMACCoreStrategy struct {
	Enigma *hmac.HMACStrategy
	Config interface {
		fosite.AccessTokenLifespanProvider
		fosite.RefreshTokenLifespanProvider
		fosite.AuthorizeCodeLifespanProvider
	}
	prefix *string
}

// AccessTokenSignature implements oauth2.AccessTokenStrategy.
func (h *HMACCoreStrategy) AccessTokenSignature(ctx context.Context, token string) string {
	return h.Enigma.Signature(token)
}

// GenerateAccessToken implements oauth2.AccessTokenStrategy.
func (h *HMACCoreStrategy) GenerateAccessToken(ctx context.Context, _ fosite.Requester) (token string, signature string, err error) {
	token, sig, err := h.Enigma.Generate(ctx)
	if err != nil {
		return "", "", err
	}

	return h.setPrefix(token, "at"), sig, nil
}

// ValidateAccessToken implements oauth2.AccessTokenStrategy.
func (h *HMACCoreStrategy) ValidateAccessToken(ctx context.Context, r fosite.Requester, token string) (err error) {
	var exp = r.GetSession().GetExpiresAt(fosite.AccessToken)
	if exp.IsZero() && r.GetRequestedAt().Add(h.Config.GetAccessTokenLifespan(ctx)).Before(time.Now().UTC()) {
		return errorsx.WithStack(fosite.ErrTokenExpired.WithHintf("Access token expired at '%s'.", r.GetRequestedAt().Add(h.Config.GetAccessTokenLifespan(ctx))))
	}

	if !exp.IsZero() && exp.Before(time.Now().UTC()) {
		return errorsx.WithStack(fosite.ErrTokenExpired.WithHintf("Access token expired at '%s'.", exp))
	}

	return h.Enigma.Validate(ctx, h.trimPrefix(token, "at"))
}

// RefreshTokenSignature implements oauth2.RefreshTokenStrategy.
func (h *HMACCoreStrategy) RefreshTokenSignature(ctx context.Context, token string) string {
	return h.Enigma.Signature(token)
}

// GenerateRefreshToken implements oauth2.RefreshTokenStrategy.
func (h *HMACCoreStrategy) GenerateRefreshToken(ctx context.Context, _ fosite.Requester) (token string, signature string, err error) {
	token, sig, err := h.Enigma.Generate(ctx)
	if err != nil {
		return "", "", err
	}

	return h.setPrefix(token, "rt"), sig, nil
}

// ValidateRefreshToken implements oauth2.RefreshTokenStrategy.
func (h *HMACCoreStrategy) ValidateRefreshToken(ctx context.Context, r fosite.Requester, token string) (err error) {
	var exp = r.GetSession().GetExpiresAt(fosite.RefreshToken)
	if exp.IsZero() {
		return h.Enigma.Validate(ctx, h.trimPrefix(token, "rt"))
	}

	if !exp.IsZero() && exp.Before(time.Now().UTC()) {
		return errorsx.WithStack(fosite.ErrTokenExpired.WithHintf("Refresh token expired at '%s'.", exp))
	}

	return h.Enigma.Validate(ctx, h.trimPrefix(token, "rt"))
}

// AuthorizeCodeSignature implements oauth2.AuthorizeCodeStrategy.
func (h *HMACCoreStrategy) AuthorizeCodeSignature(ctx context.Context, token string) string {
	return h.Enigma.Signature(token)
}

// GenerateAuthorizeCode implements oauth2.AuthorizeCodeStrategy.
func (h *HMACCoreStrategy) GenerateAuthorizeCode(ctx context.Context, _ fosite.Requester) (token string, signature string, err error) {
	token, sig, err := h.Enigma.Generate(ctx)
	if err != nil {
		return "", "", err
	}

	return h.setPrefix(token, "ac"), sig, nil
}

// ValidateAuthorizeCode implements oauth2.AuthorizeCodeStrategy.
func (h *HMACCoreStrategy) ValidateAuthorizeCode(ctx context.Context, r fosite.Requester, token string) (err error) {
	var exp = r.GetSession().GetExpiresAt(fosite.AuthorizeCode)
	if exp.IsZero() && r.GetRequestedAt().Add(h.Config.GetAuthorizeCodeLifespan(ctx)).Before(time.Now().UTC()) {
		return errorsx.WithStack(fosite.ErrTokenExpired.WithHintf("Authorize code expired at '%s'.", r.GetRequestedAt().Add(h.Config.GetAuthorizeCodeLifespan(ctx))))
	}

	if !exp.IsZero() && exp.Before(time.Now().UTC()) {
		return errorsx.WithStack(fosite.ErrTokenExpired.WithHintf("Authorize code expired at '%s'.", exp))
	}

	return h.Enigma.Validate(ctx, h.trimPrefix(token, "ac"))
}

func (h *HMACCoreStrategy) getPrefix(part string) string {
	if h.prefix == nil {
		prefix := "ory_%s_"
		h.prefix = &prefix
	} else if len(*h.prefix) == 0 {
		return ""
	}

	return fmt.Sprintf(*h.prefix, part)
}

func (h *HMACCoreStrategy) trimPrefix(token, part string) string {
	return strings.TrimPrefix(token, h.getPrefix(part))
}

func (h *HMACCoreStrategy) setPrefix(token, part string) string {
	return h.getPrefix(part) + token
}
