package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Stud_Reg_No     string             `bson:"stud_reg_no"`
	Stud_Name       string             `bson:"stud_name"`
	Stud_Department string             `bson:"stud_department"`
	Stud_DOB        string             `bson:"stud_dob"`
	Stud_Phone      string             `bson:"stud_phone"`
	Stud_MailID     string             `bson:"stud_mail_id"`
	Books_Due       []BookDue          `bson:"books_due"`
}

type BookDue struct {
	Book_ID   string `bson:"book_id"`
	Book_Name string `bson:"book_name"`
	Quantity  int    `bson:"quantity"`
}
