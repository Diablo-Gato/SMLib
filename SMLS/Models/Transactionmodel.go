package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/yourusername/smls/models"
)

// MongoDB configuration
const (
	uri    = "mongodb://localhost:27017"
	dbName = "SMLS"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Choose collection and operation
	fmt.Println("Select collection:")
	fmt.Println("1. Books")
	fmt.Println("2. Students")
	fmt.Println("3. Transactions")
	var collection string
	fmt.Scanln(&collection)

	switch collection {
	case "1":
		performBookOperation(client)
	case "2":
		performStudentOperation(client)
	case "3":
		performTransactionOperation(client)
	default:
		fmt.Println("Invalid choice.")
	}
}

func performBookOperation(client *mongo.Client) {
	fmt.Println("Choose operation:")
	fmt.Println("1. Create")
	fmt.Println("2. Read")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	var operation string
	fmt.Scanln(&operation)

	switch operation {
	case "1":
		var book models.Book
		fmt.Println("Enter BookID:")
		fmt.Scanln(&book.BookID)
		fmt.Println("Enter BookName:")
		fmt.Scanln(&book.BookName)
		fmt.Println("Enter TotalNumber:")
		fmt.Scanln(&book.TotalNumber)
		fmt.Println("Enter CurrentNumber:")
		fmt.Scanln(&book.CurrentNumber)
		err := models.CreateBook(client, book)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Book record created successfully.")
	case "2":
		books, err := models.ReadBooks(client)
		if err != nil {
			log.Fatal(err)
		}
		for _, book := range books {
			fmt.Printf("BookID: %d, BookName: %s, TotalNumber: %d, CurrentNumber: %d\n", book.BookID, book.BookName, book.TotalNumber, book.CurrentNumber)
		}
	case "3":
		var book models.Book
		fmt.Println("Enter BookID to update:")
		fmt.Scanln(&book.BookID)
		fmt.Println("Enter updated BookName:")
		fmt.Scanln(&book.BookName)
		fmt.Println("Enter updated TotalNumber:")
		fmt.Scanln(&book.TotalNumber)
		fmt.Println("Enter updated CurrentNumber:")
		fmt.Scanln(&book.CurrentNumber)
		err := models.UpdateBook(client, book)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Book record updated successfully.")
	case "4":
		var bookID int
		fmt.Println("Enter BookID to delete:")
		fmt.Scanln(&bookID)
		err := models.DeleteBook(client, bookID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Book record deleted successfully.")
	default:
		fmt.Println("Invalid choice.")
	}
}

func performStudentOperation(client *mongo.Client) {
	fmt.Println("Choose operation:")
	fmt.Println("1. Create")
	fmt.Println("2. Read")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	var operation string
	fmt.Scanln(&operation)

	switch operation {
	case "1":
		var student models.Student
		fmt.Println("Enter Student Reg No:")
		fmt.Scanln(&student.RegNo)
		fmt.Println("Enter Student Name:")
		fmt.Scanln(&student.Name)
		fmt.Println("Enter Student Department:")
		fmt.Scanln(&student.Department)
		fmt.Println("Enter Student DOB:")
		fmt.Scanln(&student.DOB)
		fmt.Println("Enter Student Phone:")
		fmt.Scanln(&student.Phone)
		fmt.Println("Enter Student Email:")
		fmt.Scanln(&student.Email)
		fmt.Println("Enter Student Books Due:")
		fmt.Scanln(&student.BooksDue)
		err := models.CreateStudent(client, student)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Student record created successfully.")
	case "2":
		students, err := models.ReadStudents(client)
		if err != nil {
			log.Fatal(err)
		}
		for _, student := range students {
			fmt.Printf("RegNo: %s, Name: %s, Department: %s, DOB: %s, Phone: %s, Email: %s, BooksDue: %s\n", student.RegNo, student.Name, student.Department, student.DOB, student.Phone, student.Email, student.BooksDue)
		}
	case "3":
		var student models.Student
		fmt.Println("Enter Student Reg No to update:")
		fmt.Scanln(&student.RegNo)
		fmt.Println("Enter updated Name:")
		fmt.Scanln(&student.Name)
		fmt.Println("Enter updated Department:")
		fmt.Scanln(&student.Department)
		fmt.Println("Enter updated DOB:")
		fmt.Scanln(&student.DOB)
		fmt.Println("Enter updated Phone:")
		fmt.Scanln(&student.Phone)
		fmt.Println("Enter updated Email:")
		fmt.Scanln(&student.Email)
		fmt.Println("Enter updated BooksDue:")
		fmt.Scanln(&student.BooksDue)
		err := models.UpdateStudent(client, student)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Student record updated successfully.")
	case "4":
		var regNo string
		fmt.Println("Enter Student Reg No to delete:")
		fmt.Scanln(&regNo)
		err := models.DeleteStudent(client, regNo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Student record deleted successfully.")
	default:
		fmt.Println("Invalid choice.")
	}
}

func performTransactionOperation(client *mongo.Client) {
	fmt.Println("Choose operation:")
	fmt.Println("1. Create")
	fmt.Println("2. Read")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	var operation string
	fmt.Scanln(&operation)

	switch operation {
	case "1":
		var transaction models.Transaction
		fmt.Println("Enter Book ID:")
		fmt.Scanln(&transaction.BookID)
		fmt.Println("Enter Book Name:")
		fmt.Scanln(&transaction.BookName)
		fmt.Println("Enter Total Number:")
		fmt.Scanln(&transaction.TotalNumber)
		fmt.Println("Enter Current Number:")
		fmt.Scanln(&transaction.CurrentNumber)
		fmt.Println("Enter Student Name:")
		fmt.Scanln(&transaction.StudentInfo.Name)
		fmt.Println("Enter Student Reg No:")
		fmt.Scanln(&transaction.StudentInfo.RegNo)
		fmt.Println("Enter Student Email:")
		fmt.Scanln(&transaction.StudentInfo.Email)
		fmt.Println("Enter Due Date:")
		fmt.Scanln(&transaction.DueDate)
		err := models.CreateTransaction(client, transaction)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Transaction record created successfully.")
	case "2":
		transactions, err := models.ReadTransactions(client)
		if err != nil {
			log.Fatal(err)
		}
		for _, transaction := range transactions {
			fmt.Printf("BookID: %s, BookName: %s, TotalNumber: %d, CurrentNumber: %d, Student Name: %s, RegNo: %s, Email: %s, DueDate: %s\n", transaction.BookID, transaction.BookName, transaction.TotalNumber, transaction.CurrentNumber, transaction.StudentInfo.Name, transaction.StudentInfo.RegNo, transaction.StudentInfo.Email, transaction.DueDate)
		}
	case "3":
		var transaction models.Transaction
		fmt.Println("Enter Book ID to update:")
		fmt.Scanln(&transaction.BookID)
		fmt.Println("Enter updated Book Name:")
		fmt.Scanln(&transaction.BookName)
		fmt.Println("Enter updated Total Number:")
		fmt.Scanln(&transaction.TotalNumber)
		fmt.Println("Enter updated Current Number:")
		fmt.Scanln(&transaction.CurrentNumber)
		fmt.Println("Enter updated Student Name:")
		fmt.Scanln(&transaction.StudentInfo.Name)
		fmt.Println("Enter updated Student Reg No:")
		fmt.Scanln(&transaction.StudentInfo.RegNo)
		fmt.Println("Enter updated Student Email:")
		fmt.Scanln(&transaction.StudentInfo.Email)
		fmt.Println("Enter updated Due Date:")
		fmt.Scanln(&transaction.DueDate)
		err := models.UpdateTransaction(client, transaction)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Transaction record updated successfully.")
	case "4":
		var bookID string
		fmt.Println("Enter Book ID to delete:")
		fmt.Scanln(&bookID)
		err := models.DeleteTransaction(client, bookID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Transaction record deleted successfully.")
	default:
		fmt.Println("Invalid choice.")
	}
}
