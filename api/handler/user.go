package handler

import (
	pb "github.com/dilshodforever/restaurant-auth/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user
// @Summary Create User
// @Description Create page
// @Tags User
// @Accept  json
// @Produce  json
// @Param   Create     body    pb.User     true        "Create"
// @Success 200 {string} string  "Create Successful"
// @Failure 401 {string} string  "Error while Created"
// @Router /user/create [post]
func (h *Handler) RegisterUser(ctx *gin.Context){
		arr:=pb.User{}
		err:=ctx.BindJSON(&arr)
		if err!=nil{
			panic(err)
		}
		_, err=h.User.CreateUser(ctx, &arr)
		if err!=nil{
			panic(err)
		}
		ctx.JSON(200, "Success!!!")
}

// UpdateUser handles the creation of a new user
// @Summary Update User
// @Description Update page
// @Tags User
// @Accept  json
// @Produce  json
// @Param     id path string true "User ID"
// @Param   Update     body    pb.User     true        "Update"
// @Success 200 {string} string  "Update Successful"
// @Failure 401 {string} string  "Error while created"
// @Router /User/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context){
	arr:=pb.User{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	_, err=h.User.UpdateUser(ctx, &arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}


// DeleteUser handles the creation of a new User
// @Summary Delete User
// @Description Delete page
// @Tags User
// @Accept  json
// @Produce  json
// @Param     id path string true "User ID"
// @Success 200 {string} string  "Delete Successful"
// @Failure 401 {string} string  "Error while Deleted"
// @Router /User/delete/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.User.DeleteUser(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllUser handles the creation of a new User
// @Summary GetAll User
// @Description GetAll page
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} pb.GetAllUser   "GetAll Successful"
// @Failure 401 {string} string  "Error while GetAlld"
// @Router /User/getall [get]
func (h *Handler) GetAllUser(ctx *gin.Context){
	arr:=&pb.User{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	res, err:=h.User.GetAllUser(ctx, arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdUser handles the creation of a new User
// @Summary GetById User
// @Description GetById page
// @Tags User
// @Accept  json
// @Produce  json
// @Param     id path string true "User ID"
// @Success 200 {object} pb.User   "GetById Successful"
// @Failure 401 {string} string "Error while GetByIdd"
// @Router /User/getbyid/{id} [get]
func (h *Handler) GetbyIdUser(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.User.GetByIdUser(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdUser handles the creation of a new User
// @Summary GetById User
// @Description GetById page
// @Tags User
// @Accept  json
// @Produce  json
// @Param     id path string true "User ID"
// @Success 200 {object} pb.User   "GetById Successful"
// @Failure 401 {string} string "Error while GetByIdd"
// @Router /User/getbyid/{id} [get]
func (h *Handler) LoginUser(ctx *gin.Context){
	user:=pb.User{}
	err:=ctx.ShouldBindJSON(&user)
	if err!=nil{
		panic(err)
	}
	res, err:=h.User.
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

