package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	"github.com/Ganasa18/document-be/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

// LoginOrRegister implements AuthRepository.
func (repository *AuthRepositoryImpl) LoginOrRegister(ctx context.Context, user domain.UserModel, OpenId string) (domain.UserModel, error) {
	// Check if the user with the given email already exists or create the user
	var plainPassword string
	if OpenId != utils.OPEN_API_GOOGLE {
		plainPassword = *user.Password
	}

	err := repository.DB.Where(domain.UserModel{Email: user.Email, OpenId: user.OpenId}).First(&user).Error

	// GET USER ROLE
	_ = repository.DB.Where(domain.UserModel{Email: user.Email}).Preload("RoleMasterModel").Find(&user).Error

	// REGISTER USER
	if err != nil {
		fmt.Println("Error fetching the user:", err.Error())
		if OpenId != utils.OPEN_API_GOOGLE {

			if plainPassword == "" {
				return user, errors.New("with email must have password")
			}

			// Hashing the password with the default cost of 10
			hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
			utils.IsErrorDoPanic(errHashedPassword)
			hashedPasswordStr := string(hashedPassword)
			user.Password = &hashedPasswordStr
		}
		err = repository.DB.Create(&user).Error
		if err != nil {
			fmt.Println("Error creating user:", err.Error())
			return user, err
		}
		return user, nil
	}

	if OpenId != utils.OPEN_API_GOOGLE {

		if plainPassword == "" {
			return user, errors.New("authentication failed")
		}

		storedPasswordHash := *user.Password
		// Compare passwords
		err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(plainPassword))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Handle the case where authentication failed
			return user, errors.New("authentication failed")
		}
	}

	return user, nil

}
