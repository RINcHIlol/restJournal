package restJournal

type User struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Surname     string `json:"surname" db:"surname" binding:"required"`
	Email       string `json:"email" db:"email" binding:"required"`
	Password    string `json:"password" db:"password" binding:"required"`
	Role        string `json:"role" db:"role" binding:"required"`
	SpecialtyId int    `json:"specialty_id" db:"specialty_id" binding:"required"`
}

// нет в бд
type UserParse struct {
	Name      string `db:"name"`
	Surname   string `db:"surname"`
	Email     string `db:"email"`
	Role      string `db:"role"`
	Specialty string `db:"specialty"`
}

type Group struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
	SpecialtyId int    `json:"specialty_id" db:"specialty_id" binding:"required"`
	MaxCapacity int    `json:"max_capacity" db:"max_capacity" binding:"required"`
}

type GroupUser struct {
	Id      int `json:"-" db:"id"`
	GroupId int `json:"group_id" db:"group_id" binding:"required"`
	UserId  int `json:"user_id" db:"user_id" binding:"required"`
}

type Specialty struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
}

type Subjects struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
}

type SpecialtySubjects struct {
	SpecialtyId int `json:"specialty_id" db:"specialty_id" binding:"required"`
	SubjectId   int `json:"subject_id" db:"subject_id" binding:"required"`
}
