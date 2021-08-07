package service

import (
	"claps-admin/common"
	"claps-admin/merico"
	"claps-admin/model"
	"claps-admin/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const RootGroup = "Claps"

// 创建新项目,返回项目id
func CreateProject(projectName string, description string, avatarUrl string) (string, *util.Err) {
	var newProject model.Project
	// 获取父组id
	parentGroupId, err := merico.GetGroupIdByName(RootGroup)
	if !util.IsOk(err) {
		log.Panicln("获取parentGroupId出错！" + err.Message)
	}
	// 添加子组
	projectId, err := merico.AddGroup(projectName, description, parentGroupId)
	if !util.IsOk(err) {
		log.Panicln("获取添加子组时出错！" + err.Message)
	}

	// 判断项目在数据库中是否存在
	DB := common.GetDB()
	if IsProjectExistByName(projectName) {
		return "", util.Fail("项目添加失败！项目名已存在！", nil)
	}
	newProject.Id = projectId
	newProject.AvatarUrl = avatarUrl
	newProject.Name = projectName
	newProject.Description = description
	if err := DB.Create(&newProject).Error; err != nil {
		log.Panicln(err)
	}
	return projectId, util.Success()
}

// 在数据库中按名字查找项目是否存在
func IsProjectExistByName(name string) bool {
	db := common.GetDB()
	var project model.Project
	db.Where("name = ?", name).First(&project)
	if len(project.Id) != 0 {
		return true
	}
	return false
}

// 在数据库中按id查找项目是否存在
func IsProjectExistById(projectId string) bool {
	db := common.GetDB()
	var project model.Project
	db.Where("id = ?", projectId).First(&project)
	if len(project.Id) != 0 {
		return true
	}
	return false
}

// 按项目名称查找并返回项目id
func GetProjectIdByName(name string) (string, *util.Err) {
	db := common.GetDB()
	var project model.Project
	db.Where("name = ?", name).First(&project)
	if len(project.Id) == 0 {
		log.Panicln("查找不到该项目：", name)
	}
	return project.Id, util.Success()
}

// 按项目id返回项目指针
func GetProjectById(projectId string) (*model.Project, *util.Err) {
	db := common.GetDB()
	var project model.Project
	db.Where("id = ?", projectId).First(&project)
	if len(project.Id) == 0 {
		return &project, util.Fail("查找不到该项目id", nil)
	}
	return &project, util.Success()
}

// 按项目名称获取项目
func GetProjectByName(projectName string) (*model.Project, *util.Err) {
	db := common.GetDB()
	var project model.Project
	db.Where("name = ?", projectName).First(&project)
	if len(project.Id) == 0 {
		return &project, util.Fail("查找不到该项目名称", nil)
	}
	return &project, util.Success()
}

// 按用户id获取项目数组
func GetProjectsByUserId(userId int64) (*[]model.Project, int, *util.Err) {
	db := common.GetDB()
	var projects []model.Project
	var members []model.Member
	db.Where("user_id = ?", userId).Find(&members)
	count := len(members)
	// 由members中的项目id获取项目并添加到项目数组中
	for i := 0; i < count; i++ {
		project, err := GetProjectById(members[i].ProjectId)
		if !util.IsOk(err) {
			fmt.Println("在按用户id获取项目数组的过程中， " + err.Message)
			continue
		}
		projects = append(projects, *project)
	}
	return &projects, len(projects), util.Success()
}

// 从数据库中获取所有项目，返回数组指针和长度
func GetAllProjects() (*[]model.Project, int) {
	db := common.GetDB()
	var projects []model.Project
	var count int
	// 读取长度
	db.Table("project").Count(&count)
	// 读取所有记录
	db.Find(&projects)
	return &projects, count
}

// 获取传送过来的项目id
func GetProjectId(ctx *gin.Context) (string, *util.Err) {
	var pIdMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&pIdMap)
	projectId := pIdMap["projectId"]
	if len(projectId) == 0 {
		log.Panicln("项目id获取失败！")
	}
	return projectId, util.Success()
}

// 根据项目id删除项目
func DeleteProjectById(projectId string) *util.Err {
	// 检查项目是否存在
	db := common.GetDB()
	if !IsProjectExistById(projectId) {
		fmt.Println("删除失败！项目不存在！")
		return util.Fail("删除失败！项目不存在！", nil)
	}
	// 删除
	db.Where("id = ?", projectId).Delete(&model.Project{})
	err := merico.DeleteGroup(projectId)
	if !util.IsOk(err) {
		log.Println(err)
		return err
	}
	return util.Success()
}

// 更新数据库中的项目信息
func UpdateProjectInfo(projectName string, editedName string, avatarUrl string, description string) *util.Err {
	db := common.GetDB()
	//var project model.Project
	//projectId, err := GetProjectIdByName(projectName)
	project, err := GetProjectByName(projectName)
	if !util.IsOk(err) {
		log.Panicln("项目信息更新失败！", err)
	}
	if len(editedName) != 0 {
		db.Model(&project).Update("name", editedName)
	}
	if len(avatarUrl) != 0 {
		if avatarUrl == "defaultUrl" {
			avatarUrl = ""
		}
		db.Model(&project).Update("avatar_url", avatarUrl)
	}
	if len(description) != 0 {
		db.Model(&project).Update("description", description)
	}
	return util.Success()
}
