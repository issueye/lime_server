package home

import (
	"context"
	_ "embed"
	"fmt"
	"lime/internal/common/config"
	"lime/internal/global"
	"time"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//go:embed logo.png
var logo []byte

// ::private::
type TFrmHomeFields struct {
	IsServerRun  bool
	ShowLogCount int32
	IsTrueClose  bool
	ActiveItem   int64
	ctx          context.Context
	cancel       context.CancelFunc

	serverRunCtx    context.Context
	serverRunCancel context.CancelFunc
}

func (f *TFrmHome) OnFormCreate(sender vcl.IObject) {
	f.IsServerRun = false

	f.SetCaption(global.APP_NAME)
	f.TForm.SetPosition(types.PoScreenCenter)
	f.ShowHint()
	f.Tray_icon.SetVisible(true)
	f.Tray_icon.SetHint(global.APP_NAME)
	f.StatusBar.Panels().Items(0).SetWidth(100)
	f.ctx, f.cancel = context.WithCancel(context.Background())

	f.InitData()
	f.SetEvents()

	f.Monitor()
}

func (f *TFrmHome) SetEvents() {
	f.TForm.SetOnClose(f.OnFormClose)
	f.TForm.SetOnDestroy(f.OnFormDestroy)

	f.Tray_icon.SetOnClick(f.OnTrayIconClick)

	f.PM_close.SetOnClick(f.OnAppCloseClick)
	f.PM_show.SetOnClick(f.OnAppShowClick)
	f.PM_open_web_page.SetOnClick(f.OnOpenWebPageClick)
}

func (f *TFrmHome) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	f.Hide()
	if !f.IsTrueClose {
		*action = types.CaHide
	}
}

func (f *TFrmHome) OnFormDestroy(sender vcl.IObject) {
	f.cancel()
}

func (f *TFrmHome) OnAppCloseClick(sender vcl.IObject) {
	f.IsTrueClose = true
	f.Close()
}

func (f *TFrmHome) OnAppShowClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmHome) OnOpenWebPageClick(sender vcl.IObject) {
	port := config.GetParam(config.SERVER, "port", config.DEF_PORT).Int()
	url := fmt.Sprintf("http://127.0.0.1:%d/web", port)
	openBrowser(url)
}

func (f *TFrmHome) OnTrayIconClick(sender vcl.IObject) {
	f.Show()
}

func GetRunDate(start time.Time) int {
	diff := time.Since(start)
	return int(diff.Hours() / 24)
}
