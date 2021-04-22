package service

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/util"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func CreateICode(user *model.User, rd *entity.CreateICodeRequestData) (int, error) {
	// 获取有权限的主机列表
	hul, err := model.HostUserList(map[string]interface{}{"user_id": user.Id})
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("未找到任何主机, err:%v", err)
	}

	deployPath := util.Config.Section("agent").Key("deploy_path").String()
	deployShell := util.Config.Section("agent").Key("deploy_shell").String()

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(hul))
	host := hul[index].GetHost()
	port := host.GetUnusedPort()
	cmd := fmt.Sprintf(
		"cd %s &&"+
			"curl -fsSL %s/icode.sh --output icode.sh >> /dev/null &&"+
			"chmod +x icode.sh &&"+
			"./icode.sh %v %v",
		deployPath, deployShell, user.Id, port)
	res, err := host.RunCmd(cmd, 0)

	if err != nil {
		return http.StatusInternalServerError, err
	}
	resArray := strings.Split(res, "\n")

	ic := &model.ICode{
		UserId:        user.Id,
		HostId:        host.Id,
		Name:          *rd.Name,
		Port:          port,
		Password:      resArray[1],
		ContainerId:   resArray[0],
		ContainerPort: 8080,
	}
	if err = model.ICodeAdd(ic); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func GetICodeList(user *model.User) (int, []model.ICode, error) {
	iCodes, err := model.ICodeList(map[string]interface{}{"user_id": user.Id})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, iCodes, nil
}

func DeleteICode(user *model.User, rd *entity.DeleteICodeRequestData) (int, error) {
	iCode, has := model.ICodeOne(map[string]interface{}{"user_id": user.Id, "id": *rd.Id})
	if !has {
		return http.StatusBadRequest, fmt.Errorf("未找到对应的开发机, id:%v", *rd.Id)
	}

	// 删除对应的记录
	containerId := iCode.ContainerId
	if err := model.ICodeDelete(iCode); err != nil {
		return http.StatusInternalServerError, err
	}

	deployPath := util.Config.Section("agent").Key("deploy_path").String()

	// 删除对应的 container
	cmd := fmt.Sprintf("cd %v"+
		"docker rm -f %v &&"+
		"cat i_code_ids | awk '{if ($1!=%v) print $1}' > i_code_ids_tmp &&"+
		"mv i_code_ids_tmp i_code_ids -f &&"+
		"rm -f i_code_ids_tmp",
		deployPath, containerId, containerId,
	)
	_, err := iCode.GetHost().RunCmd(cmd, 0)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
