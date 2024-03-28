package function

import (
	"context"
	"fmt"
	model "myproject/model"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func Test() {
	user := model.User{Username: "admin", Password: "admin"}
	fmt.Print(user)
}
func Login(ctx *gin.Context, client *firestore.Client) error {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return err
	}
	if user.Password != "admin" || user.Username != "admin" {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}
func UpdateStatus(ctx *gin.Context, client *firestore.Client, week int, studentId string) error {
	// Define the query
	query := client.Collection("statuses").Where("student_id", "==", studentId).Limit(1)

	// Execute the query
	iter := query.Documents(ctx)
	defer iter.Stop()

	// Define a variable to hold the result
	var result model.Status

	// Retrieve the first document that matches the query
	doc, err := iter.Next()
	if err == iterator.Done {
		return fmt.Errorf("No document found")
	}
	if err != nil {
		return err
	}
	if err := doc.DataTo(&result); err != nil {
		return err
	}
	if result.Status[week-1] == true {
		return nil
	} else {
		docRef := doc.Ref
		result.Status[week-1] = true
		_, err := docRef.Update(context.Background(), []firestore.Update{
			{Path: "status", Value: result.Status},
		})
		return err
	}
}
func GetAllStudent(ctx *gin.Context, client *firestore.Client) ([]*model.Student, error) {
	var studentlist []*model.Student
	iter := client.Collection("students").Documents(context.Background())
	defer iter.Stop()
	for {
		// Retrieve the next document
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return studentlist, err
		}

		// Map the document data to a Student struct
		var student model.Student
		if err := doc.DataTo(&student); err != nil {
			return nil, err
		}

		// Append the Student struct to the slice
		studentlist = append(studentlist, &student)
	}
	return studentlist, nil
}
func GetOneStudent(ctx *gin.Context, client *firestore.Client, studentID string) *model.Status {
	query := client.Collection("statuses").Where("student_id", "==", studentID).Limit(1)

	// Execute the query
	iter := query.Documents(ctx)
	defer iter.Stop()

	// Define a variable to hold the result
	var result model.Status

	// Retrieve the first document that matches the query
	doc, err := iter.Next()
	if err == iterator.Done {
		return nil
	}
	if err != nil {
		return nil
	}
	if err := doc.DataTo(&result); err != nil {
		return nil
	} else {
		return &result
	}
}
func InsertStudent(c *gin.Context, client *firestore.Client, student *model.Student) error {
	if _, _, err := client.Collection("students").Add(context.Background(), student);err != nil{
		return err
	}
	var stt []bool
	for i:=0;i<10;i++{
		stt = append(stt,false)
	}
	status := &model.Status{
		StudentId: student.StudentId,
		Status: stt,
	}
	if _,_,err := client.Collection("statuses").Add(context.Background(),status);err != nil{
		return err
	}
	return nil
}
