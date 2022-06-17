package user_handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"investidea.tech.test/internal/dtos"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/internal/repository"
	"investidea.tech.test/internal/services/database"
	"investidea.tech.test/internal/services/log"
	"investidea.tech.test/pkg"
	"investidea.tech.test/pkg/auth"
	"investidea.tech.test/pkg/config"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
	"time"
)

const (
	walletAddress = "0xFC7C98fF48Aa50D75b77A3CA3E7f528817b88255"
)

var (
	db             *database.DB
	mock           sqlmock.Sqlmock
	logger         = log.NewLogger(config.LogConfig{})
	urlRequest, _  = url.Parse(fmt.Sprintf("localhost:8080/v1/auth?wallet_address=%v", walletAddress))
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

func InitServer() userHandler {
	//logger.Info("Initializing redis...")
	//c, err := cache.New(context.TODO(),
	//	fmt.Sprintf("%v:%v",
	//		viper.GetString("redis.host"),
	//		viper.GetString("redis.port"),
	//	),
	//	logger)
	//if err != nil {
	//	panic(err)
	//}
	return New(logger, repository.New(logger, db, nil))
}

func TestUserHandler_Signup(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	v, f := viper.New(), pflag.NewFlagSet(string(pkg.APIAppName), pflag.ExitOnError)
	config.New(v, f)

	buyer := defaultBuyer()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`role`) VALUES (?,?,?,?)")).
		WithArgs(buyer.CreatedAt, buyer.UpdatedAt, nil, buyer.Role)
	mock.ExpectCommit()

	handler := InitServer()
	defer db.Close()

	d, _ := json.Marshal(buyer)
	c.Request = &defaultRequest
	c.Request.Body = io.NopCloser(bytes.NewReader(d))
	handler.Signup(c)

	obj := auth.Authentication{}
	json.Unmarshal(w.Body.Bytes(), &obj)
	assert.True(t, obj.Success || obj.Error == "Already registered!")
}

func TestUserHandler_SignupFailEmptyBody(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	v, f := viper.New(), pflag.NewFlagSet(string(pkg.APIAppName), pflag.ExitOnError)
	config.New(v, f)
	handler := InitServer()
	defer db.Close()

	handler.Signup(c)

	obj := auth.Authentication{}
	json.Unmarshal(w.Body.Bytes(), &obj)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func defaultBuyer() dtos.UserDTO {
	return dtos.UserDTO{
		User: entities.User{
			Role:     auth.BuyerRole.String(),
			Email:    "example.com@gmail.com",
			Name:     "Trung",
			Password: "12345678",
			Address:  walletAddress,
			Base: entities.Base{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}
}
