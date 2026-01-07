package handlers


import (
	"net/http"

	"github.com/gin-gonic/gin"
	"miniauth/internal/services"
)



type AuthHandler struct{
	authService services.AuthService
}


func NewAuthHandler(authService services.AuthService) *AuthHandler{
	return &AuthHandler{
		authService: authService,
	}
}

type loginRequest struct{
	Email	 string `json:"email" binding:"required,email"`
	Password string	`json:"password" binding: "required" `
}

type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// func (h *AuthHandler) Login(c *gin.Context){
// 	var req loginRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid request",
// 		})
// 		return
// 	}

// 	ctx := c.Request.Context()

// 	user, err := h.authService.Login(ctx, req.Email, req.Password)
// 	if err != nil{
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated,gin.H{
// 		"id": user.ID,
// 		"email": user.Email,
// 	})
// }


func (h *AuthHandler) Register(c *gin.Context){
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	
	ctx := c.Request.Context()

	user, err := h.authService.Register(ctx, req.Email, req.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Id": user.ID,
		"email": user.Email,
	})
}
