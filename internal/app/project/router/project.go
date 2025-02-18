package router

import (
	v1 "lime/internal/app/project/controller/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	InitProjectRouter(r)
}

func InitProjectRouter(r *gin.RouterGroup) {
	project := r.Group("project")
	// project.Use(middleware.AuthMiddleware())
	{
		project.GET(":id", v1.GetProject)
		project.POST("list", v1.ProjectList)
		project.POST("", v1.CreateProject)
		project.PUT("", v1.UpdateProject)
		project.DELETE(":id", v1.DeleteProject)
		project.POST("sync/:id", v1.SyncProject)

		// Branch routes
		branch := project.Group("branch")
		{
			branch.POST("", v1.CreateBranch)
			branch.PUT("", v1.UpdateBranch)
			branch.PUT("status", v1.UpdateBranchStatus)
			branch.DELETE(":id", v1.DeleteBranch)
			branch.POST("list", v1.BranchList)
		}

		// Tag routes
		tag := project.Group("tag")
		{
			tag.POST("", v1.CreateTag)
			tag.PUT("", v1.UpdateTag)
			tag.PUT("release", v1.UpdateTagReleaseStatus)
			tag.DELETE(":id", v1.DeleteTag)
			tag.POST("list", v1.TagList)
		}

		// Version routes
		version := project.Group("version")
		{
			version.POST("", v1.CreateVersion)
			version.PUT("", v1.UpdateVersion)
			version.PUT("build", v1.UpdateVersionBuildStatus)
			version.DELETE(":id", v1.DeleteVersion)
			version.POST("list", v1.VersionList)
		}

		// Compile routes
		compile := project.Group("compile")
		{
			compile.GET("run/:id", v1.Compile)
			compile.POST("config", v1.SaveCompileConfig)
			compile.GET("config/:id", v1.GetConfigByProjectId)
		}

		// Package routes
		pkg := project.Group("package")
		{
			pkg.DELETE(":id", v1.DeletePackage)
			pkg.POST("list", v1.PackageList)
			pkg.POST("download/:id", v1.DownloadPackage)
		}
	}
}
