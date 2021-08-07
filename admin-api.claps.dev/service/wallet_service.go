package service

import (
	"claps-admin/common"
	"claps-admin/model"
	"claps-admin/util"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

const AssetType int = 8

var assets = [AssetType]string{model.BTC, model.BCH, model.ETH, model.EOS, model.XRP, model.XEM, model.DOGE, model.USDT}

// 通过map创建项目钱包，每个币种对应一个项目钱包，因此需要创建8个
func CreateProjectWallets(db *gorm.DB, projectId string, distribution string) *util.Err {

	// 将在项目钱包中创建8条记录
	for j := 0; j < AssetType; j++ {
		wallet, err := createWallet(projectId, distribution, assets[j])
		if !util.IsOk(err) {
			log.Panicln("项目钱包创建失败！")
			return err
		}
		db.Create(&wallet)
	}

	return util.Success()
}

// 创建一个项目钱包
func createWallet(projectId string, distribution string, assetId string) (*model.Wallet, *util.Err) {
	var wallet model.Wallet
	var err *util.Err
	wallet.ProjectId = projectId
	wallet.BotId, err = GetBotId(projectId, distribution)
	if !util.IsOk(err) {
		log.Panicln("BotID 获取失败！")
	}
	wallet.AssetId = assetId
	wallet.SyncedAt = time.Now()
	return &wallet, util.Success()
}

// 根据项目id创建成员钱包,每个成员对应一个钱包，每个币种(8种)对应一条记录
func CreateMemberWallets(db *gorm.DB, projectId string, distribution string) *util.Err {
	members, err := GetAllMembersByProjectId(projectId)
	if !util.IsOk(err) {
		return err
	}
	// n个成员生成n*8条记录
	for i := 0; i < len(*members); i++ {
		for j := 0; j < AssetType; j++ {
			/*if IsMemberWalletExisted(projectId, (*members)[i].Id) {
				continue
			}*/
			memberWallet, err := createMemberWallet((*members)[i].Id, projectId, distribution, assets[j])
			if !util.IsOk(err) {
				log.Panicln("成员钱包创建失败！", err)
			}
			if err := db.Create(memberWallet).Error; err != nil {
				continue
			}
		}
	}
	return util.Success()
}

// 为项目添加一个成员钱包
func CreateMemberWallet(projectId string, userId int64) *util.Err {
	db := common.GetDB()
	distribution := GetDistributionByProjectId(projectId)
	for j := 0; j < AssetType; j++ {
		memberWallet, err := createMemberWallet(userId, projectId, distribution, assets[j])
		if !util.IsOk(err) {
			log.Panicln("成员钱包创建失败！", err)
		}
		db.Create(memberWallet)
	}
	return util.Success()
}

// 创建一个成员钱包
func createMemberWallet(userId int64, projectId string, distribution string, assetId string) (*model.MemberWallet, *util.Err) {
	var memberWallet model.MemberWallet
	var err *util.Err
	memberWallet.UserId = userId
	memberWallet.ProjectId = projectId
	memberWallet.AssetId = assetId
	memberWallet.BotId, err = GetBotId(projectId, distribution)
	if !util.IsOk(err) {
		log.Panicln("BotID 获取失败！")
	}
	return &memberWallet, util.Success()
}

// 根据项目id删除所有项目钱包
func DeleteProjectWallets(projectId string) *util.Err {
	db := common.GetDB()
	var wallets []model.Wallet
	db.Where("project_id = ?", projectId).Delete(&wallets)
	return util.Success()
}

// 根据项目id删除所有成员钱包
func DeleteMemberWallets(projectId string) *util.Err {
	db := common.GetDB()
	var memberWallets []model.MemberWallet
	db.Where("project_id = ?", projectId).Delete(&memberWallets)
	return util.Success()
}

// 根据项目id获取分配方式类型
func GetDistributionByProjectId(projectId string) string {
	db := common.GetDB()
	var bot model.Bot
	err := db.Where("project_id = ?", projectId).First(&bot).Error
	if err != nil {
		log.Panicln(err)
	}
	return bot.Distribution
}
