package user_test

import (
	"database/sql"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/mauFade/infinity/internal/models"
	"github.com/mauFade/infinity/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo *repositories.UserRepository
	user *models.User
}

func (rs *RepositorySuite) SetupSuite() {
	var (
		err error
	)

	rs.conn, rs.mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 rs.conn,
		PreferSimpleProtocol: true,
	})

	rs.DB, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(rs.T(), err)

	rs.repo = repositories.NewUserRepository(rs.DB)
	assert.IsType(rs.T(), &repositories.UserRepository{}, err)

	rs.user = models.NewUser(
		uuid.New(),
		"Test",
		"test@test.com",
		"5541991944035",
		"test",
		"0001",
		"001",
		"123456-8",
		"dev",
		false,
		nil,
		time.Now(),
		time.Now(),
	)
}
