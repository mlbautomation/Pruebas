package main

import "fmt"

type UserModel struct {
	nombre    string
	edad      int
	categoria string
}

type UserModels []UserModel

// Puerto por el que se van a comunicar los datos de salida
type Storage interface {
	Create(m *UserModel) error
}

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

func (u User) Create(m *UserModel) error {
	m.categoria = "Primera"

	err := u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	return nil
}
