package controllers

import (
	"Swopp/app/models"
	"fmt"

	"github.com/revel/revel"
)

// Account class
type Account struct {
	*revel.Controller
}

// Register controller
func (c Account) Register(user models.User) revel.Result {
	//c.Params.BindJSON(&user)
	//model, _ := reflect.ValueOf(user).Interface().(models.User)
	test, _ := c.Request.GetForm()
	fmt.Print(test)
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Render(user)
	}

	return c.Redirect(App.Index, user)
}

// Login page
func (c Account) Login(loginModel *models.LoginModel) revel.Result {
	return c.Render(loginModel)
}

// LoginUser submit login info
func (c Account) LoginUser() revel.Result {
	// Create model with parameters
	model := models.LoginModel{
		UserName: c.Params.Get("UserName"),
		Password: c.Params.Get("Password"),
	}
	// User login logic
	loginStatus := userLogin(model)

	// Return to login screen with model if user login fails
	if !loginStatus {
		return c.Redirect(Account.Login, model)
	}

	// After login logic
	return c.Render(model)
}

// Login service pseudo
func userLogin(model models.LoginModel) bool {
	return model.UserName == "doa" && model.Password == "1234"
}

// RegisterUser API
func (c Account) RegisterUser(user models.User) revel.Result {
	fmt.Printf(user.FirstName)
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		c.Redirect(Account.Register)
	}

	return c.Render(user)
}
