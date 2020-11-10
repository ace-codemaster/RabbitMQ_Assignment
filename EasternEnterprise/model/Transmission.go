package model

import "gorm.io/gorm"

//Hotel fr hotel model
type Hotel struct {
	gorm.Model
	ID        string  `json:"hotel_id";gorm:"type:longtext"`
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Telephone string  `json:"telephone"`
	// Amenities   []string `json:"amenities"`
	Description string `json:"description"`
	RoomCount   int    `json:"room_count"`
	Currency    string `json:"currency"`
}

//Room for room model
type Room struct {
	gorm.Model
	HotelID string `json:"hotel_id"`
	Hotel   Hotel  `gorm:"foreignKey:ID;references:HotelID"`

	ID          string `json:"room_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	// Capacity    struct {
	// 	MaxAdults     int `json:"max_adults"`
	// 	ExtraChildren int `json:"extra_children"`
	// } `json:"capacity"`
}

//RatePlan for rate model
type RatePlan struct {
	gorm.Model
	HotelID string `json:"hotel_id"`
	Hotel   Hotel  `gorm:"foreignKey:ID;references:HotelID"`

	RatePlanID string `json:"rate_plan_id",gorm:"primaryKey"`
	// CancellationPolicy []struct {
	// 	Type              string `json:"type"`
	// 	ExpiresDaysBefore int    `json:"expires_days_before"`
	// } `json:"cancellation_policy"`
	Name string `json:"name"`
	//OtherConditions []string `json:"other_conditions"`
	MealPlan string `json:"meal_plan"`
}

//Transmission object for transmitting object
type Transmission struct {
	Offers []struct {
		CmOfferID    string   `json:"cm_offer_id"`
		Hotel        Hotel    `json:"hotel"`
		Room         Room     `json:"room"`
		RatePlan     RatePlan `json:"rate_plan"`
		OriginalData struct {
			GuaranteePolicy struct {
				Required bool `json:"Required"`
			} `json:"GuaranteePolicy"`
		} `json:"original_data"`
		Capacity struct {
			MaxAdults     int `json:"max_adults"`
			ExtraChildren int `json:"extra_children"`
		} `json:"capacity"`
		Number   int    `json:"number"`
		Price    int    `json:"price"`
		Currency string `json:"currency"`
		CheckIn  string `json:"check_in"`
		CheckOut string `json:"check_out"`
		Fees     []struct {
			Type        string  `json:"type"`
			Description string  `json:"description"`
			Included    bool    `json:"included"`
			Percent     float64 `json:"percent"`
		} `json:"fees"`
	} `json:"offers"`
}
