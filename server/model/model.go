package model
import()

type User struct{
	Username string `json:"username" firestore:"username"`
	Password string `json:"password" firestore:"password"`
}
type Student struct{
	Firstname string `json:"firstname" firestore:"firstname"`
	Lastname string `json:"lastname" firestore:"lastname"`
	StudentId string `json:"studentID" firestore:"studentID"`

}
type Status struct{
	StudentId string `json:"student_id" firestore:"student_id"`
	Status []bool `json:"status" firestore:"status"`
}