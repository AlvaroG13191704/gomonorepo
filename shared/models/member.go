package models

// Model of member (Doctor, Assitant or Administrator == Doctor)
type Doctor struct {
	Uid         string      `json:"uid" bson:"uid"`
	Credentials Credentials `json:"credentials"`
	Profile     Profile     `json:"profile" bson:"profile"`
	Clinics     []Clinics   `json:"clinics" bson:"clinics"`
	Status      string      `json:"status" bson:"status"`
}

type Clinics struct {
	Id     string `json:"id" bson:"id"`
	Status string `json:"status" bson:"status"`
}

type Profile struct {
	Uid        string   `json:"uid" bson:"uid"`
	Name       string   `json:"name" bson:"name"`
	Email      string   `json:"email" bson:"email"`
	Phone      string   `json:"phone" bson:"phone"`
	Speciality []string `json:"speciality" bson:"speciality"`
	DoctorID   string   `json:"doctorID" bson:"doctorID"`
	Role       string   `json:"role" bson:"role"`
}
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
