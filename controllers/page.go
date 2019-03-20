package controllers

import "GoMD/models"

/********************

新增单独的页面	诸如 关于我 友情连接

路由 ： /?:name
*********************/
func (this *FrontendController) Page() {
	//文章查看页面
	id := this.GetString(":name")
	article := models.GetOneArticle(id, "name")
	temp := *article
	this.Data["id"] = temp[0].Id
	this.Data["article"] = article
	this.Data["comments"] = models.GetArticleComments(temp[0].Id)
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme + "/page.html"
	if theme == "QuietV1" {
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["Sidebar"] = theme + "/sidebar.html"
		this.LayoutSections["Comment"] = theme + "/comment.html"
	}
}
