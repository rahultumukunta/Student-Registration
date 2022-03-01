package beans


type Student struct {
	ID          int `gorm:"primary_key"`
	Name        string
	YearOfStudy int
	Department  string
	BloodGroup  string
	PhoneNo     int
	Email       string
	Location    string
}
