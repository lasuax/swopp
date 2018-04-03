package models

import "github.com/revel/revel"

// User model
type User struct {
	UserName        string
	FirstName       string
	LastName        string
	Age             int
	Password        string
	PasswordConfirm string
	Email           string
	EmailConfirm    string
	TermsOfUse      bool
}

// LoginModel model
type LoginModel struct {
	Username string
	Password string
}

// Validate loginmodel struct
func (model *LoginModel) Validate(v *revel.Validation) {
	v.Required(model.Username).Message("Username is required")
	v.MinSize(model.Username, 3).Message("Username must be 3 characters at least")
	v.Required(model.Password).Message("Password is required")
}

// Validate user struct
func (user *User) Validate(v *revel.Validation) {
	v.Required(user.UserName).Message("Username is required")
	v.MinSize(user.UserName, 3).Message("Username must be 3 characters at least")
	v.Required(user.FirstName).Message("Firstname is required")
	v.Required(user.LastName).Message("Lastname is required")
	v.Required(user.Age)
	v.Range(user.Age, 13, 201)
	v.Required(user.Password)
	v.Required(user.PasswordConfirm)
	v.MinSize(user.Password, 6)
	v.Required(user.Password == user.PasswordConfirm).Message("Passwords do not match")
	v.Email(user.Email)
	v.Required(user.EmailConfirm)
	v.Required(user.Email == user.EmailConfirm).Message("Emails do not match")
	v.Required(user.TermsOfUse)
}
