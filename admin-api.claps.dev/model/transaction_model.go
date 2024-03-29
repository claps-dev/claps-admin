package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	Id        string          `json:"id,omitempty" gorm:"type:varchar(50);primary_key;not null;"`
	ProjectId string          `json:"project_id,omitempty" gorm:"type:varchar(50);not null"`
	AssetId   string          `json:"asset_id,omitempty" gorm:"type:varchar(50);not null"`
	Amount    decimal.Decimal `json:"amount,omitempty" gorm:"type:varchar(128);not null"`
	CreatedAt time.Time       `json:"created_at,omitempty" gorm:"type:timestamp with time zone"`
	Sender    string          `json:"sender,omitempty" gorm:"type:varchar(50);default:null"`
	Receiver  string          `json:"receiver,omitempty" gorm:"type:varchar(50);default:null"`
}

type TransactionDto struct {
	Id          string          `json:"id,omitempty"`
	ProjectId   string          `json:"project_id,omitempty"`
	ProjectName string          `json:"project_name,omitempty"`
	AssetId     string          `json:"asset_id,omitempty"`
	Asset       string          `json:"asset,omitempty"`
	Amount      decimal.Decimal `json:"amount,omitempty"`
	CreatedAt   time.Time       `json:"created_at,omitempty"`
	Sender      string          `json:"sender,omitempty"`
	Receiver    string          `json:"receiver,omitempty"`
}
