package services

import (
	"context"
	"errors"
	"fmt"
	"job-portal-api/internal/database"
	"job-portal-api/internal/models"
	"job-portal-api/internal/pkg"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) CreateUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	//method that creates a new record in  db
	hashedPass, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return models.User{}, err
	}
	//prepare user record
	userDetails := models.User{
		Name:         userData.Name,
		Email:        userData.Email,
		PasswordHash: string(hashedPass),
	}
	userDetails, err = s.userRepo.CreateUser(userDetails)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil
}
func (s *Service) UserLogin(ctx context.Context, email, password string) (jwt.RegisteredClaims, error) {
	//checking the email in database
	userDetails, err := s.userRepo.UserLogin(email)
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	// We check if the provided password matches the hashed password in the database.
	err = pkg.CheckPassword(password, userDetails.PasswordHash)
	if err != nil {
		log.Info().Err(err).Send()
		return jwt.RegisteredClaims{}, errors.New("entered password is wrong")
	}
	// Successful authentication! Generate JWT claims.
	claims := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// And return those claims.
	return claims, nil

	// token, err := s.auth.GenerateAuthToken(claims)
	// if err != nil {
	// 	return "", err
	// }

	// return token, nil
}
func (s *Service) CheckEmail(email string) (bool, error) {
	_, err := s.userRepo.CheckUserEmail(email)
	if err != nil {
		log.Error().Err(err).Msg("email is not there in database")
		return false, err
	}
	otp := generateOTP()
	otpString := strconv.Itoa(otp)

	from := "mtejaswini243@gmail.com"
	password := "rdmv becs jhsn yzlh"

	// Recipient's email address
	to := email

	// SMTP server and port
	smtpServer := "smtp.gmail.com"
	smtpPort := 587
	subject := "Regarding reset password"
	body := fmt.Sprintf("one time password :%s", otpString)
	auth := smtp.PlainAuth("", from, password, smtpServer)
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)
	err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return false, err
	}

	fmt.Println("Email sent successfully.")
	rc := database.RedisClient()
	ctx := context.Background()
	err = rc.Set(ctx, email, otpString, 5*time.Minute).Err()
	if err != nil {
		log.Error().Msgf("%v", err)
		return false, err
	}
	r, _ := rc.Get(ctx, email).Result()
	fmt.Println(r)
	return true, nil
}
func (s *Service) VerifyOtp(reset models.Reset) (bool, error) {

	redis := database.RedisClient()

	cacheOtp, err := redis.Get(context.Background(), reset.Email).Result()

	if err != nil {
		log.Error().Err(err)
		return false, err
	}

	if cacheOtp == reset.Otp {
		_, err = s.userRepo.UpdatePassword(reset.Email, reset.NewPassword)
		if err != nil {
			return false, err
		}
	}

	from := "mtejaswini243@gmail.com"
	password := "rdmv becs jhsn yzlh"

	// Recipient's email address
	to := reset.Email

	// SMTP server and port
	smtpServer := "smtp.gmail.com"
	smtpPort := 587
	subject := "Reset passoword"
	body := fmt.Sprintf(":%s", "password has been successfully reset")
	auth := smtp.PlainAuth("", from, password, smtpServer)
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)
	err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return false, err
	}

	return true, nil

}

// return false, err

func generateOTP() int {
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number
	otp := rand.Intn(900000) + 100000
	return otp
}
