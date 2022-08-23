//go:build integration || pkg || users || all

package tests

import (
	"net/http"
	"testing"

	"github.com/muharamdani/go-base-rest-api/pkg/users/routers"
	"github.com/muharamdani/go-base-rest-api/pkg/users/models"
	"github.com/muharamdani/go-base-rest-api/utils"
	"github.com/muharamdani/go-base-rest-api/db"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserTestSuite struct {
	suite.Suite
	DB *gorm.DB
	Server *gin.Engine
	Client *httpexpect.Expect
}

// Setup for the entire suite, for specific test setup use SetupTest() below
func (suite *UserTestSuite) SetupSuite() {
    db.ConnectSqlite("../../../test.sqlite3")
	suite.DB = db.DB
	
	suite.Server = gin.Default()
	routers.Router(suite.Server)

	suite.Client = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(suite.Server),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(suite.T()),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(suite.T(), true),
		},
	})

	suite.DB.AutoMigrate(&models.Users{})
}

// Teardown for the entire suite, for specific test teardown use TearDownTest() below
func (suite *UserTestSuite) TearDownSuite() {
	db, _ := suite.DB.DB()
	db.Close()

	err := os.Remove("../../../test.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
}

// Executed before each test
func (suite *UserTestSuite) SetupTest() {
	
}

// Executed after each test
func (suite *UserTestSuite) TearDownTest() {
	suite.DB.Exec("DELETE FROM users")
}

func (suite *UserTestSuite) TestGetUsers() {
	suite.Client.GET("/users").
		Expect().
		Status(http.StatusOK).JSON().Object().ValueEqual("message", "Get user data success")
}

func (suite *UserTestSuite) TestInsertUser() {
	resp := suite.Client.POST("/users").
		WithJSON(map[string]interface{}{
			"first_name":   "Muhamad",
			"last_name":    "Ramdani",
			"username":     "muharamdani",
			"phone_number": "812317641",
			"email":        "muharamdani@gmail.com",
			"address":      "Rumah no 91",
		}).Expect().Status(http.StatusOK).JSON().Object()

	resp.Value("data").Object().ValueEqual("first_name", "Muhamad")
}

func TestUserTestSuite(t *testing.T) {
    suite.Run(t, new(UserTestSuite))
}