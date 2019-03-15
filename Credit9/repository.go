package credit9

type Repository interface {
	SentMessageOTP(phone string, otp string, key string) (interface{}, error)
	ValidateMessageOTP(Secret string, phone string) (interface{}, error)
}
