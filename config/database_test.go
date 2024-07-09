package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConnectToDB_DoesNotReturnNil(t *testing.T) {
	db := ConnectToDB()

	assert.Equal(t, db.Error, nil, "Success Connect to Database")
	assert.NotNil(t, db)
}
