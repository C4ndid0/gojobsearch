package schemas

import "time"

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int
}

type OpenigResponse struct {
  ID uint 
  CreatedAt time.Time
  UpdatadAt time.Time
  Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int

}
