package utils

func HashPassword(password string) {

}

func VerifyPassword() {

}

/*
func (user *User) HashPassword(password string) error {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashpassword)
	return nil
}

func (user *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
*/
