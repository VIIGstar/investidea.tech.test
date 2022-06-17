package session_handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"investidea.tech.test/internal/dtos"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/internal/repository"
	"investidea.tech.test/internal/services/database"
	"investidea.tech.test/internal/services/log"
	"investidea.tech.test/pkg/auth"
	base_entity "investidea.tech.test/pkg/base-entity"
	"investidea.tech.test/pkg/config"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
)

var (
	db             *database.DB
	mock           sqlmock.Sqlmock
	logger         = log.NewLogger(config.LogConfig{})
	urlRequest, _  = url.Parse("localhost:8080/v1/auth")
	defaultRequest = http.Request{
		Method: http.MethodPost,
		URL:    urlRequest,
		Header: map[string][]string{
			"User-Agent": {"PostmanRuntime/7.29.0"},
		},
		RemoteAddr: "[::1]:60844",
	}
)

func init() {
	var sqlDB *sql.DB
	sqlDB, mock, _ = sqlmock.New()
	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	db = database.TestDB(gormDB)
}

func TestSessionHandler_Login(t *testing.T) {
	dto := dtos.UserDTO{
		User: entities.User{
			Base: base_entity.Base{
				ID: 1,
			},
			Role:     auth.BuyerRole.String(),
			Email:    "example.com",
			Password: "12345678",
		},
	}
	parsedEntity, _ := dto.ToEntity()
	user := parsedEntity.(entities.User)
	bData, _ := json.Marshal(dto)

	mock.MatchExpectationsInOrder(false)
	mock.ExpectQuery(
		regexp.QuoteMeta(
			"SELECT * FROM `users` WHERE email = ? AND password = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
		),
	).
		WithArgs(user.Email, user.Password).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "role"}).
				AddRow(user.ID, auth.BuyerRole.String()))
	handler := New(logger, repository.New(logger, db, nil))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &defaultRequest
	c.Request.Body = io.NopCloser(bytes.NewReader(bData))
	handler.Login(c)

	obj := auth.Authentication{}
	json.Unmarshal(w.Body.Bytes(), &obj)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, obj.Success)
}
