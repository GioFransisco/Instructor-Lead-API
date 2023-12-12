package controller

import (
	"net/http"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
	utils "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
)


type NoteController struct {
	uc             usecase.NoteUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.JwtMiddleware
}

func (n *NoteController) createHandler(ctx *gin.Context) {
	var payload model.Note

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.ResponseError(ctx, utils.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	createdNoteDTO, err := n.uc.Create(payload)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusCreated, createdNoteDTO)
}

func (n *NoteController) updateHandler(ctx *gin.Context) {
	noteID := ctx.Param("id")

	var payload model.Note
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := n.uc.Update(noteID, payload)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, note)
}

func (n *NoteController) listHandler(ctx *gin.Context) {
	notes, err := n.uc.FindAll()
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, notes)
}

func (n *NoteController) Route() {
	noteGroup := n.rg.Group(config.NoteGroupPath, n.authMiddleware.AuthMiddleware("Trainer"))
	noteGroup.GET(config.NoteGetPath, n.listHandler)
	noteGroup.POST(config.NoteCreatePath, n.createHandler)
	noteGroup.PUT(config.NoteUpdatePath, n.updateHandler)
	noteGroup.GET(config.NoteGetByIDPath, n.getByIDHandler)
	noteGroup.DELETE(config.NoteDeletePath, n.deleteHandler)
}

func (n *NoteController) getByIDHandler(ctx *gin.Context) {
	noteID := ctx.Param("id")
	note, err := n.uc.FindByID(noteID)

	if err == utils.NotFoundErrorByID {
		common.ResponseError(ctx, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, note)
}

func (n *NoteController) deleteHandler(ctx *gin.Context) {
	noteID := ctx.Param("id")

	err := n.uc.Delete(noteID)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ResponseSuccess(ctx, "Note deleted successfully", http.StatusNoContent, nil)
}


func NewNoteController(uc usecase.NoteUseCase, rg *gin.RouterGroup, authMidlleware middleware.JwtMiddleware) *NoteController {
	return &NoteController{uc: uc, rg: rg, authMiddleware: authMidlleware}
}
