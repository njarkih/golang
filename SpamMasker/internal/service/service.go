package service

// читает файл и предоставляет его контент
type producer interface {
	Produce() ([]string, error)
}

// выводит результат
type presenter interface {
	Present([]string) error
}

// новая сущность сервиса
type Service struct {
	prod producer
	pres presenter
}

// конструктор, возвращает указатель на Service
func NewService(prod producer, pres presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}

// собственно маскируем
func (s *Service) Run() error {
	lines, err := s.prod.Produce() // зачитали
	if err != nil {
		return err
	}

	var result []string
	for _, line := range lines {
		result = append(result, HideLinks(line))
	}

	return s.pres.Present(result) // записали
}
