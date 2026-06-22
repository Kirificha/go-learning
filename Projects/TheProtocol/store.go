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
	TotalMinutes() (int, error)
}

type MemoryStore struct {
	Sessions []Session
	ID       int
}

func (m *MemoryStore) AddSession(s Session) error {
	m.ID++
	m.Sessions = append(m.Sessions, s)
	return nil
}

func (m *MemoryStore) ListSessionsByDate(date string) ([]Session, error) {
	var goods []Session

	for _, v := range m.Sessions {
		if v.Date == date {
			goods = append(goods, v)
		}
	}

	return goods, nil
}

func (m *MemoryStore) ListSessionsByPeriod(firstDate string, lastDate string) ([]Session, error) {
	var goods []Session

	for _, v := range m.Sessions {
		if v.Date >= firstDate && v.Date <= lastDate {
			goods = append(goods, v)
		}
	}
	return goods, nil
}

func (m *MemoryStore) TotalMinutes() (int, error) {
	minutes := 0
	for _, v := range m.Sessions {
		minutes += v.Minutes
	}
	return minutes, nil
}
