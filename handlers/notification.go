package handlers

import (
	"net/http"

	"piscord-backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type NotificationHandler struct {
	NotificationService *services.NotificationService
	AuthService         *services.AuthService
}

func NewNotificationHandler(notificationService *services.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		NotificationService: notificationService,
	}
}

func (h *NotificationHandler) GetUnreadNotificationCount(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userObjectID, err := bson.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	count, err := h.NotificationService.GetUnreadNotificationCount(userObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch unread notification count"})
		return
	}

	c.JSON(http.StatusOK, count)
}

func (h *NotificationHandler) GetMyNotifications(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userObjectID, err := bson.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// page := 1
	// limit := 50
	// if p := c.Query("page"); p != "" {
	// 	if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
	// 		page = parsed
	// 	}
	// }
	// if l := c.Query("limit"); l != "" {
	// 	if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
	// 		limit = parsed
	// 	}
	// }

	notifications, err := h.NotificationService.GetNotificationsByUserID(userObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  notifications,
		"total": len(notifications),
	})
}

func (h *NotificationHandler) MarkNotificationAsRead(c *gin.Context) {
	notificationID, _ := c.Params.Get("id")
	notificationObjectID, err := bson.ObjectIDFromHex(notificationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userObjectID, err := bson.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.NotificationService.MarkNotificationAsRead(notificationObjectID, userObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

func (h *NotificationHandler) MarkAllNotificationsAsRead(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userObjectID, err := bson.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.NotificationService.MarkAllNotificationsAsRead(userObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark all notifications as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	notificationID, _ := c.Params.Get("id")
	notificationObjectID, err := bson.ObjectIDFromHex(notificationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userObjectID, err := bson.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.NotificationService.DeleteNotification(notificationObjectID, userObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

func (h *NotificationHandler) DeleteAllNotifications(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userObjectID, err := bson.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.NotificationService.DeleteAllNotifications(userObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}
