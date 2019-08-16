package photoprism

import (
	"github.com/jinzhu/gorm"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	db *gorm.DB
}

func NewUserManager(conf config.Config) (*UserManager, error) {
	usermgr := UserManager{}
	usermgr.db = conf.Db()

	return &usermgr, nil
}

func (state *UserManager) HasUser(username string) bool {
	if err := state.db.Where("username=?", username).Find(&models.User{}).Error; err != nil {
		return false
	}
	return true
}

func (state *UserManager) FindUser(username string) *models.User {
	user := models.User{}
	state.db.Where("username=?", username).Find(&user)
	return &user
}

func (state *UserManager) FindUserByUUID(uuid string) *models.User {
	user := models.User{}
	state.db.Where("uuid=?", uuid).Find(&user)
	return &user
}

func (state *UserManager) AddUser(username, password string) *models.User {
	passwordHash := state.HashPassword(username, password)
	guid := uuid.NewV4()
	user := &models.User{
		Username: username,
		Password: passwordHash,
		UUID:     guid.String(),
	}
	state.db.Create(&user)
	return user
}

func (state *UserManager) HashPassword(username, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}

func (state *UserManager) CheckPassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false
	}
	return true
}
