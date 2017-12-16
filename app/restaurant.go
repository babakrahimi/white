package app

type (
	Restaurant struct {
		Name        string `json:"name"`
		IsActive    bool   `json:"isActive"`
		DeliveryFee int    `json:"deliveryFee"`
		Foods       []Food `json:"foods"`
		Description string `json:"description"`
	}

	Food struct {
		Name        string `json:"name"`
		Fee         int    `json:"fee"`
		Description string `json:"description"`
	}
)

func GetRestaurants() []Restaurant {
	restaurants := []Restaurant{
		{Name: "نان و نمک", IsActive: false},
		{Name: "ای تی اف", IsActive: true, Foods: []Food{
			{Name: "برگر", Fee: 18000},
		}},
	}
	return restaurants
}
