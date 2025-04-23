package service

import (
	"errors"
	"lime/internal/app/admin/requests"
	"sync"

	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

type Casbin struct {
	db *gorm.DB
}

func NewCasbin(db *gorm.DB) *Casbin {
	return &Casbin{db: db}
}

func (srv *Casbin) UpdateCasbin(roleCode string, casbinInfos []requests.CasbinInfo) error {
	srv.ClearCasbin(0, roleCode)
	rules := [][]string{}
	//做权限去重处理
	deduplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := roleCode + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, []string{roleCode, v.Path, v.Method})
		}
	}
	if len(rules) == 0 {
		return nil
	} // 设置空权限无需调用 AddPolicies 方法
	e := srv.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func (srv *Casbin) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := srv.db.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	if err != nil {
		return err
	}

	e := srv.Casbin()
	err = e.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

func (srv *Casbin) GetPolicyPathByRoleCode(roleCode string) ([]requests.CasbinInfo, error) {
	e := srv.Casbin()
	list, err := e.GetFilteredPolicy(0, roleCode)
	if err != nil {
		return nil, err
	}

	pathMaps := make([]requests.CasbinInfo, len(list))
	for index, v := range list {
		pathMaps[index] = requests.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		}
	}
	return pathMaps, nil
}

func (srv *Casbin) GetRoleApis(roleCode string) ([]requests.CasbinInfo, error) {
	list := make([]gormadapter.CasbinRule, 0)
	err := srv.db.Model(&gormadapter.CasbinRule{}).Where("v0 =?", roleCode).Find(&list).Error
	if err != nil {
		return nil, err
	}

	listMap := make([]requests.CasbinInfo, len(list))
	for index, v := range list {
		listMap[index] = requests.CasbinInfo{
			Path:   v.V1,
			Method: v.V2,
		}
	}

	return listMap, err
}

func (srv *Casbin) ClearCasbin(v int, p ...string) bool {
	e := srv.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

func (srv *Casbin) RemoveFilteredPolicy(roleCode string) error {
	return srv.db.Delete(&gormadapter.CasbinRule{}, "v0 = ?", roleCode).Error
}

func (srv *Casbin) RemoveRoleApi(roleCode string, path string, method string) error {
	return srv.db.Delete(&gormadapter.CasbinRule{}, "v0 =? AND v1 =? AND v2 =?", roleCode, path, method).Error
}

func (srv *Casbin) AddRoleApi(roleCode string, path string, method string) error {
	return srv.db.Create(&gormadapter.CasbinRule{
		Ptype: "p",
		V0:    roleCode,
		V1:    path,
		V2:    method,
	}).Error
}

func (srv *Casbin) SyncPolicy(roleCode string, rules [][]string) error {
	err := srv.RemoveFilteredPolicy(roleCode)
	if err != nil {
		return err
	}
	return srv.AddPolicies(rules)
}

func (srv *Casbin) AddPolicies(rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return srv.db.Create(&casbinRules).Error
}

func (srv *Casbin) FreshCasbin() (err error) {
	e := srv.Casbin()
	err = e.LoadPolicy()
	return err
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (srv *Casbin) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(srv.db)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}

		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}
