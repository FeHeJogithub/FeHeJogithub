package service

import (
	"aprenderGo/internal/model"
	"aprenderGo/internal/store"
	"errors"
)

/*
	type Logger interface {
		Log(msg, error string)
	}
*/
type Service struct {
	store store.Store
	//logger Logger
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
		//logger: nil,
	}

}

func (s *Service) ObtenTodosLosLibros() ([]*model.Libro, error) {

	//s.logger.Log("Estamos obteniendo los libros", "")
	libros, err := s.store.GetAll()
	if err != nil {
		//	s.logger.Log("el error es %v\n", err.Error())
		return nil, err
	}
	return libros, nil
}

func (s *Service) ObtenLbroPorID(id int) (*model.Libro, error) {
	return s.store.GetByID(id)
}

func (s *Service) CrearLbro(libro model.Libro) (*model.Libro, error) {

	if libro.Titulo == "" {
		return nil, errors.New("necesitamos el titulo")
	}

	return s.store.Create(&libro)
}
func (s *Service) UpdateAllLibro(id int, libro model.Libro) (*model.Libro, error) {

	if libro.Titulo == "" {
		return nil, errors.New("necesitamos el titulo")
	}

	return s.store.Update(id, &libro)
}
func (s *Service) RemoverLibro(id int) error {

	return s.store.Delete(id)
}
