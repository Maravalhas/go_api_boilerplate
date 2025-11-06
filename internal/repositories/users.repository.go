package repositories

import (
	"api/internal/models"

	"gorm.io/gorm"
)

type UsersListOptions struct {
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
	Order  string `form:"order"`
	Search string `form:"search"`
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db: db}
}

func (r *usersRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *usersRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *usersRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *usersRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *usersRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}

func (r *usersRepository) List(opts *UsersListOptions) ([]models.User, int64, error) {
	var users []models.User
	var count int64
	db := r.db.Model(&models.User{})

	if opts.Search != "" {
		db = db.Where("name ILIKE ? OR email ILIKE ?", "%"+opts.Search+"%", "%"+opts.Search+"%")
	}

	db.Count(&count)

	if opts.Order != "" {
		db = db.Order(opts.Order)
	}
	if opts.Limit > 0 {
		db = db.Limit(opts.Limit)
	}
	if opts.Offset > 0 {
		db = db.Offset(opts.Offset)
	}

	err := db.Find(&users).Error

	return users, count, err
}
