package repository

import (
	"gin-gonic-api/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindAllAccount() ([]dao.Account, error)
	FindAccountById(id int) (dao.Account, error)
	Save(account *dao.Account) (dao.Account, error)
	DeleteById(id int) error
}

type AccountRepositoryImpl struct {
	db *gorm.DB
}

func (u AccountRepositoryImpl) FindAllAccount() ([]dao.Account, error) {
	var accounts []dao.Account

	var err = u.db.Find(&accounts).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return accounts, nil
}

func (u AccountRepositoryImpl) FindAccountById(id int) (dao.Account, error) {
	account := dao.Account{
		ID: id,
	}
	err := u.db.First(&account).Error
	if err != nil {
		log.Error("Got and error when find account by id. Error: ", err)
		return dao.Account{}, err
	}
	return account, nil
}

func (u AccountRepositoryImpl) Save(account *dao.Account) (dao.Account, error) {
	var err = u.db.Save(account).Error
	if err != nil {
		log.Error("Got an error when save account. Error: ", err)
		return dao.Account{}, err
	}
	return *account, nil
}

func (u AccountRepositoryImpl) DeleteById(id int) error {
	err := u.db.Delete(&dao.Account{}, id).Error
	if err != nil {
		log.Error("Got an error when delete account. Error: ", err)
		return err
	}
	return nil
}

func AccountRepositoryInit(db *gorm.DB) *AccountRepositoryImpl {
	db.AutoMigrate(&dao.Account{})
	return &AccountRepositoryImpl{
		db: db,
	}
}
