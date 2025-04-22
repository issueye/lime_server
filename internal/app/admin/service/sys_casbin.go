package service

import (
	"errors"
	"lime/internal/app/admin/requests"
	"strconv"
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

func (srv *Casbin) UpdateCasbin(roleID uint, casbinInfos []requests.CasbinInfo) error {
	id := strconv.Itoa(int(roleID))
	srv.ClearCasbin(0, id)
	rules := [][]string{}
	//做权限去重处理
	deduplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := id + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, []string{id, v.Path, v.Method})
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

func (srv *Casbin) GetPolicyPathByRoleID(roleID uint) ([]requests.CasbinInfo, error) {
	e := srv.Casbin()
	id := strconv.Itoa(int(roleID))
	list, err := e.GetFilteredPolicy(0, id)
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

func (srv *Casbin) ClearCasbin(v int, p ...string) bool {
	e := srv.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

func (srv *Casbin) RemoveFilteredPolicy(db *gorm.DB, roleID string) error {
	return db.Delete(&gormadapter.CasbinRule{}, "v0 = ?", roleID).Error
}

func (srv *Casbin) SyncPolicy(db *gorm.DB, roleID string, rules [][]string) error {
	err := srv.RemoveFilteredPolicy(db, roleID)
	if err != nil {
		return err
	}
	return srv.AddPolicies(db, rules)
}

func (srv *Casbin) AddPolicies(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
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
