package services

import (
	"context"
	"piscord-backend/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type NotificationService struct {
	MongoService *MongoService
	RedisService *RedisService
}

func NewNotificationService(mongoService *MongoService, redisService *RedisService) *NotificationService {
	return &NotificationService{
		MongoService: mongoService,
		RedisService: redisService,
	}
}

func (ns *NotificationService) GetNotifications(userID bson.ObjectID, limit int64, offset int64) ([]*models.Notification, error) {
	var notifications []*models.Notification
	opts := options.Find().SetSort(bson.M{"createdAt": -1}).SetLimit(limit).SetSkip(offset)
	cursor, err := ns.MongoService.GetCollection("notifications").Find(context.Background(), bson.M{"userId": userID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (ns *NotificationService) GetUnreadNotificationCount(userObjectID bson.ObjectID) (int64, error) {
	notificationsCollection := ns.MongoService.GetCollection("notifications")
	count, err := notificationsCollection.CountDocuments(context.Background(), bson.M{
		"userId": userObjectID,
		"isRead": false,
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ns *NotificationService) CreateNotification(notification *models.Notification) error {
	return ns.RedisService.Publish("notification", "notification.create", notification)
}

func (ns *NotificationService) GetNotificationsByUserID(userID bson.ObjectID) ([]models.Notification, error) {
	var notifications []models.Notification
	cursor, err := ns.MongoService.GetCollection("notifications").Find(context.Background(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var notification models.Notification
		if err := cursor.Decode(&notification); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func (ns *NotificationService) MarkNotificationAsRead(notificationID, userID bson.ObjectID) error {
	return ns.RedisService.Publish("notification", "notification.read", map[string]bson.ObjectID{
		"notificationId": notificationID,
		"userId":         userID,
	})
}

func (ns *NotificationService) MarkAllNotificationsAsRead(userID bson.ObjectID) error {
	return ns.RedisService.Publish("notification", "notification.read_all", map[string]bson.ObjectID{
		"userId": userID,
	})
}

func (ns *NotificationService) DeleteNotification(notificationID, userID bson.ObjectID) error {
	return ns.RedisService.Publish("notification", "notification.delete", map[string]bson.ObjectID{
		"notificationId": notificationID,
		"userId":         userID,
	})
}

func (ns *NotificationService) DeleteAllNotifications(userID bson.ObjectID) error {
	return ns.RedisService.Publish("notification", "notification.delete_all", map[string]bson.ObjectID{
		"userId": userID,
	})
}
