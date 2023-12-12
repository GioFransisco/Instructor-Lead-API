package common

import (
	"errors"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenereteToken(userId, email, role string) (model.TokenModel, error)
	VerifyToken(model.TokenModel) (jwt.MapClaims, error)
}

type jwtToken struct {
	cfg *config.JwtConfig
}

// GenereteToken implements JwtToken.
func (j *jwtToken) GenereteToken(userId, email, role string) (model.TokenModel, error) {
	var payloadToken model.TokenModel

	claims := utilsmodel.JwtClaimsToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.Issuer,
			ExpiresAt: &jwt.NumericDate{time.Now().Add(j.cfg.TokenLifeTime * time.Hour)},
			IssuedAt:  &jwt.NumericDate{time.Now()},
		},
		UserId: userId,
		Role:   role,
		Email:  email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.cfg.SecretKey)

	if err != nil {
		return model.TokenModel{}, errors.New("failed to sign the token")
	}

	payloadToken.Token = tokenString

	return payloadToken, nil
}

// VerifyToken implements JwtToken.
func (j *jwtToken) VerifyToken(tokenPayload model.TokenModel) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenPayload.Token, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.GetSigningMethod("HS256") {
			return nil, errors.New("method not match")
		}

		return j.cfg.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)

	if !token.Valid || !ok || mapClaims["iss"] != j.cfg.Issuer {
		return nil, errors.New("token invalid")
	}

	return mapClaims, nil
}

func NewJwtToken(cfg *config.Config) JwtToken {
	return &jwtToken{&cfg.JwtConfig}
}
