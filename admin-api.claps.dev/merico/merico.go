package merico

import (
	"claps-admin/util"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

const MERICO = "https://cloud.merico.cn"
const CONTENTTYPE = "application/json"
const NonceStr = "ibuaiVcKdpRxkhJ9A"

//签名类
type SignUtils struct {
	params map[string]interface{}
	key    string
}

//实例化签名
func NewSign() (sign *SignUtils) {
	sign = new(SignUtils)
	sign.key = viper.GetString("merico.key")
	sign.params = make(map[string]interface{})
	sign.params["appid"] = viper.GetString("merico.appId")
	return sign
}

func (s *SignUtils) SetNonceStr(nonceStr string) {
	s.params["nonce_str"] = nonceStr
}

/*
对map的key排序，返回有序的slice
*/
func (s *SignUtils) sortMapByKey() (keys []string) {
	for k := range s.params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

/*
设置k-v键值对
*/
func (s *SignUtils) Set(key, value string) {
	s.params[key] = value
}

/*
设置数组和对象k-v
*/
func (s *SignUtils) SetObjectOrArray(key string, value interface{}) (err error) {
	s.params[key] = value
	return
}

/*生成sign*/
func (s *SignUtils) sign() (result string) {
	//对key排序
	keys := s.sortMapByKey()
	//拼接
	for _, val := range keys {
		//对于array和object需要先转义再拼接
		v, err := s.params[val].(string)
		//断言失败ok为false
		if !err {
			//序列化
			if b, err2 := json.Marshal(s.params[val]); err2 == nil {
				v = string(b)
			} else {
				fmt.Println(err2)
				return
			}
		}
		result += val + "=" + v + "&"
	}
	result += "key=" + viper.GetString("merico.key")
	//md5加密
	result = fmt.Sprintf("%x", md5.Sum([]byte(result)))
	//转化大写
	result = strings.ToUpper(result)
	return
}

/*
获取要post的数据
*/
func (s *SignUtils) GetPostData() (result *strings.Reader, err error) {
	//获取sign值
	s.params["sign"] = s.sign()
	//序列化 两次序列化导致转义
	b, err := json.Marshal(s.params)
	if err != nil {
		return
	}
	result = strings.NewReader(string(b))
	return
}

// 根据组名获取id
func GetGroupIdByName(groupName string) (string, *util.Err) {
	signTool := NewSign()
	signTool.SetNonceStr(NonceStr)
	res, err := signTool.GetPostData()
	if err != nil {
		log.Panicln(err)
	}

	url := MERICO + "/openapi/openapi/group/list"
	resp, err := http.Post(url, CONTENTTYPE, res)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// 从resp.Body中提取信息
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	var f interface{}
	if err := json.Unmarshal(bytes, &f); err != nil {
		log.Panicln(err)
	}
	responseMap := f.(map[string]interface{})
	// 判断code是否请求成功
	if int(responseMap["code"].(float64)) != 200 {
		log.Panicln(f)
	}
	// 获取resp.Body.data中的map数组
	groupMaps := responseMap["data"].([]interface{})
	for i := 0; i < len(groupMaps); i++ {
		var gm = make(map[string]interface{})
		gm = groupMaps[i].(map[string]interface{})
		if gm["name"] == groupName {
			return gm["id"].(string), util.Success()
		}
	}
	log.Println("Group name \"" + groupName + "\" doesn't exit!")
	return "", util.Fail("", nil)
}

// 为一个父组添加子组，返回子组的id
func AddGroup(groupName string, description string, parentGroupId string) (string, *util.Err) {
	signTool := NewSign()
	signTool.SetNonceStr(NonceStr)
	signTool.Set("name", groupName)
	signTool.Set("description", description)
	signTool.Set("parentGroupId", parentGroupId)

	res, err := signTool.GetPostData()
	if err != nil {
		log.Panicln(err)
	}

	url := MERICO + "/openapi/openapi/group/add"
	resp, err := http.Post(url, CONTENTTYPE, res)
	//fmt.Println("Ready to POST: ",*res)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	// 解析body
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	var f interface{}
	if err := json.Unmarshal(bytes, &f); err != nil {
		log.Panicln(err)
	}
	responseMap := f.(map[string]interface{})
	// 判断code是否请求成功
	if int(responseMap["code"].(float64)) != 200 {
		log.Panicln(f)
	}
	// 于resp.Body.data中的map，获取返回来的id
	projectId := responseMap["data"].(map[string]interface{})["id"]
	return projectId.(string), util.Success()
}

// 为一个组添加仓库，返回仓库的id
func AddRepository(repoUrl string, projectId string) (string, *util.Err) {
	signTool := NewSign()
	signTool.SetNonceStr(NonceStr)
	signTool.Set("gitUrl", repoUrl)
	signTool.Set("groupId", projectId)

	res, err := signTool.GetPostData()
	if err != nil {
		log.Panicln(err)
	}
	url := MERICO + "/openapi/openapi/project/add"
	resp, err := http.Post(url, CONTENTTYPE, res)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	// 解析body
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	var f interface{}
	if err := json.Unmarshal(bytes, &f); err != nil {
		log.Panicln(err)
	}
	responseMap := f.(map[string]interface{})
	// 判断code是否请求成功
	if int(responseMap["code"].(float64)) != 200 {
		log.Panicln(f)
	}

	// 获取返回来的id
	repoId := responseMap["data"].(map[string]interface{})["project"].(map[string]interface{})["id"]
	return repoId.(string), util.Success()
}

// 删除组
func DeleteGroup(id string) *util.Err {
	signTool := NewSign()
	signTool.SetNonceStr(NonceStr)
	signTool.Set("id", id)

	res, err := signTool.GetPostData()
	if err != nil {
		log.Panicln(err)
	}
	url := MERICO + "/openapi/openapi/group/delete"
	resp, err := http.Post(url, CONTENTTYPE, res)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	// 解析body
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	var f interface{}
	if err := json.Unmarshal(bytes, &f); err != nil {
		log.Panicln(err)
	}
	responseMap := f.(map[string]interface{})
	// 判断code是否请求成功
	if int(responseMap["code"].(float64)) != 200 {
		log.Panicln(f)
	}
	log.Println("merico: Group删除成功！")
	return util.Success()
}

// 获取一个仓库
func GetRepository(repoId string) {
	signTool := NewSign()
	signTool.SetNonceStr(NonceStr)
	//id或者gitUrl
	signTool.Set("id", repoId)

	res, err := signTool.GetPostData()
	if err != nil {
		fmt.Println(err)
		return
	}

	url := MERICO + "/openapi/openapi/project/get"
	fmt.Println(url)

	fmt.Println(*res)
	resp, err := http.Post(url, CONTENTTYPE, res)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(string(b))
	fmt.Println(resp.Status)
}

// 删除仓库
func DeleteRepository(repoId string) *util.Err {
	signTool := NewSign()
	signTool.SetNonceStr(NonceStr)
	signTool.Set("id", repoId)

	res, err := signTool.GetPostData()
	if err != nil {
		log.Panicln(err)
	}
	url := MERICO + "/openapi/openapi/project/delete"
	resp, err := http.Post(url, CONTENTTYPE, res)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	// 解析body
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	var f interface{}
	if err := json.Unmarshal(bytes, &f); err != nil {
		log.Panicln(err)
	}
	responseMap := f.(map[string]interface{})
	// 判断code是否请求成功
	if int(responseMap["code"].(float64)) != 200 {
		log.Panicln(f)
	}
	log.Println("merico仓库删除成功！")
	return util.Success()
}
