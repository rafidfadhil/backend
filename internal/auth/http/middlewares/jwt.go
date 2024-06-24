package middlewares

import (
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/BIC-Final-Project/backend/internal/auth/entity"
	"github.com/golang-jwt/jwt/v5"
)

const (
    privateKeyPath = "configs/keys/private.pem"
    publicKeyPath  = "configs/keys/public.pem"
)

var (
    verifyKey *rsa.PublicKey
    signKey   *rsa.PrivateKey
)

func init() {
    signBytes, err := os.ReadFile(privateKeyPath)
    if err != nil {
        log.Fatalf("Failed to read private key from %s: %v", privateKeyPath, err)
    }
    signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
    if err != nil {
        log.Fatalf("Failed to parse private key: %v", err)
    }

    verifyBytes, err := os.ReadFile(publicKeyPath)
    if err != nil {
        log.Fatalf("Failed to read public key from %s: %v", publicKeyPath, err)
    }
    verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
    if err != nil {
        log.Fatalf("Failed to parse public key: %v", err)
    }
}

// JWTClaim represents the claims in the JWT token.
type JWTClaim struct {
    NamaLengkap string `json:"nama_lengkap"`
    NoHandphone string `json:"no_handphone"`
    Email       string `json:"email"`
    Role        string `json:"role"`
    jwt.RegisteredClaims
}

// SignJWT signs a JWT token for the given user and duration.
func SignJWT(user entity.User, hour time.Duration) (string, error) {
    expirationTime := jwt.NewNumericDate(time.Now().Add(hour * time.Hour))
    claims := &JWTClaim{
        NamaLengkap: user.NamaLengkap,
        NoHandphone: user.NoHandphone,
        Email:       user.Email,
        Role:        user.Role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: expirationTime,
        },
    }
    
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

    tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

    return tokenString, nil
}

func VerifyJWT(tokenString string) (*JWTClaim, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
        return verifyKey, nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaim)
    if !ok || !token.Valid {
        return nil, jwt.ErrInvalidKey
    }

    return claims, nil
}