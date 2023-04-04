package models

import "gorm.io/gorm"

type Contract struct {
	gorm.Model
	Title         string  `gorm:"not null"`
	Value         float64 `gorm:"not null"`
	Client_name   string  `gorm:"not null"`
	Contract_date string  `gorm:"not null"`
}

func (c *Contract) TableName() string {
	return "contracts"
}

type ContractRepository struct {
	DB *gorm.DB
}

func (cr *ContractRepository) CreateContract(contract *Contract) error {
	return cr.DB.Create(contract).Error
}

func (cr *ContractRepository) GetContractById(id string) (*Contract, error) {
	contract := &Contract{}
	if err := cr.DB.First(contract, id).Error; err != nil {
		return nil, err
	}
	return contract, nil
}

func (cr *ContractRepository) GetContracts() ([]*Contract, error) {
	contracts := []*Contract{}
	var count int64
	cr.DB.Model(&Contract{}).Count(&count)
	if err := cr.DB.Limit(int(count)).Find(&contracts).Error; err != nil {
		return nil, err
	}
	return contracts, nil
}

func (cr *ContractRepository) UpdateContract(contract *Contract) error {
	return cr.DB.Save(contract).Error
}

func (cr *ContractRepository) DeleteContract(contract *Contract) error {
	return cr.DB.Delete(contract).Error
}

func NewContractRepository(db *gorm.DB) *ContractRepository {
	return &ContractRepository{
		DB: db,
	}
}