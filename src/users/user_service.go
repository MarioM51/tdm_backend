package users

import (
	"fmt"
	"users_api/src/crypto"
	"users_api/src/errorss"
)

type IUserService interface {
	findAll() *[]UserModel

	save(newUser UserModel) *UserModel

	findById(id uint) *UserModel

	update(newInfo *UserModel) *UserModel

	delete(id uint) *UserModel

	activate(id uint, code string) *errorss.ErrorResponseModel

	login(toLoggin *UserModel) (token string)
}

type UserService struct {
}

var userRepo IUserRepository = UserRepository{}

var badCredential = errorss.ErrorResponseModel{HttpStatus: 401, Cause: "Bad credentials"}
var userNotFound = errorss.ErrorResponseModel{HttpStatus: 404, Cause: "User not found"}

func (_ UserService) findAll() *[]UserModel {
	return userRepo.findAll()
}

func (_ UserService) save(newUser UserModel) *UserModel {

	plain, hash := crypto.GenerateRandomHash()
	Logger.LogF(true, "Activation code:", plain)
	newUser.ActivationHash = hash

	passHash := crypto.GetHash(newUser.Password)
	newUser.Password = passHash

	return userRepo.saveUser(&newUser)
}

func (_ UserService) findById(id uint) *UserModel {
	return userRepo.findUserById(id)
}

func (uServ UserService) update(newInfo *UserModel) *UserModel {
	oldUser := uServ.findById(newInfo.ID)
	if oldUser == nil {
		panic(userNotFound)
	}

	return userRepo.updateUser(oldUser, newInfo)
}

func (uServ UserService) delete(id uint) *UserModel {
	toDel := uServ.findById(id)
	if toDel == nil {
		panic(userNotFound)
	}

	userRepo.deleteUser(toDel.ID)

	return toDel
}

func (uServ UserService) activate(id uint, code string) *errorss.ErrorResponseModel {
	userFinded := uServ.findById(id)
	if userFinded == nil {
		return &badCredential
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
		return &badCredential
	}

}

func (uServ UserService) login(toLoggin *UserModel) (token string) {
	userFinded := userRepo.findByEmail(toLoggin.Email)
	if userFinded == nil {
		panic(badCredential)
	}

	if userFinded.ActivationHash != "_" && userFinded.ActivationHash != "" {
		panic(errorss.ErrorResponseModel{HttpStatus: 401, Cause: "Email validation requied"})
	}

	isMatch := crypto.PasswordMatches(userFinded.Password, toLoggin.Password)
	if !isMatch {
		panic(badCredential)
	}

	token = crypto.GenerateToken(fmt.Sprint(userFinded.ID))

	return token
}
