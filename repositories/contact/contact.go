package contactrepository

import (
	"api-rs/models"
	"api-rs/schemas"
	"errors"
	"math"

	"gorm.io/gorm"
)

type ContactRepository interface {
	ListContactPagination(request schemas.Common) (roles []*models.Contact, meta *schemas.Meta, err error)
	SaveContact(contact models.Contact) (err error)
	GetContact(ID uint64) (contact *models.Contact, err error)
	DeleteContact(contact models.Contact) (err error)
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *contactRepository {
	return &contactRepository{db}
}

func (r *contactRepository) ListContactPagination(request schemas.Common) (contacts []*models.Contact, meta *schemas.Meta, err error) {
	var (
		offset = (request.Page - 1) * request.Limit
		db     = r.db.Model(&contacts)
		count  int64
	)

	if request.Search != nil && *request.Search != "" {
		search := "%" + *request.Search + "%"
		db = db.Where("name LIKE ? OR value LIKE ?", search, search)
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	db.Offset(int(offset)).Limit(int(request.Limit))

	if err := db.Find(&contacts).Error; err != nil {
		return nil, nil, err
	}

	meta = &schemas.Meta{
		Page:       request.Page,
		PerPage:    request.Limit,
		TotalPages: int64(math.Ceil(float64(count) / float64(request.Limit))),
		TotalRows:  count,
	}

	return contacts, meta, nil
}

func (r *contactRepository) SaveContact(contact models.Contact) (err error) {
	err = r.db.Save(&contact).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *contactRepository) GetContact(ID uint64) (contact *models.Contact, err error) {
	err = r.db.First(&contact, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact not found")
		}

		return nil, err
	}

	return contact, nil
}

func (r *contactRepository) DeleteContact(contact models.Contact) (err error) {
	err = r.db.Delete(&contact).Error
	if err != nil {
		return err
	}

	return nil
}
