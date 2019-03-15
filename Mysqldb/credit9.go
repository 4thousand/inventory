package mysqldbs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/pquerna/otp/totp"

	"github.com/pquerna/otp"

	credit9 "gitlab.com./c9/Credit9"
)

type credit9Repository struct{ db *sqlx.DB }

func NewCredit9Repository(db *sqlx.DB) credit9.Repository {
	return &credit9Repository{db}
}

type ValidateOpts struct {

	// Digits as part of the input. Defaults to 6.
	Digits otp.Digits
	// Algorithm to use for HMAC. Defaults to SHA1.
	Algorithm otp.Algorithm
}

func (repo *credit9Repository) SentMessageOTP(phone string, otps string, key string) (resp interface{}, err error) {

	message := map[string]interface{}{
		"to":        phone,
		"from":      "Nopadol",
		"text":      otps,
		"apiKey":    "b6c39957-75c1-4234-9fab-f7576bcf24ee",
		"apiSecret": "44e49e15-01d4-483e-ab46-358d991cb15a",
	}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err = http.Post("https://api.apitel.co/sms", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	sql := `insert into keyprivate(phone,pvkey) values(?,?)`
	resp, err = repo.db.Exec(sql, phone, key)
	if err != nil {
		fmt.Println("error = ", err.Error())
	}
	fmt.Println("sql = ", sql)

	return map[string]interface{}{
		"response": "ok",
	}, nil

}
func (repo *credit9Repository) ValidateMessageOTP(Secret string, phone string) (resp interface{}, err error) {
	var keys string
	sql := `select pvkey from keyprivate where phone = ? order by id desc limit 1`
	err = repo.db.Get(&keys, sql, phone)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	fmt.Println(keys)
	valid := totp.Validate(Secret, keys)
	if valid {
		return map[string]interface{}{
			"response": "ok",
		}, nil

	} else {
		return map[string]interface{}{
			"response": "not math",
		}, nil

	}

}
