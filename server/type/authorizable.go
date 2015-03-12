package servertype

import (
	"github.com/helderfarias/oauthprovider-go/encode"
	"github.com/helderfarias/oauthprovider-go/http"
	"github.com/helderfarias/oauthprovider-go/model"
	"time"
)

type Authorizable interface {
	FindByCredencials(clientId, clientSecret string) *model.Client

	FindRefreshTokenById(refreshToken string) *model.RefreshToken

	IssuerAccessToken() string

	RevokeToken(request http.Request) error

	DeleteTokens(refreshToken *model.RefreshToken, accessToken *model.AccessToken)

	IssuerExpireTimeForAccessToken() time.Time

	IssuerExpireTimeForRefreshToken() time.Time

	StoreAccessToken(token *model.AccessToken)

	StoreRefreshToken(token *model.RefreshToken)

	HasGrantType(identified string) bool

	CreateResponse(accessToken *model.AccessToken, refreshToken *model.RefreshToken) encode.Message
}