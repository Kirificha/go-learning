package main

import (
	"testing"
)

func TestMemoryStore_AddSession(t *testing.T) {
	store := &MemoryStore{}
	s := []Session{
		{Date: "2008-04-04", Minutes: 40, ID: 0, Category: CodingSession},
	}

	err := store.AddSession(s)
	if err != nil {
		t.Fatalf("Ошибка добавления в память: %v", err)
	}

	sessions, err := store.ListSessionsByDate("2008-04-04")
	if err != nil {
		t.Fatalf("Ошибка взятия из памяти: %v", err) // t.Fatal(err) или t.Fatal("ошибка %v: ", err)
	}

	if len(sessions) != 1 {
		t.Errorf("ожидалось %v, получили %v", 1, len(sessions))
	}

	if store.ID != 1 {
		t.Errorf("ожидалось %v, получили %v", 1, store.ID)
	}

	if sessions[0].Minutes != s[0].Minutes {
		t.Errorf("ожидалось %v, получили %v", s[0].Minutes, sessions[0].Minutes)
	}

}

func TestMemoryStore_ListSessionsByDate(t *testing.T) {
	store := &MemoryStore{}
	date := "2008-04-04"
	s := []Session{
		{Date: "2008-04-04", Minutes: 30, ID: 0, Category: LearningSession},
		{Date: "2003-04-04", Minutes: 30, ID: 0, Category: LearningSession},
	}

	store.AddSession(s)

	sessions, err := store.ListSessionsByDate(date)

	if err != nil {
		t.Fatalf("Возникла ошибка %v", err)
	}

	if len(sessions) != 1 {
		t.Errorf("ожидалось %v, получили %v", 1, len(sessions))
	}

	if sessions[0].Minutes != s[0].Minutes {
		t.Errorf("Ошибка")
	}

}

func TestMemoryStore_ListSessionsByPeriod(t *testing.T) {
	store := &MemoryStore{}
	s := []Session{
		{Date: "2008-04-04", Minutes: 40, Category: HybridSession, ID: 0},
		{Date: "2008-05-04", Minutes: 30, Category: HybridSession, ID: 0},
		{Date: "2007-06-04", Minutes: 20, Category: HybridSession, ID: 0},
	}

	store.AddSession(s)
	firstDate := "2008-04-04"
	lastDate := "2026-06-23"

	sessions, err := store.ListSessionsByPeriod(firstDate, lastDate)

	if err != nil {
		t.Fatalf("Ошибка %v", err)
	}

	if len(sessions) != 2 {
		t.Fatalf("Ожидали %v, получили %v", 2, len(sessions))
	}

	if sessions[0].Minutes != 40 {
		t.Errorf("Ошибка")
	}

	if sessions[1].Minutes != 30 {
		t.Errorf("Ошибка")
	}
}

func TestMemoryStore_TotalMinutes(t *testing.T) {
	store := &MemoryStore{}
	s := []Session{
		{ID: 0, Date: "2008-04-04", Minutes: 40, Category: HybridSession},
		{ID: 0, Date: "2008-04-04", Minutes: 70, Category: HybridSession},
		{ID: 0, Date: "2008-04-04", Minutes: 32, Category: HybridSession},
	}

	store.AddSession(s)
	result, err := store.TotalMinutes()

	if err != nil {
		t.Errorf("Ошибка: %v", err)
	}
	if result != 142 {
		t.Errorf("Ожидали %v, получили %v", 142, result)
	}
}
