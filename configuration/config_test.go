package configuration

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigDataTestSuite struct {
	suite.Suite
	configLoader ConfigLoader
	configData   *ConfigData
}

func TestConfigDataTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigDataTestSuite))
}
func (suite *ConfigDataTestSuite) SetupSuite() {
	suite.configLoader = NewConfigLoader()
	var err error
	suite.configData, err = suite.configLoader.LoadConfig("./config.json")
	if err != nil {
		return
	}
	suite.Nil(err)
}

func (suite *ConfigDataTestSuite) TestShouldLoadConfigFromFile() {

	suite.Equal(15, suite.configData.RandomCharacterLength)
	suite.Equal("http://localhost:8080", suite.configData.ServiceDomain)
}
