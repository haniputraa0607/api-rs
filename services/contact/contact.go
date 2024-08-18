package contactservice

import (
	"api-rs/models"
	contactrepository "api-rs/repositories/contact"
	"api-rs/schemas"
	"errors"
	"strconv"
)

type ContactService interface {
	ListContactsPagination(request schemas.Common) (response []*schemas.ListContactResponse, meta *schemas.Meta, err error)
	CreateContact(request schemas.CreateUpdateContactRequest) (err error)
	DeleteContact(ID string) (err error)
	GetContact(ID string) (response *schemas.ContactResponse, err error)
	UpdateContact(ID string, request schemas.CreateUpdateContactRequest) (err error)
	ListContacts() (response []*schemas.ListContactResponse, err error)
}

type contactService struct {
	contactRepository contactrepository.ContactRepository
}

func NewContactService(
	contactRepository contactrepository.ContactRepository,
) *contactService {
	return &contactService{
		contactRepository: contactRepository,
	}
}

func (s *contactService) ListContactsPagination(request schemas.Common) (response []*schemas.ListContactResponse, meta *schemas.Meta, err error) {
	contacts, meta, err := s.contactRepository.ListContactPagination(request)
	if err != nil {
		return nil, nil, err
	}

	for _, contact := range contacts {

		response = append(response, &schemas.ListContactResponse{
			ID:    contact.ID,
			Name:  contact.Name,
			Value: contact.Value,
		})
	}

	return response, meta, nil
}

func (s *contactService) CreateContact(request schemas.CreateUpdateContactRequest) (err error) {
	contact := models.Contact{
		Name:  request.Name,
		Value: request.Value,
	}

	err = s.contactRepository.SaveContact(contact)
	if err != nil {
		return err
	}

	return nil
}

func (s *contactService) DeleteContact(ID string) (err error) {
	contactID, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	contact, err := s.contactRepository.GetContact(uint64(contactID))
	if err != nil {
		return err
	}

	err = s.contactRepository.DeleteContact(*contact)
	if err != nil {
		return err
	}

	return nil
}

func (s *contactService) GetContact(ID string) (response *schemas.ContactResponse, err error) {
	contactID, err := strconv.Atoi(ID)
	if err != nil {
		return nil, err
	}

	contact, err := s.contactRepository.GetContact(uint64(contactID))
	if err != nil {
		return nil, err
	}

	if contact == nil {
		return nil, errors.New("contact not found")
	}

	response = &schemas.ContactResponse{
		ID:    contact.ID,
		Name:  contact.Name,
		Value: contact.Value,
		Icon:  contact.Icon,
	}

	return response, nil
}

func (s *contactService) UpdateContact(ID string, request schemas.CreateUpdateContactRequest) (err error) {
	contactID, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	contact, err := s.contactRepository.GetContact(uint64(contactID))
	if err != nil {
		return err
	}

	contact.Name = request.Name
	contact.Value = request.Value

	if request.Icon != nil {
		contact.Icon = request.Icon
	}

	err = s.contactRepository.SaveContact(*contact)
	if err != nil {
		return err
	}

	return nil
}

func (s *contactService) ListContacts() (response []*schemas.ListContactResponse, err error) {
	contacts, err := s.contactRepository.ListContact()
	if err != nil {
		return nil, err
	}

	for _, contact := range contacts {
		response = append(response, &schemas.ListContactResponse{
			ID:    contact.ID,
			Name:  contact.Name,
			Value: contact.Value,
		})
	}

	return response, nil
}
