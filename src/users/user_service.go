package users

import (
	"fmt"
	"users_api/src/crypto"
	"users_api/src/errorss"
)

type IUserService interface {
	findAll() *[]UserModel

	saveUser(newUser UserModel) *UserModel

	findUserById(id uint) *UserModel

	updateUser(newInfo *UserModel) *UserModel

	deleteUser(id uint) *UserModel

	activate(id uint, code string) *errorss.ErrorResponseModel
}

type UserService struct {
}

var userRepo IUserRepository = UserRepository{}

func (_ UserService) findAll() *[]UserModel {
	return userRepo.findAll()
}

func (_ UserService) saveUser(newUser UserModel) *UserModel {

	plain, hash := crypto.GenerateToken()
	Logger.LogF(true, "Activation code:", plain)
	newUser.ActivationHash = hash

	passHash := crypto.GetHash(newUser.Password)
	newUser.Password = passHash

	return userRepo.saveUser(&newUser)
}

func (_ UserService) findUserById(id uint) *UserModel {
	return userRepo.findUserById(id)
}

func (uServ UserService) updateUser(newInfo *UserModel) *UserModel {
	oldUser := uServ.findUserById(newInfo.ID)
	if oldUser == nil {
		panic(errorss.ErrorResponseModel{
			HttpStatus: 404,
			Cause:      fmt.Sprintf("User id %v not found", newInfo.ID),
		})
	}

	return userRepo.updateUser(oldUser, newInfo)
}

func (uServ UserService) deleteUser(id uint) *UserModel {
	return userRepo.deleteUser(id)
}

func (uServ UserService) activate(id uint, code string) *errorss.ErrorResponseModel {
	errResp := &errorss.ErrorResponseModel{HttpStatus: 401, Cause: "Bad credentials"}

	userFinded := uServ.findUserById(id)
	if userFinded == nil {
		return errResp
	}

	if userFinded.ActivationHash == "_" {
		return &errorss.ErrorResponseModel{HttpStatus: 400, Cause: "User already activated"}
	}

	Logger.LogF(true, "To compare: in db: %v, income: %v\n", userFinded.ActivationHash, code)

	isMatch := crypto.PasswordMatches(userFinded.ActivationHash, code)
	if isMatch {
		userRepo.updateUser(userFinded, &UserModel{ActivationHash: "_"})
		return nil
	} else {
		return errResp
	}

}
