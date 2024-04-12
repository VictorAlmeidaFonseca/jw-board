package repository

import (
	"fmt"

	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/config"
)

type RoleRepository struct {
	poll *config.Pool
	role entity.Role
}

func NewRoleRepository() *RoleRepository {
	poll, err := config.NewPool()
	if err != nil {
		fmt.Println("Connection Error: ", err)
		panic(err)
	}

	poll.Conn.AutoMigrate(&entity.Role{})

	return &RoleRepository{
		role: entity.Role{},
		poll: poll,
	}
}

func (s *RoleRepository) FindAll() ([]entity.Role, error) {
	var roles []entity.Role
	result := s.poll.Conn.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}

func (s *RoleRepository) FindByID(id int64) (*entity.Role, error) {
	var role entity.Role
	result := s.poll.Conn.First(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (s *RoleRepository) Create(role entity.Role) (int64, error) {
	result := s.poll.Conn.Create(&role)
	if result.Error != nil {
		return 0, result.Error
	}

	return role.ID, nil
}

func (s *RoleRepository) Update(role entity.Role) (int64, error) {
	result := s.poll.Conn.Save(&role)
	if result.Error != nil {
		return 0, result.Error
	}

	return role.ID, nil
}

func (s *RoleRepository) Delete(id int64) (int64, error) {
	result := s.poll.Conn.Delete(&s.role, id)
	if result.Error != nil {
		return 0, result.Error
	}

	return id, nil
}
