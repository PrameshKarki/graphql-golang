package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/PrameshKarki/event-management-golang/graph/model"
	services "github.com/PrameshKarki/event-management-golang/graph/services/session"
	userEventService "github.com/PrameshKarki/event-management-golang/graph/services/userEvents"
	"github.com/PrameshKarki/event-management-golang/utils"
	"github.com/sirupsen/logrus"
)

// CreateSession is the resolver for the createSession field.
func (r *mutationResolver) CreateSession(ctx context.Context, eventID string, data *model.SessionInput) (*model.Response, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	allowedRoles := []string{"ADMIN", "OWNER", "CONTRIBUTOR"}
	userRole, _ := userEventService.GetRoleOfUser(userID, eventID)
	hasPermission := utils.Includes(allowedRoles, userRole)

	if !hasPermission {
		return nil, fmt.Errorf("you don't have permission to add members to the event")
	}
	_, err := services.CreateSession(eventID, data)

	if err != nil {
		logrus.Error(err.Error())
		return &model.Response{Success: false, Message: "internal server error"}, nil
	} else {
		return &model.Response{Success: true, Message: "session created successfully"}, nil
	}
}

// UpdateSession is the resolver for the updateSession field.
func (r *mutationResolver) UpdateSession(ctx context.Context, id string, data *model.SessionInput) (*model.Response, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	allowedRoles := []string{"ADMIN", "OWNER", "CONTRIBUTOR"}
	eventID, _ := services.GetEventIDFromSession(id)
	userRole, _ := userEventService.GetRoleOfUser(userID, eventID)
	hasPermission := utils.Includes(allowedRoles, userRole)

	if !hasPermission {
		return nil, fmt.Errorf("you don't have permission update session")
	}
	_, err := services.UpdateSession(id, data)
	if err != nil {
		logrus.Error(err.Error())
		return &model.Response{Success: false, Message: "internal server error"}, nil
	} else {
		return &model.Response{Success: true, Message: "session updated successfully"}, nil
	}
}

// DeleteSession is the resolver for the deleteSession field.
func (r *mutationResolver) DeleteSession(ctx context.Context, id string) (*model.Response, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	allowedRoles := []string{"ADMIN", "OWNER", "CONTRIBUTOR"}
	eventID, _ := services.GetEventIDFromSession(id)
	userRole, _ := userEventService.GetRoleOfUser(userID, eventID)
	hasPermission := utils.Includes(allowedRoles, userRole)

	if !hasPermission {
		return nil, fmt.Errorf("you don't have permission to add members to the event")
	}
	_, err := services.DeleteSession(id)
	if err != nil {
		logrus.Error(err.Error())
		return &model.Response{Success: false, Message: "internal server error"}, nil
	} else {
		return &model.Response{Success: true, Message: "session deleted successfully"}, nil
	}
}

// GetEventSessions is the resolver for the getEventSessions field.
func (r *queryResolver) GetEventSessions(ctx context.Context, id string) ([]*model.Session, error) {
	return services.GetEventSession(id)
}
