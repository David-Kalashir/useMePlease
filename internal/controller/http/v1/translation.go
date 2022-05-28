package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/David-Kalashir/crs-server/internal/entity"
	"github.com/David-Kalashir/crs-server/internal/usecase"
	"github.com/David-Kalashir/crs-server/pkg/logger"
)

type loginRoutes struct {
	t usecase.Rigestery
	l logger.Interface
}

func newLoginRoutes(handler *gin.RouterGroup, t usecase.Rigestery, l logger.Interface) {
	r := &loginRoutes{t, l}

	h := handler.Group("/request")
	{
		h.POST("/login", r.login)
		//h.POST("/do-register", r.dorigester)
	}
}

type loginResponse struct {
	Login []entity.Login `json:"login"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} loginResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *loginRoutes) login(c *gin.Context) {
	translations, err := r.t.Login(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - Login")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, loginResponse{translations})
}

/*type doLoginRequest struct {
	Name  string `json:"Name"       example:"amir"`
	Email string `json:"Email"  example:"example@example.cpm"`
}*/
