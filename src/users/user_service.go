package users

type IUserService interface {
	findAll() *[]UserModel

	saveUser(newUser UserModel) *UserModel

	findUserById(id uint) *UserModel

	updateUser(newInfo UserModel) *UserModel

	deleteUser(id uint) *UserModel
}

type UserService struct {
}

var userRepo IUserRepository = UserRepository{}

func (_ UserService) findAll() *[]UserModel {
	return userRepo.findAll()
}

func (_ UserService) saveUser(newUser UserModel) *UserModel {
	return userRepo.saveUser(newUser)
}

func (_ UserService) findUserById(id uint) *UserModel {
	return userRepo.findUserById(id)
}

func (_ UserService) updateUser(newInfo UserModel) *UserModel {
	return userRepo.updateUser(newInfo)
}

func (uServ UserService) deleteUser(id uint) *UserModel {
	return userRepo.deleteUser(id)
}
