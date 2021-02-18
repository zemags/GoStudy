package main

import (
	"gorm.io/gorm"
)

var comets []Comet = []Comet{
	{CometType: 1, CometShortName: "Lovejoy", Zone: 1},
}

type Comet struct {
	gorm.Model
	CometType      int
	CometShortName string `sql:"type:VARCHAR(255);not null;unique"`
	Zone           int
}
