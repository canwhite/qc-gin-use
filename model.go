package main

import "gorm.io/gorm"

// model
type User struct {
	gorm.Model //会帮忙建id 、create and update time
	Name       string
	Age        uint
}

type Params struct {
	name  string
	email string
}
