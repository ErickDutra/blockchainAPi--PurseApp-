package controller

import (
	"go-api/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BlockController struct {
	BlockUsecase usecase.BlockUsecase

}

func NewBlockController(usecase usecase.BlockUsecase) BlockController {
	return BlockController{
		BlockUsecase: usecase,
	}
}


func (bc *BlockController) PostBlockGenes(ctx *gin.Context) {
	var req struct {
		Transactions string `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	block, err := bc.BlockUsecase.NewBlockGenesis(req.Transactions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, block)
}


func (bc *BlockController) GetBlocks(ctx *gin.Context) {
	blocks, err := bc.BlockUsecase.GetLastBlock()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blocks)
}



