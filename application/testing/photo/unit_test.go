package photo

// import (
// 	"testing"

// 	"hacktiv/final-project/application/mocks"
// 	photoUsecase "hacktiv/final-project/application/usecases/photo"

// 	errorDomain "hacktiv/final-project/domain/errors"
// 	photoDomain "hacktiv/final-project/domain/photo"

// 	"github.com/stretchr/testify/suite"
// )

// type UnitTestSuite struct {
// 	suite.Suite
// 	photo     photoUsecase.PhotoTesting
// 	photoMock *mocks.PhotoTesting
// }

// func TestUnitTestSuite(t *testing.T) {
// 	suite.Run(t, &UnitTestSuite{})
// }

// func (uts *UnitTestSuite) SetupTest() {
// 	photoMock := mocks.PhotoTesting{}
// 	photo := photoUsecase.NewTesting(&photoMock)

// 	uts.photo = photo
// 	uts.photoMock = &photoMock
// }

// func (uts *UnitTestSuite) TestGetAll() {
// 	uts.photoMock.On("GetAll", 1, 20).Return([]*photoDomain.PaginationResultPhoto{}, nil)

// 	actual, err := uts.photo.GetAll(1, 20)

// 	uts.GreaterOrEqual(1, len(*actual.Data))
// 	uts.EqualError(err, errorDomain.NotFound)
// }

// func (uts *UnitTestSuite) TestGetAll_Error() {
// 	expectedError := errors.New(errorDomain.NotFound)

// 	uts.photoMock.On("List", mock.Anything).Return([]*photoDomain.Photo{}, expectedError)

// 	actual, err := uts.photo.GetAll(0, 0)

// 	uts.Equal(0, len(*actual.Data))
// 	uts.Equal(expectedError, err)

// }
