package logic

import (
	"errors"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	commonModel "lime/internal/common/model"
	commonSrv "lime/internal/common/service"
	"lime/internal/global"
	"lime/pkg/utils"
	"log/slog"
)

type MenuLogic struct{}

func NewMenuLogic() *MenuLogic {
	return &MenuLogic{}
}

func (lc *MenuLogic) SaveRoleMenus(code string, menu_codes []string) error {
	return service.NewRole().SaveRoleMenus(code, menu_codes)
}

func (lc *MenuLogic) GetRoleMenus(Code string) ([]*model.Menu, error) {
	data, err := service.NewRole().GetRoleMenus(Code)
	if err != nil {
		return nil, err
	}

	list := lc.MakeTree(data)
	return list, nil
}

// 创建数据
func (lc *MenuLogic) Create(r *requests.CreateMenu) error {
	srv := service.NewMenu()

	data, err := srv.GetByField("code", r.Code)
	if err != nil {
		return err
	}

	if data.ID != 0 {
		return errors.New("角色编码已存在")
	}

	info := &model.Menu{
		MenuBase: model.MenuBase{
			Code:        r.Code,
			Name:        r.Name,
			Description: r.Description,
			Frontpath:   r.Frontpath,
			Order:       r.Order,
			Icon:        r.Icon,
			ParentCode:  r.ParentCode,
			Visible:     true,
		},
	}

	return service.NewMenu().Create(info)
}

// 更新数据
func (lc *MenuLogic) Update(r *requests.UpdateMenu) error {

	menuSrv := service.NewMenu()
	// 查询是否存在
	menu, err := menuSrv.GetById(uint(r.Id))
	if err != nil {
		return err
	}

	if menu.ID == 0 {
		return errors.New("菜单不存在")
	}

	data := make(map[string]any)
	data["code"] = r.Code
	data["name"] = r.Name
	data["description"] = r.Description
	data["frontpath"] = r.Frontpath
	data["order"] = r.Order
	data["icon"] = r.Icon
	data["parent_code"] = r.ParentCode
	data["visible"] = true
	data["menu_type"] = r.MenuType
	data["is_link"] = r.IsLink

	// 开启事务
	menuSrv.Begin()
	defer func() {
		if err != nil {
			menuSrv.Rollback()
			slog.Error("更新菜单失败", "error", err)
			return
		}

		menuSrv.Commit()
		slog.Info("更新菜单成功", slog.Any("更新内容", data))
	}()

	err = menuSrv.UpdateByMap(uint(r.Id), data)
	if err != nil {
		return err
	}

	// 如果修改了标识码，则需要更新子菜单和角色菜单的标识码
	if menu.Code != r.Code {
		// 更新子菜单
		err = menuSrv.UpdatedatasByMap(map[string]any{"parent_code": menu.Code}, map[string]any{"parent_code": r.Code})
		if err != nil {
			return err
		}
	}

	// 更新角色菜单
	err = service.NewRoleMenu(menuSrv.GetDB()).UpdatedatasByMap(map[string]any{"menu_code": menu.Code}, map[string]any{"menu_code": r.Code})
	if err != nil {
		return err
	}

	return nil
}

// 根据ID查询数据
func (lc *MenuLogic) GetMenuById(id uint) (*model.Menu, error) {
	return service.NewMenu().GetById(id)
}

func (lc *MenuLogic) GetMenus() ([]*model.Menu, error) {
	menus, err := service.NewMenu().GetAll()
	if err != nil {
		return nil, err
	}

	return menus, nil
}

// 根据条件查询数据
func (lc *MenuLogic) ListMenu(condition *commonModel.PageQuery[*requests.QueryMenu]) (*commonModel.ResPage[model.Menu], error) {
	srv := service.NewMenu()
	res, err := srv.ListMenu(condition)
	if err != nil {
		return nil, err
	}

	// 重新根据所有的父级菜单和子集菜单的标识码进行查询
	codes := make([]string, 0)
	for _, menu := range res.List {
		codes = append(codes, menu.Code)
		if menu.ParentCode != "" {
			codes = append(codes, menu.ParentCode)
		}
	}

	// 将切片去重
	codes = utils.Unique(codes)
	resList, err := srv.GetMenuByCodes(codes)
	if err != nil {
		return nil, err
	}

	res.List = lc.MakeTree(resList)
	return res, nil
}

// 删除数据
func (lc *MenuLogic) Del(id uint) error {
	// 查询是否存在
	menuSrv := service.NewMenu()
	menu, err := menuSrv.GetById(id)
	if err != nil {
		return err
	}

	if menu.ID == 0 {
		return errors.New("菜单不存在")
	}

	// 判断是否有子菜单
	count, err := menuSrv.GetCountByFields(map[string]any{"parent_code": menu.Code})
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("菜单下有子菜单，不允许删除")
	}

	// 判断菜单是否可以被移除
	if menu.IsCanNotRemove == 1 {
		return errors.New("菜单不允许被移除")
	}

	// 开启事务
	menuSrv.Begin()
	defer func() {
		if err != nil {
			menuSrv.Rollback()
			slog.Error("删除菜单失败", "error", err)
			return
		}

		menuSrv.Commit()
		slog.Info("删除菜单成功", slog.Any("菜单名称", menu.Name), slog.String("菜单标识码", menu.Code))
	}()

	err = menuSrv.Delete(id)
	if err != nil {
		return err
	}

	// 删除角色菜单
	err = service.NewRoleMenu(menuSrv.GetDB()).DeleteByFields(map[string]any{"menu_code": menu.Code})
	if err != nil {
		return err
	}

	return nil
}

func (lc *MenuLogic) GetMenuTree(Role_code string) ([]*model.Menu, error) {
	// 获取角色菜单
	if Role_code == "" {
		return nil, errors.New("角色编码不能为空")
	}

	roleMenuSrv := service.NewRoleMenu()
	// 获取所有菜单
	menuList, err := roleMenuSrv.GetDatasByField("role_code", Role_code)
	if err != nil {
		return nil, err
	}

	menuCodes := make([]string, 0)
	for _, menu := range menuList {
		menuCodes = append(menuCodes, menu.MenuCode)
	}

	menuSrv := service.NewMenu()
	condition := commonSrv.Condition{
		Field: "code",
		Value: menuCodes,
		Exp:   "in",
	}
	menus, err := menuSrv.GetDatasByFields([]commonSrv.Condition{condition})
	if err != nil {
		return nil, err
	}

	return lc.MakeTree(menus), nil
}

func (lc *MenuLogic) MakeTree(list []*model.Menu) []*model.Menu {
	fList := findFirst(list)
	for _, menu := range fList {
		data := findChildren(list, menu.Code)
		menu.Children = data
	}

	return fList
}

func findFirst(list []*model.Menu) []*model.Menu {
	// 如果 parentCode 为空，则返回第一个元素
	if len(list) == 0 {
		return nil
	}

	rtnList := make([]*model.Menu, 0)

	for _, menu := range list {
		if menu.ParentCode == "" {
			rtnList = append(rtnList, menu)
		}
	}
	return rtnList
}

func findChildren(menus []*model.Menu, code string) []*model.Menu {
	children := make([]*model.Menu, 0)
	for _, menu := range menus {
		if menu.ParentCode == code {
			menu.Children = findChildren(menus, menu.Code)
			children = append(children, menu)
		}
	}

	return children
}

func (lc *MenuLogic) MenuIsNotExistAdd(menu *model.Menu) {
	menuSrv := service.NewMenu()
	isHave, err := menuSrv.CheckMenuExist(menu)
	if err != nil {
		global.Logger.Sugar().Errorf("检查菜单是否存在失败: %s", err.Error())
		return
	}

	if !isHave {
		err = menuSrv.AddMenu(menu)
		if err != nil {
			global.Logger.Sugar().Errorf("添加菜单失败: %s", err.Error())
		}
	}
}
