package registerHandler

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var config = util.ErrorConfig{
	Options: []util.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "Fullname",
			Message: "fullname is required on body",
		},
		{
			Tag:     "lowercase",
			Field:   "Fullname",
			Message: "fullname must be using lowercase",
		},
		{
			Tag:     "required",
			Field:   "Email",
			Message: "email is required on body",
		},
		{
			Tag:     "email",
			Field:   "Email",
			Message: "email format is not valid",
		},
		{
			Tag:     "required",
			Field:   "Password",
			Message: "password is required on body",
		},
		{
			Tag:     "gte",
			Field:   "Password",
			Message: "password minimum must be 8 character",
		},
	},
}

func ErrRegisterHandler(resultRegister *models.EntityUsers, ctx *gin.Context, errRegister string) {
	switch errRegister {

	case "REGISTER_CONFLICT_409":
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "REGISTER_FAILED_403":
		util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
		accessToken, errToken := util.Sign(accessTokenData, util.GodotEnv("JWT_SECRET"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		_, errSendMail := util.SendGridMail(resultRegister.Fullname, resultRegister.Email, "Activation Account", "template_register", accessToken)

		if errSendMail != nil {
			defer logrus.Error(errSendMail.Error())
			util.APIResponse(ctx, "Sending email activation failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}