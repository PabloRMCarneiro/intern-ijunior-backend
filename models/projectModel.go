package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Delivery_date string `gorm:"not null"`
	Name         string `gorm:"not nul"`
	Contract_id  uint   `gorm:"not null"`
	User_id      string `gorm:"not null"`
}

func (c *Project) TableName() string {
	return "projects"
}

type ProjectRepository struct {
	DB *gorm.DB
}

func (pr *ProjectRepository) CreateProject(project *Project) error {
	return pr.DB.Create(project).Error
}

func (pr *ProjectRepository) GetProjectById(id string) (*Project, error) {
	project := &Project{}
	if err := pr.DB.First(project, id).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (pr *ProjectRepository) GetProjects() ([]*Project, error) {
	projects := []*Project{}
	var count int64
	pr.DB.Model(&Project{}).Count(&count)
	if err := pr.DB.Limit(int(count)).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (pr *ProjectRepository) UpdateProject(project *Project) error {
	return pr.DB.Save(project).Error
}

func (pr *ProjectRepository) DeleteProject(project *Project) error{
	return pr.DB.Delete(project).Error
}

func (pr *ProjectRepository) RemoveUserInProject(project *Project, userId string) error {
	project.User_id = ""
	return pr.DB.Save(project).Error
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}