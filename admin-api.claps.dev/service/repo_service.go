package service

import (
	"claps-admin/common"
	"claps-admin/merico"
	"claps-admin/model"
	"claps-admin/util"
	"log"
)

const GithubUrl = "https://github.com/"
const GITHUB = "GITHUB"
const GITLAB = "GITLAB"
const GIT = "GIT"
const BITBUCKET = "BITBUCKET"

// 创建仓库添加到数据库中，返回repoId
func CreateRepository(projectId string, repoType string, repoUrl string, repoName string) (string, *util.Err) {
	DB := common.GetDB()
	var newRepo model.Repository
	repoId, err := merico.AddRepository(repoUrl, projectId)
	if !util.IsOk(err) {
		return "", util.Fail("merico添加仓库失败！", nil)
	}
	newRepo.ProjectId = projectId
	newRepo.Id = repoId
	newRepo.Type = repoType
	newRepo.Slug = repoUrlCut(repoUrl, repoType)
	newRepo.Name = repoName
	if err := DB.Create(&newRepo).Error; err != nil {
		merico.DeleteRepository(repoId)
		return "", util.Fail(err.Error(), err)
	}
	return repoId, util.Success()
}

// 根据项目Id获取所有仓库,返回指针类型
func GetAllReposByProjectId(projectId string) (*[]model.Repository, *util.Err) {
	DB := common.GetDB()
	// 查仓库表获取某个项目的仓库数组
	var repos []model.Repository
	if err := DB.Where("project_id = ?", projectId).Find(&repos).Error; err != nil {
		return nil, util.Fail("数据库查询repo表失败！"+err.Error(), nil)
	}
	if len(repos) == 0 {
		return &repos, util.Fail("获取仓库失败，未查找到任何仓库！", nil)
	}
	for i := 0; i < len(repos); i++ {
		repos[i].Slug = repoUrlFill(repos[i].Slug, repos[i].Type)
	}
	return &repos, util.Success()
}

// 根据项目id删除项目的所有仓库
func DeleteProjectRepos(projectId string) *util.Err {
	DB := common.GetDB()
	var repos []model.Repository
	DB.Where("project_id = ?", projectId).Delete(&repos)
	// 循环删除仓库
	for i := 0; i < len(repos); i++ {
		if err := merico.DeleteRepository(repos[i].Id); !util.IsOk(err) {
			log.Println(err.Message)
			return util.Fail(err.Message, nil)
		}
	}
	return util.Success()
}

/*// 判断一个仓库是否在数据库中存在
func IsRepositoryExist(repo *model.Repository) bool {
	var repoInDB model.Repository
	// 按slug查数据库
	DB.Where("slug = ? AND project_id = ? AND type = ?", repo.Slug, repo.ProjectId, repo.Type).First(&repoInDB)
	if len(repoInDB.Id) != 0 {
		return true
	}
	return false
}*/

// 根据仓库id删除仓库
func DeleteRepository(repoId string) *util.Err {
	DB := common.GetDB()
	var repo model.Repository
	if err := DB.Where("id = ?", repoId).Delete(&repo).Error; err != nil {
		log.Println(err)
		return util.Fail(err.Error(), nil)
	}
	if err := merico.DeleteRepository(repoId); !util.IsOk(err) {
		log.Println(err.Message)
		return util.Fail(err.Message, nil)
	}
	return util.Success()
}

func repoUrlCut(repoUrl string, repoType string) string {
	switch repoType {
	case GITHUB:
		if len(repoUrl) > 19 {
			return repoUrl[19:]
		}
		break
	default:
		break
	}
	return repoUrl
}

func repoUrlFill(repoUrl string, repoType string) string {
	switch repoType {
	case GITHUB:
		return GithubUrl + repoUrl
	default:
		break
	}
	return repoUrl
}
