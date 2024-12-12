//@Date 2024/12/12 11:05
//@Desc

package integration

import (
	"github.com/ego-component/egorm"
	"github.com/gotomicro/ego/server/egin"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HandlerTestSuite struct {
	suite.Suite
	db     *egorm.Component
	server *egin.Component
}

func (s *HandlerTestSuite) SetupSuite() {

}

func TestUserHandler_Edit(t *testing.T) {

}

func TestUserHandler_Login(t *testing.T) {

}

func TestUserHandler_Profile(t *testing.T) {

}

func TestUserHandler_SignUp(t *testing.T) {

}

func TestUserHandler_UpdatePassword(t *testing.T) {

}
