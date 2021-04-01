package types

//Money - represents a monetary amount
//in minimum units (cents, kopecks, diramas, etc.).
type Money int64

//Category - represents the category in which
//the payment was made(auto, pharmacies, restaurants etc.).
type Category string

//Payment - represents information about the payment source.
type Payment struct {
	ID       string
	Amount   Money
	Category Category
}
