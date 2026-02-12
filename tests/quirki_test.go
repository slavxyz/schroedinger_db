package tests

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"quirky-store-go/internal/repository"
	"quirky-store-go/internal/service"
)

type FakeChaosNever struct{}

func (f *FakeChaosNever) Hit() bool { return false }

type FakeChaosAlways struct{}

func (f *FakeChaosAlways) Hit() bool { return true }

func setup(t *testing.T) *service.QuirkyStore {

	db, err := sql.Open(
		"mysql",
		"root:123456@tcp(localhost:3306)/quirky_store",
	)
	if err != nil {
		t.Fatal(err)
	}

	db.Exec("DELETE FROM store")

	repo := repository.NewStoreRepository(db)
	return service.NewQuirkyStore(repo, &FakeChaosNever{})
}

func TestPutAndGet(t *testing.T) {

	qs := setup(t)

	err := qs.Put("cat", "meow")
	if err != nil {
		t.Fatal(err)
	}

	v, ok, err := qs.Get("cat")
	if err != nil || !ok {
		t.Fatal("value not found")
	}

	if v != "meow" {
		t.Fatalf("expected meow, got %s", v)
	}
}

func TestUpdate(t *testing.T) {

	qs := setup(t)

	qs.Put("dog", "woof")
	qs.Put("dog", "bark")

	v, _, _ := qs.Get("dog")

	if v != "bark" {
		t.Fatal("update failed")
	}
}

func TestDelete(t *testing.T) {

	qs := setup(t)

	qs.Put("apple", "red")
	qs.Delete("apple")

	_, ok, _ := qs.Get("apple")

	if ok {
		t.Fatal("delete failed")
	}
}

func TestChaosGet(t *testing.T) {

	db, _ := sql.Open(
		"mysql",
		"root:123456@tcp(localhost:3306)/quirky_store",
	)

	db.Exec("DELETE FROM store")

	repo := repository.NewStoreRepository(db)
	qs := service.NewQuirkyStore(repo, &FakeChaosAlways{})

	qs.Put("cat", "meow")
	qs.Put("dog", "woof")

	v, _, _ := qs.Get("cat")

	if v != "meow" && v != "woof" {
		t.Fatal("unexpected value")
	}
}

func TestDump(t *testing.T) {

	qs := setup(t)

	qs.Put("a", "1")
	qs.Put("b", "2")

	rows, _ := qs.Dump()

	if len(rows) != 2 {
		t.Fatal("dump failed")
	}
}
