package controller

import (
	"net/http"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type SystemsController struct {
	systemsUseCase usecases.SystemsUseCase
}

func NewSystemsController(systemsUseCase usecases.SystemsUseCase) *SystemsController {
	return &SystemsController{
		systemsUseCase: systemsUseCase,
	}
}

// NOP handles the NOP system call
// @Description  Perform NOP operation
// @Tags         systems
// @Accept       json
// @Produce      json
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /systems/nop [get]
func (c *SystemsController) NOP(ctx *gin.Context) {
	err := c.systemsUseCase.NOP(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "NOP executed successfully"})
}

// HLT handles the HLT system call
// @Description  Check if system is halted
// @Tags         systems
// @Accept       json
// @Produce      json
// @Success      200   {object}  map[string]bool
// @Failure      500   {object}  map[string]string
// @Router       /systems/hlt [get]
func (c *SystemsController) HLT(ctx *gin.Context) {
	halted, err := c.systemsUseCase.HLT(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"halted": halted})
}

// LABELS handles the LABELS system call
// @Description  Retrieve system labels
// @Tags         systems
// @Accept       json
// @Produce      json
// @Param        label   query     string  true  "Label Filter"
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /systems/labels [post]
func (c *SystemsController) LABELS(ctx *gin.Context) {
	label := ctx.Query("label")
	labels, err := c.systemsUseCase.LABELS(ctx, label)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, labels)
}

func (c *SystemsController) SystemsRoutes(r *gin.RouterGroup) {
	r.GET("/nop", c.NOP)
	r.GET("/hlt", c.HLT)
	r.POST("/labels", c.LABELS)
}
