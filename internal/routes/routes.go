package routes

import (
	"github.com/besyuzkirk/feature-flag-management/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(flagHandler *handlers.FeatureFlagHandler, segmentHandler *handlers.SegmentHandler, rolloutHandler *handlers.RolloutStrategyHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		featureFlags := v1.Group("/feature-flags")
		{
			featureFlags.POST("/", flagHandler.CreateFeatureFlag)
			featureFlags.PUT("/:id", flagHandler.UpdateFeatureFlag)
			featureFlags.DELETE("/:id", flagHandler.DeleteFeatureFlag)
			featureFlags.GET("/:id", flagHandler.GetFeatureFlag)
			featureFlags.GET("/", flagHandler.ListFeatureFlags)
			featureFlags.POST("/rollback/:id", flagHandler.RollbackFeatureFlag)
		}

		segments := v1.Group("/segments")
		{
			segments.POST("/", segmentHandler.CreateSegment)
			segments.PUT("/:id", segmentHandler.UpdateSegment)
			segments.DELETE("/:id", segmentHandler.DeleteSegment)
			segments.GET("/:id", segmentHandler.GetSegment)
			segments.GET("/", segmentHandler.ListSegments)
		}

		rollout := v1.Group("/rollout-strategies")
		{
			rollout.POST("/", rolloutHandler.CreateRolloutStrategy)
			rollout.PUT("/:id", rolloutHandler.UpdateRolloutStrategy)
			rollout.DELETE("/:id", rolloutHandler.DeleteRolloutStrategy)
			rollout.GET("/:id", rolloutHandler.GetRolloutStrategy)
			rollout.GET("/flag/:flag_id", rolloutHandler.ListRolloutStrategiesByFlag)
			rollout.GET("/flag/:flag_id/progress", rolloutHandler.TrackRolloutProgress)
		}
	}

	return router
}
