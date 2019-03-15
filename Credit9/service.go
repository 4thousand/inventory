package credit9

import (
	"crypto/hmac"

	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}
type ValidateOpts struct {
	Digits    otp.Digits
	Algorithm otp.Algorithm
}

type Service interface {
	SentMessageOTP(phone string) (interface{}, error)
	ValidateMessageOTP(Secret string, phone string) (interface{}, error)
}

func (s *service) ValidateMessageOTP(Secret string, phone string) (interface{}, error) {
	repo, err := s.repo.ValidateMessageOTP(Secret, phone)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (s *service) SentMessageOTP(phone string) (interface{}, error) {

	if phone[0] != '+' {
		return nil, errors.New("Please input county code Phone number")
	}
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: phone,
	})

	otp, _ := GenerateCodeCustom(key.Secret(), time.Now().UTC(), ValidateOpts{
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	fmt.Println(key.Secret(), otp)
	repo, err := s.repo.SentMessageOTP(phone, otp, key.Secret())
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func GenerateCodeCustom(secret string, t time.Time, opts ValidateOpts) (passcode string, err error) {

	counter := uint64(math.Floor(float64(t.Unix()) / float64(30)))
	secret = strings.TrimSpace(secret)
	if n := len(secret) % 8; n != 0 {
		secret = secret + strings.Repeat("=", 8-n)
	}

	secret = strings.ToUpper(secret)

	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", otp.ErrValidateSecretInvalidBase32
	}

	buf := make([]byte, 8)
	mac := hmac.New(opts.Algorithm.Hash, secretBytes)
	binary.BigEndian.PutUint64(buf, counter)

	mac.Write(buf)
	sum := mac.Sum(nil)

	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))

	l := opts.Digits.Length()
	mod := int32(value % int64(math.Pow10(l)))

	return opts.Digits.Format(mod), nil
}
