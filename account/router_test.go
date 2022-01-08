package account

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"street/db"
	"street/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupRouter(t *testing.T) *gin.Engine {
	r := gin.Default()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	g := r.Group("")
	Routers(g, db.New(client))
	return r
}

func TestPingRoute(t *testing.T) {
	router := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/register", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
