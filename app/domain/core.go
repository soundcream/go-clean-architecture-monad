package domain

type (
	User struct {
		Name string `validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
		Age  int    `validate:"required,teen-person"`  // Required field, and client needs to implement our 'teen-person' tag format which we'll see later
	}
)
