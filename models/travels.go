package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	u "github.com/lazaropj/rmad_api/utils"
)

type Travel struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Route       string    `json:"route"`
	StartDate   time.Time `json:"start_date"`
	FinishDate  time.Time `json:"finish_date"`
	Code        string    `json:"code"gorm:"unique"`
	UserId      uint      `json:"user_id"` //The user that this contact belongs to
}

func (travel *Travel) Validate() (map[string]interface{}, bool) {

	if travel.Title == "" {
		return u.Message(false, "Travel title should be on the payload"), false
	}

	if travel.Route == "" {
		return u.Message(false, "Travel route should be on the payload"), false
	}

	if travel.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	fmt.Println("Travel validated", travel.StartDate)

	return u.Message(true, "success"), true
}

func (travel *Travel) Create() map[string]interface{} {

	if resp, ok := travel.Validate(); !ok {
		return resp
	}

	GetDB().Create(travel)

	resp := u.Message(true, "success")
	resp["travel"] = travel
	return resp
}

func GetTravel(clause string, value string) *Travel {

	travel := &Travel{}
	err := GetDB().Table("travels").Where(clause+" = ?", value).First(travel).Error
	if err != nil {
		return nil
	}

	return travel
}

func GetTravels(user uint) []*Travel {

	travels := make([]*Travel, 0)
	err := GetDB().Table("travels").Where("user_id = ?", user).Find(&travels).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return travels
}

func FinishTravel(id string) *Travel {

	travel := &Travel{}
	err := GetDB().Table("travels").Where("id = ?", id).First(travel).Error
	if err != nil {
		return nil
	}

	travel.FinishDate = time.Now()
	travel.Code = u.GenerateRandomString(10)
	fmt.Println("Travel finished:", travel)

	GetDB().Save(travel)

	return travel
}
