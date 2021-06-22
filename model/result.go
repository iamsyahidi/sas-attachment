package model

// Result for 1 file upload
type Result struct {
	IsSuccess bool
	FileName string
}

// Results for >1 file upload
type Results struct {
	IsSuccess bool
	FileNames []string
}