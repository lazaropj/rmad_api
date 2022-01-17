package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Election struct {
	gorm.Model
	Note      int64 `json:"note"`
	AccountId int
	Account   Account `json:"account" gorm:"foreignKey:AccountId"`
	TravelId  int
	Travel    Travel `json:"travel_id" gorm:"foreignKey:TravelId"`
}

func VoteOnTravel(travel Travel, account Account, note int64) {

	err := GetDB().Create(&Election{
		Account: account,
		Travel:  travel,
		Note:    note,
	}).Error
	if err != nil {
		fmt.Println(err)
	}

}
