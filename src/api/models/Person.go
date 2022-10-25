package models

import "teste1/src/api/security"

type Person struct {
	ID       uint32 `gorm:primary_key;auto_increment" json:"id"`
	Age      uint32 `json:"age"`
	Username string `gorm:"size:64;not null;unique" json:"username"`
	Password string `gorm:"size:64;not null" json:"password"`
	Family   string `gorm:"size:64" json:"family"`
	Role     string `gorm:"size:64;not null" json:"role"`
}

func (p *Person) BeforeSave() error {
	hashedPassword, err := security.Hash(p.Password)
	if err !=nil{
		return err
	}

	p.Password = string(hashedPassword)
	return nil
}