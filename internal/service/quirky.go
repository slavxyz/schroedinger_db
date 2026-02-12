package service

import (
	"quirky-store-go/internal/chaos"
	"quirky-store-go/internal/repository"
)

type QuirkyStore struct {
	repo  *repository.StoreRepository
	chaos chaos.Chaos
}

func NewQuirkyStore(
	repo *repository.StoreRepository,
	chaos chaos.Chaos,
) *QuirkyStore {
	return &QuirkyStore{
		repo:  repo,
		chaos: chaos,
	}
}

func (s *QuirkyStore) Put(key, value string) error {

	if s.chaos.Hit() {

		// silent fail
		if randBool() {
			return nil
		}

		// wrong key
		if k, ok, _ := s.repo.RandomKey(); ok {
			key = k
		}
	}

	_, exists, err := s.repo.SelectValue(key)
	if err != nil {
		return err
	}

	if !exists {
		return s.repo.Insert(key, value)
	}

	return s.repo.Update(key, value)
}

func (s *QuirkyStore) Get(key string) (string, bool, error) {

	if s.chaos.Hit() {
		if k, ok, _ := s.repo.RandomKey(); ok {
			key = k
		}
	}

	_ = s.repo.Mutate(key)

	return s.repo.SelectValue(key)
}

func (s *QuirkyStore) Delete(key string) error {

	if s.chaos.Hit() {
		if k, ok, _ := s.repo.RandomKey(); ok {
			key = k
		}
	}

	return s.repo.Delete(key)
}

func (s *QuirkyStore) Dump() ([]repository.Row, error) {
	return s.repo.Dump()
}
