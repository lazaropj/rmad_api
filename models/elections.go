package models

import (
	"fmt"
	"strconv"

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

type Average struct {
	Average string
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

func GetAverageByTravel(travelId string) float64 {

	var average Average

	err := GetDB().Table("elections").Select("sum(note) / count(*) as media").Where("travel_id = ?", travelId).Find(&average.Average).Error
	if err != nil {
		return 0
	}
	convertedAverage, _ := strconv.ParseFloat(average.Average, 64)

	return convertedAverage
}
