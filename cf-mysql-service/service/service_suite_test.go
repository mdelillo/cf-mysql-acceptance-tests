package service_test

import (
	"testing"

	"github.com/cloudfoundry-incubator/cf-mysql-acceptance-tests/helpers"
)

func TestService(t *testing.T) {
	helpers.PrepareAndRunTests("Service", t)
}
