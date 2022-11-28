package payload

type (
	RequestCreateUser struct {
		Name       string `json:"name" validate:"required"`
		Phone      string `json:"phone" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required,min=8"`
		UserTypeID int    `json:"user_type_id" validate:"required,min=1,max=2"`
	}

	RequestEditUser struct {
		ID         int64  `json:"-" validate:"required"`
		Name       string `json:"name" validate:"required"`
		Phone      string `json:"phone" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		UserTypeID int    `json:"user_type_id" validate:"required,min=1,max=2"`
	}
)
