package http

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestContext struct {
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
}

func NewTestContext() *TestContext {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	c.Request = httptest.NewRequest("GET", "/", nil)

	return &TestContext{
		Recorder: rec,
		Context:  c,
	}
}

func (c *TestContext) SetParam(key, val string) *TestContext {
	c.Context.Params = append(c.Context.Params, gin.Param{Key: key, Value: val})
	return c
}

func (c *TestContext) AssertStatus(t *testing.T, code int) bool {
	return assert.Equal(t, code, c.Recorder.Code)
}

func (c *TestContext) AssertJSONBodyEquals(t *testing.T, obj interface{}) bool {
	out := reflect.Indirect(reflect.ValueOf(obj)).Addr().Interface()
	err := json.NewDecoder(c.Recorder.Body).Decode(out)

	if err != nil {
		return assert.Failf(t, "Failed to decode body: %s, %s", c.Recorder.Body.String(), err)
	}

	return assert.Equal(t, obj, out)

}
