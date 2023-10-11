package types

type Dish struct {
	ID    int     `json:"id"`
	Name  string  `json:"dishName"`
	Time  float32 `json:"timeToPrepare"`
	Price float32 `json:"price"`
}
