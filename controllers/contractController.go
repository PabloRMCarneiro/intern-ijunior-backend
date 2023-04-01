package controllers

import (
	"context"
	"jwt-gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ContractController struct {
	ContractRepository *models.ContractRepository
}

func (cc *ContractController) GetContract(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	contractId := c.Param("id")
	contract, err := cc.ContractRepository.GetContractById(contractId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find contract"})
		return
	}

	c.JSON(http.StatusOK, contract)
}
func (cc *ContractController) GetContracts(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	contracts, err := cc.ContractRepository.GetContracts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get contracts"})
		return
	}

	c.JSON(http.StatusOK, contracts)
}
func (cc *ContractController) CreateContracts(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	var contract models.Contract
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if contract.Client_name == "" || contract.Contract_date == "" || contract.Title == "" || contract.Value == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client Name, Contract Date, Title and Value are required"})
		return		
	}

	if err := cc.ContractRepository.CreateContract(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create contract"})
		return
	}
	
	c.JSON(http.StatusOK, contract)
}
func (cc *ContractController) UpdateContracts(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}
	
	contractId := c.Param("id")
	contract, err := cc.ContractRepository.GetContractById(contractId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find contract"})
		return
	}
	
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if contract.Client_name == "" && contract.Contract_date == "" && contract.Title == "" && contract.Value == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client Name, Contract Date, Title or Value are required"})
		return		
	}

	if err := cc.ContractRepository.UpdateContract(contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update contract"})
		return
	}

	c.JSON(http.StatusOK, contract)
}

func (cc *ContractController) DeleteContracts(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	contractId := c.Param("id")
	contract, err := cc.ContractRepository.GetContractById(contractId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find contract"})
		return
	}

	if err := cc.ContractRepository.DeleteContract(contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete contract"})
		return
	}

	c.JSON(http.StatusOK, "Contract deleted sucessfully")
}
