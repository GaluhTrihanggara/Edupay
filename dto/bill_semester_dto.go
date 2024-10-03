package dto

type BillSemesterResponseDto struct {
	StudentId string  `json:"student_id"` // Relation to siswa
	Semester  string  `json:"semester"`
	Year      string  `json:"year"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	DueDate   string  `json:"due_date"`
}
