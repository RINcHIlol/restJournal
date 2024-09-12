package rest

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Surname  string `json:"surname" db:"surname" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
	Role     Roles  `json:"roles" db:"roles" binding:"required"`
}

type Roles struct {
	User        string `json:"user"`
	Teacher     string `json:"teacher"`
	HeadTeacher string `json:"head_teacher"`
}

// .l.
