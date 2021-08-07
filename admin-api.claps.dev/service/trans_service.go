package service

import (
	"claps-admin/common"
	"claps-admin/model"
	"claps-admin/util"
	"log"
)

// 从数据库中获取某个项目的捐赠流水
func GetProjectTransactions(projectId string) (*[]model.Transaction, int) {
	db := common.GetDB()
	var transactions []model.Transaction
	db.Where("project_id = ?", projectId).Find(&transactions)
	count := len(transactions)
	return &transactions, count
}

// 获取所有项目的捐赠流水
func GetAllTransactions() (*[]model.TransactionDto, *util.Err) {
	var tranDtos []model.TransactionDto
	var transactions []model.Transaction
	DB := common.GetDB()
	if err := DB.Find(&transactions).Error; err != nil {
		log.Panicln(err)
	}
	if len(transactions) == 0 {
		log.Println("捐赠流水为空")
		return &tranDtos, util.Success()
	}
	for _, tran := range transactions {
		var tranDto *model.TransactionDto
		tranDto = toTransactionDto(&tran)
		tranDtos = append(tranDtos, *tranDto)
	}
	return &tranDtos, util.Success()
}

// 获取所有用户提现记录
func GetAllTransfers() (*[]model.Transfer, *util.Err) {
	var transfers []model.Transfer
	DB := common.GetDB()
	if err := DB.Find(&transfers).Error; err != nil {
		log.Panicln(err)
	}
	return &transfers, util.Success()
}

// 从数据库中获取某个用户的提现记录
func GetUserTransfers(mixinId string) (*[]model.Transfer, int, *util.Err) {
	db := common.GetDB()
	var transfers []model.Transfer
	db.Where("mixin_id = ?", mixinId).Find(&transfers)
	count := len(transfers)
	if count == 0 {
		log.Panicln("未查找到该用户的提现记录")
	}
	return &transfers, count, util.Success()
}

// 将transaction转化成transactionDto
func toTransactionDto(transaction *model.Transaction) *model.TransactionDto {
	var transactionDto model.TransactionDto
	project, err := GetProjectById(transaction.ProjectId)
	if !util.IsOk(err) {
		log.Panicln(err)
		return nil
	}
	transactionDto.ProjectName = project.Name
	transactionDto.ProjectId = project.Id
	transactionDto.Id = transaction.Id
	transactionDto.ProjectId = transaction.ProjectId
	transactionDto.CreatedAt = transaction.CreatedAt
	transactionDto.AssetId = transaction.AssetId
	transactionDto.Amount = transaction.Amount
	transactionDto.Receiver = transaction.Receiver
	transactionDto.Sender = transaction.Sender

	return &transactionDto
}

/*// 将transfer转化成transferDto
func ToTransferDto(transfer *model.Transfer) (*model.TransferDto, *util.Err) {
	DB := common.GetDB()
	var transferDto model.TransferDto
	user, err := GetUserByMixinId(DB, transfer.MixinId)
	if !util.IsOk(err) {
		return &transferDto, err
	}
	transferDto.UserId = user.Id
	transferDto.AvatarUrl = user.AvatarUrl
	transferDto.Name = user.Name
	transferDto.Email = user.Email
	transferDto.MixinId = user.MixinId
	transferDto.Total = transfer.Amount
	transferDto.Withdraws = 1
	return &transferDto, util.Success()
}*/
