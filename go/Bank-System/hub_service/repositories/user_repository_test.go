package repositories

import (
	"database/sql"
	"fmt"
	"hub_service/entities"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var repository *UserRepository
var mock sqlmock.Sqlmock

var db *sql.DB
var err error
var dialector gorm.Dialector
var gormDB *gorm.DB

func TestMain(m *testing.M) {
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer db.Close()

	dialector = postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "mysql",
	})

	gormDB, err = gorm.Open(dialector)

	// open gorm db
	if err != nil {
		log.Fatalf("%v", err)
	}

	repository = NewUserRepository(gormDB)

	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	var data []entities.UserEntity

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "password"}).
			AddRow(uuid.NewString(), "paulo", "keller", "paulo", "test@gmail.com", "123456"))

	data, err = repository.GetAll()

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(data) < 1 {
		t.Errorf("Test failed. Should return at least one user entity.")
	}

	if data[0].Username != "paulo" {
		t.Errorf("Test failed. Wrong username")
	}

	if data[0].Email != "test@gmail.com" {
		t.Errorf("Test failed. Wrong email.")
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(uuid.NewString())).
		WillReturnError(fmt.Errorf("some error"))

	data, err = repository.GetAll()

	if data != nil || err == nil {
		t.Errorf("Test failed. Request query fail.")
	}
}

func TestGetByID(t *testing.T) {
	var data *entities.UserEntity

	id := uuid.NewString()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "password"}).
		AddRow(id, "paulo", "keller", "paulo", "test@gmail.com", "123456"))

	data, err = repository.GetByID(id)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data.Username != "paulo" {
		t.Errorf("Test failed. Wrong username")
	}

	if data.Email != "test@gmail.com" {
		t.Errorf("Test failed. Wrong email.")
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "password"}).
		AddRow(id, "paulo", "keller", "paulo", "test@gmail.com", "123456")).WillReturnError(fmt.Errorf("some error"))

	data, err = repository.GetByID(id)

	if data != nil || err == nil {
		t.Errorf("Test failed. Request query fail.")
	}
}

func TestInsert(t *testing.T) {
	id, err := uuid.NewUUID()

	if err != nil {
		t.Errorf(err.Error())
	}

	entity := &entities.UserEntity{
		ID:        id,
		FirstName: "Paulo",
		LastName:  "Keller",
		Username:  "paulo",
		Email:     "paulo@keller.com",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	t.Skip("Skipping test because it's not ready")

	slqInsert := `INSERT INTO "users" ("id","deleted_at","first_name","last_name","username","email","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id","first_name","last_name","username","email","password","created_at","updated_at"`
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(slqInsert)).
		WithArgs(entity.ID, entity.DeletedAt, entity.FirstName, entity.LastName, entity.Username, entity.Email, entity.Password, entity.CreatedAt, entity.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repository.Insert(entity)

	if err != nil {
		t.Errorf(err.Error())
	}
}
