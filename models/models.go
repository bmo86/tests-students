package models

//student

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

//test

type Test struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//questions

type Question struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	TestId   string `json:"test_id"`
}

//enrollment

type Enrollment struct {
	StudentId string `json:"student_id"`
	TestId    string `json:"test_id"`
}

//answerstudent

type Answer struct {
	TestId        string `json:"test_id"`
	StudentId     string `json:"student_id"`
	QuestionId    string `json:"question_id"`
	AnswerStudent string `json:"answer_student"`
	Correct       bool   `json:"res_ok"`
}

//testScore

type TestScore struct {
	TestId    string `json:"test_id"`
	StudentId string `json:"student_id"`
	Ok        int32  `json:"ok"`
	Ko        int32  `json:"ko"`
	Total     int32  `json:"total"`
	Score     int32  `json:"score"`
}
