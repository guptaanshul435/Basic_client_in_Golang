package models

type User struct {
	UserId int `json:"userid,omitempty"`
	Name string `json:"name,omitempty"`
	Address string  `json:"address,omitempty"`
    GmailId  string  `json:"emailid,omitempty"`
	PhoneNo  string  `json:"number"`
}
