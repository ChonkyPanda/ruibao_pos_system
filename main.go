package main

import "ruibao_pos_system/database"

import "ruibao_pos_system/controller"
import "github.com/gin-gonic/gin"
// import "ruibao_pos_system/ui"

// import "fyne.io/fyne/v2"
// import "fyne.io/fyne/v2/app"
// import "fyne.io/fyne/v2/widget"


func main() {
	database.InitDB()

	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", controller.GetAllUsers)
		userRoutes.GET("/:id", controller.GetUserByID)
		userRoutes.GET("/username/:username", controller.GetUserByUsername)
		userRoutes.POST("", controller.CreateUser)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
	}

	productRoutes := r.Group("/products")
	{
		productRoutes.GET("", controller.GetAllProducts)
		productRoutes.GET("/:id", controller.GetProductByID)
		productRoutes.POST("", controller.CreateProduct)
		productRoutes.PUT("/:id", controller.UpdateProduct)
		productRoutes.DELETE("/:id", controller.DeleteProduct)
	}

	r.Run(":8080")
	// service.CreateUser("Admin", "Jasonxie175", true)

	// a := app.NewWithID("com.ruibao.pos")
	// w := a.NewWindow("RuiBao POS System")
	// w.Resize(fyne.NewSize(1000, 700))

	// ui.ShowInventory(w, nil)

	// ui.ShowLogin(w, func(user *db.User) {
	// 	// Swap this out for a real dashboard page later, e.g. ui.ShowDashboard(w, user)
	// 	w.SetContent(widget.NewLabel("Welcome, " + user.Name + "!"))
	// })

	// w.ShowAndRun()
}