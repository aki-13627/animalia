package infra

import (
	"fmt"

	"github.com/htanos/animalia/backend-go/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(name, email string) (*models.User, error) {
	var existingUser models.User
	if err := r.db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("このメールアドレスは既に登録されています")
	}

	user := models.User{
		Email: email,
		Name:  name,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user in database: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) ExistsEmail(email string) (bool, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	// メールアドレスで検索し、最初の一致したレコードを取得
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetById(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(id string, name string, description string, newImageKey string) error {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	user.Name = name
	user.Bio = description
	user.IconImageKey = newImageKey

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Follow(fromId string, toId string) error {
	follow_relation := models.FollowRelation{
		FromID: fromId,
		ToID:   toId,
	}
	if err := r.db.Create(&follow_relation).Error; err != nil {
		return fmt.Errorf("failed to create follow relation in database: %w", err)
	}

	return nil
}
