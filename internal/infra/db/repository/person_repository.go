package repository

import (
	"fmt"

	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/config"
)

type PersonRepository struct {
	poll *config.Pool
}

func NewPersonRepository() *PersonRepository {
	conn := &config.Pool{}
	poll, err := conn.NewPool()
	if err != nil {
		fmt.Println("Connection Error: ", err)
		panic(err)
	}

	poll.Conn.AutoMigrate(&entity.Person{})

	return &PersonRepository{
		poll: poll,
	}
}

func (s *PersonRepository) FindAll() ([]entity.Person, error) {
	var Persons []entity.Person
	result := s.poll.Conn.Find(&Persons)
	if result.Error != nil {
		return nil, result.Error
	}

	return Persons, nil
}

func (s *PersonRepository) FindByID(id int64) (*entity.Person, error) {
	var Person entity.Person
	result := s.poll.Conn.First(&Person, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &Person, nil
}

func (s *PersonRepository) Create(Person entity.Person) (int64, error) {
	result := s.poll.Conn.Create(&Person)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *PersonRepository) Update(Person entity.Person) (int64, error) {
	result := s.poll.Conn.Model(&Person).Where("id = ?", Person.ID).Updates(&Person)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *PersonRepository) Delete(id int64) (int64, error) {
	result := s.poll.Conn.Delete(entity.Person{}, id)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
