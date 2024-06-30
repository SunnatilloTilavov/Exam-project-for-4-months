package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"errors"

	"api_gateway/api/models"
	"api_gateway/config"
	"api_gateway/pkg/grpc_client"
	"api_gateway/pkg/logger"
	"api_gateway/pkg/jwt"
)

type handler struct {
	log        logger.Logger
	grpcClient *grpc_client.GrpcClient
	cfg        config.Config
}

// HandlerV1Config ...
type HandlerConfig struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

const (
	// ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	// ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	// ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	// ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	// ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	// ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	// ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	// ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
	// ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	// ErrorCodeNotApproved ...
	ErrorCodeNotApproved = "NOT_APPROVED"
	// ErrorCodeWrongClub ...
	ErrorCodeWrongClub = "WRONG_CLUB"
	// ErrorCodePasswordsNotEqual ...
	ErrorCodePasswordsNotEqual = "PASSWORDS_NOT_EQUAL"
)

var (
	SigningKey = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)7Ddo")
)

// New ...
func New(c *HandlerConfig) *handler {
	return &handler{
		log:        c.Logger,
		grpcClient: c.GrpcClient,
		cfg:        c.Cfg,
	}
}

func handleGrpcErrWithDescription(c *gin.Context, l logger.Logger, err error, message string) bool {
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ErrorWithDescription{
			Code:        http.StatusBadRequest,
			Description: st.Message(),
		})
		l.Error(message, logger.Error(err))
		return true
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusNotFound, models.ErrorWithDescription{
			Code:        http.StatusNotFound,
			Description: st.Message(),
		})
		l.Error(message+", not found", logger.Error(err))
		return true
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ErrorWithDescription{
			Code:        http.StatusInternalServerError,
			Description: "Internal Server Error",
		})
		l.Error(message+", service unavailable", logger.Error(err))
		return true
	} else if st.Code() == codes.AlreadyExists {
		c.JSON(http.StatusInternalServerError, models.ErrorWithDescription{
			Code:        http.StatusInternalServerError,
			Description: st.Message(),
		})
		l.Error(message+", already exists", logger.Error(err))
		return true
	} else if st.Code() == codes.InvalidArgument {
		c.JSON(http.StatusBadRequest, models.ErrorWithDescription{
			Code:        http.StatusBadRequest,
			Description: st.Message(),
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	} else if st.Code() == codes.Code(20) {
		c.JSON(http.StatusBadRequest, models.ErrorWithDescription{
			Code:        http.StatusBadRequest,
			Description: st.Message(),
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	} else if st.Err() != nil {
		c.JSON(http.StatusBadRequest, models.ErrorWithDescription{
			Code:        http.StatusBadRequest,
			Description: st.Message(),
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	}
	return false
}


func getAuthInfo(c *gin.Context) (AuthInfo, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return AuthInfo{}, errors.New("unauthorized")
	}

	m, err := jwt.ExtractClaims(accessToken)
	if err != nil {
		return AuthInfo{}, err
	}

	role := m["user_role"].(string)
	if !(role =="SuperAdmin" || role == "Teacher" || role == "SupportTeacher" || role == "Student"|| role == "Manager"|| role == "Administration") {
		return AuthInfo{}, errors.New("unauthorized")
	}

	return AuthInfo{
		UserID:   m["user_id"].(string),
		UserRole: role,
	}, nil
}

type AuthInfo struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}