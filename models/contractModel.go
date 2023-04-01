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

func (ur *ContractRepository) CreateContract(contract *Contract) error {
	return ur.DB.Create(contract).Error
}

func (ur *ContractRepository) GetContractById(id string) (*Contract, error) {
	contract := &Contract{}
	if err := ur.DB.First(contract, id).Error; err != nil {
		return nil, err
	}
	return contract, nil
}

func (ur *ContractRepository) GetContracts() ([]*Contract, error) {
	contracts := []*Contract{}
	var count int64
	ur.DB.Model(&Contract{}).Count(&count)
	if err := ur.DB.Limit(int(count)).Find(&contracts).Error; err != nil {
		return nil, err
	}
	return contracts, nil
}

func (ur *ContractRepository) UpdateContract(contract *Contract) error {
	return ur.DB.Save(contract).Error
}

func (ur *ContractRepository) DeleteContract(contract *Contract) error {
	return ur.DB.Delete(contract).Error
}

func NewContractRepository(db *gorm.DB) *ContractRepository {
	return &ContractRepository{
		DB: db,
	}
}