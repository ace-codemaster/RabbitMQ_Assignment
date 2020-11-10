package services

import (
	"EasternEnterprise/helpers/dbhelper"
	"EasternEnterprise/model"
	"fmt"
)

// SaveHotelData for saving hotel , room and rateplan object
func SaveHotelData(receivedData *model.Transmission) bool {

	if receivedData == nil {
		return false
	}
	for _, offer := range receivedData.Offers {

		err := dbhelper.HotelDB.Save(&offer.Hotel).Error
		if err != nil {
			fmt.Println("-------", err)
			return false
		}
		err = dbhelper.HotelDB.Save(&offer.Room).Error
		if err != nil {
			return false
		}

		err = dbhelper.HotelDB.Save(&offer.RatePlan).Error
		if err != nil {
			return false
		}
	}

	return true
}
