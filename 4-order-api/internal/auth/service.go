package auth

import (
	"cart-api/internal/user"
	"errors"
	"math/rand"
	// "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) AuthByPhone(phone string) (*user.User, error) {
	existedUser, _ := service.UserRepository.FindByPhone(phone)

	if existedUser != nil {
		return existedUser, nil
	}

	sessionId := RandStringRunes(15)
	user := &user.User{
		Phone:     phone,
		SessionId: sessionId,
		Code:      "1234",
	}
	_, err := service.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (service *AuthService) CodeVerification(sessionId, code string) (string, error) {
	existedUser, _ := service.UserRepository.FindBySessionId(sessionId, code)
	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}
	// hashedSessionId, err := bcrypt.GenerateFromPassword([]byte(sessionId), bcrypt.DefaultCost)
	// if err != nil {
	// 	return "", err
	// }
	// user := &user.User{
	// 	Phone:     existedUser.Phone,
	// 	SessionId: string(hashedSessionId),
	// 	Code:      code,
	// }
	// _, err = service.UserRepository.Create(user)
	// if err != nil {
	// 	return "", err
	// }
	return existedUser.Phone, nil
}
