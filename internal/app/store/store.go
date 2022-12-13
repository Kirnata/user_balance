package store

type Store struct {
	config *Config
	repo   *Repo
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	return nil
}

func (s *Store) Close() {

}

func (s *Store) User() *Repo {
	if s.repo != nil {
		return s.repo
	}
	s.repo = &Repo{
		store: s,
	}

	return s.repo
}
