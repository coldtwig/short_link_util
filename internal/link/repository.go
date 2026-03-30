package link

import (
	"go/http-api/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	DataBase *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		DataBase: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	res := repo.CheckHashInDb(link.Hash)
	if res {
		return nil, ErrorLinkUniqueHash
	}

	result := repo.DataBase.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.DataBase.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) CheckHashInDb(hash string) bool {
	var link Link
	result := repo.DataBase.DB.First(&link, "hash = ?", hash)
	return result.Error == nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.DataBase.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repo.DataBase.DB.First(&link, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}
