package types

//Money - represents a monetary amount
//in minimum units (cents, kopecks, diramas, etc.).
type Money int64

//Category - represents the category in which
//the payment was made(auto, pharmacies, restaurants etc.).
type Category string

//Status - represents the status of the payments
type Status string

//Predefined payment statuses
const (
	StatusOK         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

//Payment - represents information about the payment source.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
