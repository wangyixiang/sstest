package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"] = append(beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"] = append(beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"] = append(beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"] = append(beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"] = append(beego.GlobalControllerRouter["sstest/todoapi/controllers:TodoWebapiTodoitemController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
