package entity

type Assignment struct {
	ID       int64  `json:"id"`
	PersonID int64  `json:"personId"`
	RoleID   int64  `json:"roleId"`
	Date     string `json:"date"`
}
