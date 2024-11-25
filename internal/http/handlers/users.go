package handlers

import (
	"context"
	"task01/internal/models"
	"task01/internal/web/users"
)

type userService interface {
	GetAll() ([]models.User, error)
	Get(id uint) (*models.User, error)
	Create(task *models.User) error
	UpdateByID(id uint, user *models.User) (bool, error)
	DeleteByID(id uint) (bool, error)
}

type userHandlers struct {
	service userService
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewUsersHandler(service userService) *userHandlers {
	return &userHandlers{
		service,
	}
}

func (u userHandlers) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	all, err := u.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := make(users.GetUsers200JSONResponse, len(all))
	for i, user := range all {
		response[i] = users.User{
			Created: &user.CreatedAt,
			Email:   &user.Email,
			Id:      &user.ID,
			Updated: &user.UpdatedAt,
		}
	}
	return response, nil
}

func (u userHandlers) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	user := &models.User{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}
	err := u.service.Create(user)
	if err != nil {
		return nil, err
	}
	return users.PostUsers201JSONResponse{
		Created: &user.CreatedAt,
		Email:   &user.Email,
		Id:      &user.ID,
		Updated: &user.UpdatedAt,
	}, err
}

func (u userHandlers) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	done, err := u.service.DeleteByID(request.Id)
	if err != nil {
		return nil, err
	}
	if done {
		return users.DeleteUsersId200Response{}, nil
	}
	return users.DeleteUsersId404Response{}, nil
}

func (u userHandlers) GetUsersId(_ context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	user, err := u.service.Get(request.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return users.GetUsersId404Response{}, nil
	}
	return users.GetUsersId200JSONResponse{
		Created: &user.CreatedAt,
		Email:   &user.Email,
		Id:      &user.ID,
		Updated: &user.UpdatedAt,
	}, err
}

func (u userHandlers) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	user := models.User{}
	if request.Body.Email != nil {
		user.Email = *request.Body.Email
	}
	if request.Body.Password != nil {
		user.Password = *request.Body.Password
	}
	done, err := u.service.UpdateByID(request.Id, &user)
	if err != nil {
		return nil, err
	}
	if done {
		return users.PatchUsersId200Response{}, nil
	}
	return users.PatchUsersId404Response{}, nil
}
