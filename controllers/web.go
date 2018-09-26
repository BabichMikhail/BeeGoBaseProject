package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) SetBaseLayouts() {
	c.Layout = "base.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "components/header.tpl"
	c.LayoutSections["Nav"] = "components/nav.tpl"
}

func (c *BaseController) Index() {
	c.SetBaseLayouts()
	c.TplName = "pages/index.tpl"
}
