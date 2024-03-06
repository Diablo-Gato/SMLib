package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Book_ID          string             `bson:"book_id"`
	Book_Name        string             `bson:"book_name"`
	Book_Total_Num   int                `bson:"book_total_num"`
	Book_Current_Num int                `bson:"book_current_num"`
}
