package controllers

import "beego-admin/services/admin_log_service"

type AdminLogController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AdminLogController) NestPrepare() {
	//fmt.Println("AdminLogController NestPrepare")
}

func (this *AdminLogController)Index()  {

	this.Data["data"] = admin_log_service.GetAllData()

	this.Layout = "public/base.html"
	this.TplName = "admin_log/index.html"
}