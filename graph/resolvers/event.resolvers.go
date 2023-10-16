package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/PrameshKarki/event-management-golang/graph"
	"github.com/PrameshKarki/event-management-golang/graph/model"
	eventService "github.com/PrameshKarki/event-management-golang/graph/services/event"
	userEventService "github.com/PrameshKarki/event-management-golang/graph/services/userEvents"
	"github.com/PrameshKarki/event-management-golang/utils"
)

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, data model.EventInput) (*model.EventResponse, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	id, err := eventService.CreateEvent(data, userID)
	// Assign user to admin role
	userEventService.CreateUserEvent(fmt.Sprint(id), userID, "OWNER")
	return &model.EventResponse{ID: &id}, err
}

// AddMembersToEvent is the resolver for the addMembersToEvent field.
func (r *mutationResolver) AddMembersToEvent(ctx context.Context, id string, data model.AddMemberInput) (string, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	allowedRoles := []string{"ADMIN", "OWNER", "CONTRIBUTOR"}
	userRole, _ := userEventService.GetRoleOfUser(userID, id)
	// Check if the user is admin or owner of the event, Only owner and admin can add members to the event
	hasPermission := utils.Includes(allowedRoles, userRole)

	if !hasPermission {
		return "", fmt.Errorf("you don't have permission to add members to the event")
	}

	isContributor := userRole == "CONTRIBUTOR"

	// If the currently logged in user is contributor, then he/she can only add attendees
	if isContributor {
		for _, member := range data.Members {
			if member.Role != "ATTENDEE" {
				return "", fmt.Errorf("you have permission to add only attendees to the event")
			}
		}
	}

	_, err := userEventService.AddMembersToEvent(id, data)
	if err != nil {
		return "", err
	} else {
		return "Success", nil
	}
}

// RemoveMemberFromEvent is the resolver for the removeMemberFromEvent field.
func (r *mutationResolver) RemoveMemberFromEvent(ctx context.Context, id string, memberID string) (*model.Response, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	allowedRoles := []string{"ADMIN", "OWNER"}
	userRole, _ := userEventService.GetRoleOfUser(userID, id)
	// Check if the user is admin or owner of the event, Only owner and admin can add members to the event
	hasPermission := utils.Includes(allowedRoles, userRole)

	if !hasPermission {
		return nil, fmt.Errorf("you don't have permission to remove members from the event")
	}
	_, err := userEventService.RemoveUserFromEvent(id, memberID)
	if err != nil {
		return &model.Response{Success: false, Message: "Internal Server Error"}, err
	} else {
		return &model.Response{Success: true, Message: "Member Removed Successfully"}, nil
	}
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (*model.Response, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	userRole, _ := userEventService.GetRoleOfUser(userID, id)
	// Check if the user is admin or owner of the event, Only owner and admin can add members to the event
	hasPermission := userRole == "OWNER"

	if !hasPermission {
		return nil, fmt.Errorf("you don't have permission remove the event")
	}
	_, err := eventService.DeleteEvent(id)
	if err != nil {
		return &model.Response{Success: false, Message: "Internal Server Error"}, err
	} else {
		return &model.Response{Success: true, Message: "Event Deleted Successfully"}, nil
	}
}

// UpdateEvent is the resolver for the updateEvent field.
func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, data model.EventInput) (*model.EventResponse, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	userRole, _ := userEventService.GetRoleOfUser(userID, id)
	// Check if the user is admin or owner of the event, Only owner and admin can add members to the event
	hasPermission := userRole == "OWNER"

	if !hasPermission {
		return nil, fmt.Errorf("you don't have permission update the event")
	}
	eventId, err := eventService.UpdateEvent(id, data)
	return &model.EventResponse{ID: &eventId}, err
}

// UpdateSchedule is the resolver for the updateSchedule field.
func (r *mutationResolver) UpdateSchedule(ctx context.Context, id string, data model.ScheduleUpdateInput) (*model.Response, error) {
	userID := ctx.Value("user").(*utils.TokenMetadata).ID
	allowedRoles := []string{"ADMIN", "OWNER"}
	userRole, _ := userEventService.GetRoleOfUser(userID, id)
	// Check if the user is admin or owner of the event, Only owner and admin can update schedule of the event
	hasPermission := utils.Includes(allowedRoles, userRole)
	if !hasPermission {
		return nil, fmt.Errorf("you don't have permission update the schedule")
	}
	_, err := eventService.UpdateEventSchedule(id, data)
	if err != nil {
		return &model.Response{Success: false, Message: "Internal Server Error"}, err
	} else {
		return &model.Response{Success: true, Message: "Schedule Updated Successfully"}, nil
	}
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	return eventService.GetEvents()
}

// Event is the resolver for the event field.
func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	return eventService.GetEvent(id)
}

// GetMembersOfEvent is the resolver for the getMembersOfEvent field.
func (r *queryResolver) GetMembersOfEvent(ctx context.Context, id string) ([]*model.Member, error) {
	return userEventService.GetMembersOfEvent(id)
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
