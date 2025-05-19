package http

import (
	"dg-server/config"
	"dg-server/internal/middleware"
	v1 "dg-server/internal/transport/http/v1"
	v2 "dg-server/internal/transport/http/v2"
	"fmt"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, onboardingHandler *v1.OnboardingHandler, controlHandler *v1.ControlHandler, evidenceTaskHandler *v1.EvidenceTaskHandler, frameworkHandler *v1.FrameworkHandler, departmentHandler *v1.DepartmentHandler) {
	for _, version := range config.APIVersions {
		apiPrefix := fmt.Sprintf("/api/%s", version)
		vGroup := e.Group(apiPrefix, middleware.DeprecationMiddleware()) // ✅ Adds versioning

		if version == "v1" {
			vGroup.GET("/health", v1.HealthCheckHandler)
			vGroup.GET("/greet", v1.GreetHandler)

			// control endpoints
			vGroup.POST("/controls", controlHandler.CreateControl)
			vGroup.GET("/controls/:id", controlHandler.GetControlById)
			vGroup.GET("/controls/code/:control_code", controlHandler.GetControlByCode)
			vGroup.GET("/controls", controlHandler.ListControls)
			vGroup.PUT("/controls", controlHandler.UpdateControl)
			vGroup.DELETE("/controls/:id", controlHandler.DeleteControl)
			vGroup.GET("/controls-summary", controlHandler.GetStatusSummary)

			// evidence endpoints
			vGroup.POST("/evidence-tasks", evidenceTaskHandler.CreateEvidenceTask)
			vGroup.GET("/evidence-tasks/:id", evidenceTaskHandler.GetEvidenceTaskByID)
			vGroup.GET("/evidence-tasks", evidenceTaskHandler.ListEvidenceTasks)
			vGroup.PUT("/evidence-tasks/:id", evidenceTaskHandler.UpdateEvidenceTask)
			vGroup.DELETE("/evidence-tasks/:id", evidenceTaskHandler.DeleteEvidenceTask)
			vGroup.GET("/evidence-summary", evidenceTaskHandler.GetStatusSummary)
			vGroup.POST("/upload-evidence", v1.UploadEvidence)

			// framework endpoints
			vGroup.POST("/frameworks", frameworkHandler.CreateFramework)
			vGroup.GET("/frameworks", frameworkHandler.ListFrameworks)
			vGroup.GET("/frameworks/:id", frameworkHandler.GetFrameworkByID)
			vGroup.GET("/frameworks/name/:name", frameworkHandler.GetFrameworkByName)
			vGroup.PUT("/frameworks/:id", frameworkHandler.UpdateFramework)
			vGroup.DELETE("/frameworks/:id", frameworkHandler.DeleteFramework)
			vGroup.GET("/frameworks/:id/compliance", frameworkHandler.GetCompliancePercentage)
			vGroup.POST("/frameworks/:framework_id/controls/:control_id", frameworkHandler.AddControlToFramework)
			vGroup.DELETE("/frameworks/:framework_id/controls/:control_id", frameworkHandler.RemoveControlFromFramework)
			vGroup.GET("/frameworks/:id/evidence-percentage", frameworkHandler.GetEvidenceTaskPercentage)

			// department
			vGroup.POST("/departments", departmentHandler.CreateDepartment)

			vGroup.POST("/register", onboardingHandler.OrgRegistration)
		}

		if version == "v2" {
			vGroup.GET("/health", v1.HealthCheckHandler) // Shared from v1
			vGroup.GET("/greet", v2.GreetHandler)        // v2-specific greeting
			vGroup.GET("/info", v2.InfoHandler)          // New v2-only feature

			vGroup.POST("/register", onboardingHandler.OrgRegistration)
		}
	}
}
