

# Service

### Setup debugger
- Package path:
n4a3/clean-architecture/cmd
- Working Dir:
/my-clean-templ
___
### Comment Swagger doc
````
// GetAccount @Summary Show an account
// @Description get string by ID
// @Tags Demo
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} http.Response
// @Failure 404 {object} http.Response
// @Router /accounts/{id} [get]
func GetAccount(c *gin.Context) {
    // Your code here
}
````
___
### Generate Swagger
1. cd to root project dir `` my-clean-templ``
````
cd {{my-directory}}/my-clean-templ
````
2. Run command
````
swag init -g cmd/main.go
````
___

Validator
https://github.com/go-playground/validator

