package users

import (
	"fmt"
	"users_api/src/crypto"
	"users_api/src/errorss"
)

type IUserService interface {
	findAll() *[]UserModel

	save(newUser UserModel) *UserModel

	FindById(id uint) *UserModel

	update(newInfo *UserModel) *UserModel

	delete(id uint) *UserModel

	activate(id uint, code string) *errorss.ErrorResponseModel

	login(toLoggin *UserModel) (token string, user *UserModel)

	CheckRol(rolToSearch []string, token *crypto.TokenModel) bool
}

type UserService struct {
}

var userRepo IUserRepository = UserRepository{}

var badCredentials = errorss.ErrorResponseModel{HttpStatus: 400, Cause: "Bad credentials"}
var userNotFound = errorss.ErrorResponseModel{HttpStatus: 404, Cause: "User not found"}

func (UserService) findAll() *[]UserModel {
	return userRepo.findAll()
}

func (UserService) save(newUser UserModel) *UserModel {

	//plain, hash := crypto.GenerateRandomHash()
	//Logger.LogF(true, "Activation code:", plain)
	//newUser.ActivationHash = hash

	passHash := crypto.GetHash(newUser.Password)
	newUser.Password = passHash

	return userRepo.saveUser(&newUser)
}

func (UserService) FindById(id uint) *UserModel {
	userFinded := userRepo.findUserById(id)
	if userFinded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 404, Cause: "User not found"})
	}
	return userFinded
}

func (uServ UserService) update(newInfo *UserModel) *UserModel {
	oldUser := uServ.FindById(newInfo.ID)
	if oldUser == nil {
		panic(userNotFound)
	}

	return userRepo.updateUser(oldUser, newInfo)
}

func (uServ UserService) delete(id uint) *UserModel {
	toDel := uServ.FindById(id)
	if toDel == nil {
		panic(userNotFound)
	}

	userRepo.deleteUser(toDel.ID)

	return toDel
}

func (uServ UserService) activate(id uint, code string) *errorss.ErrorResponseModel {
	userFinded := uServ.FindById(id)
	if userFinded == nil {
		return &badCredentials
	}

	if uServ.isUserActiivated(userFinded) {
		return &errorss.ErrorResponseModel{HttpStatus: 400, Cause: "User already activated"}
	}

	isMatch := crypto.PasswordMatches(userFinded.ActivationHash, code)
	if isMatch {
		userRepo.updateUser(userFinded, &UserModel{ActivationHash: "_"})
		return nil
	} else {
		return &badCredentials
	}

}

func (uServ UserService) login(toLoggin *UserModel) (token string, user *UserModel) {
	userFinded := userRepo.findByEmail(toLoggin.Email)
	if userFinded == nil {
		panic(badCredentials)
	}

	// if !uServ.isUserActiivated(userFinded) {
	//	 panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "Email validation requied"})
	// }

	isMatch := crypto.PasswordMatches(userFinded.Password, toLoggin.Password)
	if !isMatch {
		panic(badCredentials)
	}

	token = crypto.GenerateToken(fmt.Sprint(userFinded.ID))

	return token, userFinded
}

func (uServ UserService) isUserActiivated(u *UserModel) bool {
	return u.ActivationHash == "_" || u.ActivationHash == ""
}

func (uServ UserService) CheckRol(rolToSearch []string, token *crypto.TokenModel) bool {
	userLogged := usrServ.FindById(token.IdUser)
	for i := range userLogged.Rols {
		for k := range rolToSearch {
			if userLogged.Rols[i].Name == rolToSearch[k] {
				return true
			}
		}
	}
	return false
}
