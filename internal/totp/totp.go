package totp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"github.com/skip2/go-qrcode"
)

const (
	Issuer = "AuthCLI"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateSecret(username string) (string, string, error) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      Issuer,
			AccountName: username,
		},
	)

	if err != nil {
		return "", "", fmt.Errorf("failed to generate totp secret : %w", err)
	}

	return key.Secret(), key.URL(), nil
}

// this function validates a TOTP code which entered by user
func (s *Service) ValidateCode(code, secret string) bool {
	return totp.Validate(code, secret)
}

func (s *Service) GenerateKey(username string) (*otp.Key, error) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      Issuer,
			AccountName: username,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to generate totp key : %w", err)
	}

	return key, nil
}

func (s *Service) GenerateQRCode(
	username string,
	otpURL string,
) (string, error) {

	dir := "data/qrcodes"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	filePath := filepath.Join(
		dir,
		fmt.Sprintf("%s.png", username),
	)

	err := qrcode.WriteFile(
		otpURL,
		qrcode.Medium,
		256,
		filePath,
	)

	if err != nil {
		return "", err
	}

	return filePath, nil
}
