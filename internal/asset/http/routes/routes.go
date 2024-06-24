package routes

import (
	"github.com/BIC-Final-Project/backend/configs/env"
	"github.com/BIC-Final-Project/backend/internal/asset/di"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAssetRoutes(app *fiber.App, db *mongo.Database, env env.EnvVars) {
	asset := app.Group("api/v1/manage-aset")

	// KELOLA ASET
	asetHandler := di.InitAset(db, env)

	aset := asset.Group("/aset")

	aset.Get("", asetHandler.GetAllAset)
	aset.Get("/:aset_id", asetHandler.GetAset)
	aset.Post("", asetHandler.CreateAset)
	aset.Put("/:aset_id", asetHandler.UpdateAset)
	aset.Delete("/:aset_id", asetHandler.DeleteAset)

	// VENDOR
	vendorHandler := di.InitVendor(db, env)

	Vendor := asset.Group("/vendor")

	Vendor.Get("", vendorHandler.GetAllVendor)
	Vendor.Get("/:vendor_id", vendorHandler.GetVendor)
	Vendor.Post("", vendorHandler.CreateVendor)
	Vendor.Put("/:vendor_id", vendorHandler.UpdateVendor)
	Vendor.Delete("/:vendor_id", vendorHandler.DeleteVendor)

	// PERENCANAAN ASET
	perencanaanHandler := di.InitPerencanaan(db, env)

	Rencana := asset.Group("/rencana")

	Rencana.Get("", perencanaanHandler.GetAllRencana)
	Rencana.Get("/:id", perencanaanHandler.GetRencana)
	Rencana.Post("", perencanaanHandler.CreatePerencanaan)
	Rencana.Put("/:id", perencanaanHandler.UpdatePerencanaan)
	Rencana.Delete("/:id", perencanaanHandler.DeleteRencana)

	// PEMELIHARAAN ASET
	pemeliharaanHandler := di.InitPemeliharaan(db, env)

	Pelihara := asset.Group("/pelihara")

	Pelihara.Get("", pemeliharaanHandler.GetAllPelihara)
	Pelihara.Get("/:id", pemeliharaanHandler.GetPelihara)
	Pelihara.Post("", pemeliharaanHandler.CreatePelihara)
	Pelihara.Put("/:id", pemeliharaanHandler.UpdatePelihara)
	Pelihara.Delete("/:id", pemeliharaanHandler.DeletePelihara)
}

// func SetupVendorRoutes(app *fiber.App, db *mongo.Database, env env.EnvVars) {
// 	vendor := app.Group("api/v1/vendor")

// 	// VENDOR
// 	vendorHandler := di.InitVendor(db, env)

// 	Vendor := vendor.Group("/vendor")

// 	Vendor.Get("", vendorHandler.GetAllVendor)
// 	Vendor.Get("/:vendor_id", vendorHandler.GetVendor)
// 	Vendor.Post("", vendorHandler.CreateVendor)
// 	Vendor.Put("/:vendor_id", vendorHandler.UpdateVendor)
// 	Vendor.Delete("/:vendor_id", vendorHandler.DeleteVendor)
// }

// func SetupRencanaRoutes(app *fiber.App, db *mongo.Database, env env.EnvVars) {
// 	rencana := app.Group("apa/v1/rencana")

// 	// RENCANA
// 	perencanaanHandler := di.InitPerencanaan(db, env)

// 	Rencana := rencana.Group("/rencana")

// 	Rencana.Get("", perencanaanHandler.GetAllRencana)
// 	Rencana.Get("/:id", perencanaanHandler.GetRencana)
// 	Rencana.Post("", perencanaanHandler.CreatePerencanaan)
// 	Rencana.Put("/:id", perencanaanHandler.UpdatePerencanaan)
// 	Rencana.Delete("/:id", perencanaanHandler.DeleteRencana)
// }
