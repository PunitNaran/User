package user

import (
	"io/ioutil"
	"os/user"
	"testing"

	. "gopkg.in/check.v1"
	"gopkg.in/yaml.v2"
)

func TestUser(t *testing.T) { TestingT(t) }

var _ = Suite(&UserTestSuite{})

type UserTestSuite struct {
	config *user.User
}

func (s *UserTestSuite) SuiteSetup(c *C) {
	// Read in the correct configuration
	yamlFile, err := ioutil.ReadFile("testdata/user.yaml")
	c.Check(err, NotNil)
	userDetails := &user.User{}
	err = yaml.Unmarshal(yamlFile, userDetails)
	c.Check(err, NotNil)
	// Save the correct configuration of a user profile
	s.config = userDetails
}

func (s *UserTestSuite) TestWrongPassword() {

}
