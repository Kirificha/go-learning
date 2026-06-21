package main

type Session struct {
	Date     string
	Minutes  int
	ID       int
	Category SessionType
}

type SessionType string

const (
	CodingSession   SessionType = "coding"
	LearningSession SessionType = "learning"
	HybridSession   SessionType = "hybrid"
)

type Store interface {
	AddSession(s Session) error
	ListSessionsByDate(date string) ([]Session, error)
	ListSessionsByPeriod(firstDate string, lastDate string) ([]Session, error)
	TotalMinutes() int
}
