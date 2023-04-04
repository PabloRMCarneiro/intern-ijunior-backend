package controllers

import (
	"context"
	"jwt-gin/api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	ProjectRepository *models.ProjectRepository
}

func (pc *ProjectController) CreateProject(c *gin.Context) {

	_, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user from session"})
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if project.Delivery_date == "" || project.Name == "" || project.Contract_id == 0 || project.User_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery date, Name, Contract id and User id are required"})
		return	
	}
	
	if err := pc.ProjectRepository.CreateProject(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create project"})
		return 
	}

	c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) GetProject(c *gin.Context) {

	_, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user from session"})
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	projectId := c.Param("id")
	project, err := pc.ProjectRepository.GetProjectById(projectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) GetProjects(c *gin.Context) {

	_, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user from session"})
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	projects, err := pc.ProjectRepository.GetProjects()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {

	_, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user from session"})
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	projectID := c.Param("id")
	project, err := pc.ProjectRepository.GetProjectById(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find project"})
		return
	}
	
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if project.Delivery_date == "" && project.Name == "" && project.Contract_id == 0 && project.User_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery date, Name, Contract id or User id are required"})
		return
	}

	if err := pc.ProjectRepository.UpdateProject(project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) UpdateUserInProject(c *gin.Context) {

	_, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user from session"})
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	projectID := c.Param("id")
	userId := c.Param("userId")
	project, err := pc.ProjectRepository.GetProjectById(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find project"})
		return
	}

	if project.User_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User id is required"})
		return
	}

	project.User_id = userId

	if err := pc.ProjectRepository.UpdateProject(project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) DeleteProject(c *gin.Context) {

	_, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user from session"})
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	projectID := c.Param("id")
	project, err := pc.ProjectRepository.GetProjectById(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find project"})
		return
	}

	if err := pc.ProjectRepository.DeleteProject(project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) RemoveUserInProject(c *gin.Context) {

	_, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user from session"})
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	projectID := c.Param("id")
	project, err := pc.ProjectRepository.GetProjectById(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find project"})
		return
	}

	if project.User_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User id is required"})
		return
	}

	project.User_id = ""

	if err := pc.ProjectRepository.UpdateProject(project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, project)
}
