package main

import (
	"checksmt/ttt"
	"fmt"
	"time"
)

type WSAD struct {
}

func (r *WSAD) Get(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("Empty string")
	}
	return fmt.Sprintf("I've got this, %s!", s), nil
}

func (r *WSAD) GetTime() (string, error) {
	s := time.Now().String()
	if s == "" {
		return "", nil
	}
	return s, nil

}

type Dependency struct {
	Repo ttt.IRepo
}

type Handler struct {
	repo ttt.IRepo
}

func NewHandler(d *Dependency) *Handler {
	return &Handler{
		repo: d.Repo,
	}
}

func (h *Handler) GetUnit(id string) (string, error) {
	result, err := h.repo.Get(id)
	if err != nil {
		return "", err
	}

	fmt.Println(result)

	return fmt.Sprintf("%v", result), nil
}

func (h *Handler) GetTimeVarUnit() (string, error) {
	result, err := h.repo.GetTime()
	if err != nil {
		return "", err
	}
	return result, nil
}

func main() {
	h := NewHandler(&Dependency{
		Repo: &WSAD{},
	})

	str, err := h.GetUnit("someID")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	str, err = h.GetTimeVarUnit()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
