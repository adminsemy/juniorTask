package model

type User struct {
	ID         int
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Nationality
}
