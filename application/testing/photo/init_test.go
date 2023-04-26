package photo

import (
	"testing"

	photoUsecase "hexagonal-fiber/application/usecases/photo"
	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
	mssgConst "hexagonal-fiber/utils/constant/message"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type IntTestSuite struct {
	suite.Suite
	db        *gorm.DB
	photoCase photoUsecase.Service
}

func TestIntTestSuite(t *testing.T) {
	suite.Run(t, &IntTestSuite{})
}

func (its *IntTestSuite) SetupSuite() {
	inDB, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		its.FailNowf("unable to connect to database", err.Error())
	}

	inDB = setupDatabase(its, inDB)

	photoRepo := photoRepository.Repository{DB: inDB}
	photoCase := photoUsecase.Service{PhotoRepository: photoRepo}

	its.db = inDB
	its.photoCase = photoCase
}

func (its *IntTestSuite) BeforeTest(suiteName, testName string) {
	if testName == "TestGetAll_Error" {
		return
	}
	seedTestTable(its, its.db)
}

func (its *IntTestSuite) TearDownSuite() {
	tearDownDatabase(its)
}

func (its *IntTestSuite) TearDownTest() {
	cleanTable(its)
}

func (its *IntTestSuite) TestGetByID() {
	actual, err := its.photoCase.GetByID(1)

	its.Nil(err)
	its.Equal(uint(1), actual.ID)

}

func (its *IntTestSuite) TestGetByID_Error() {
	actual, err := its.photoCase.GetByID(0)

	its.EqualError(err, mssgConst.StatusNotFound)
	its.Equal(uint(0), actual.ID)

}

func (its *IntTestSuite) TestGetAll() {
	actual, err := its.photoCase.GetAll(1, 20)

	its.Nil(err)
	its.Greater(len(*actual.Data), 0)

}

func (its *IntTestSuite) TestGetAll_Error() {
	actual, err := its.photoCase.GetAll(1, 1)

	its.Nil(err)
	its.Equal(0, len(*actual.Data))

}
