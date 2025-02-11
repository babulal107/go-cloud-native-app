package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/babulal107/go-cloud-native-app/internal/config"
	"github.com/babulal107/go-cloud-native-app/internal/model"
	"log"
	"time"
)

type UserSvc interface {
	AddUser(ctx context.Context, reqData model.UserRequest) (int, error)
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUser(ctx context.Context, userId int) (*model.User, error)
}

type UserSvcImpl struct {
	db *sql.DB
}

func NewUserSvc(appContainer config.AppContainer) UserSvc {
	return UserSvcImpl{db: appContainer.DB}
}

func (s UserSvcImpl) AddUser(ctx context.Context, reqData model.UserRequest) (int, error) {

	user := model.User{}

	// if db operation takes more than 1 second, then context will cancel request and return error
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx,
		"INSERT INTO users (id, name, email) VALUES ($1, $2, $3) RETURNING id",
		reqData.Id, reqData.Name, reqData.Email,
	).Scan(&user.Id)

	// Check for timeout or cancellation
	if ctx.Err() != nil {
		log.Printf("Context cancelled or timed out: %v", ctx.Err())
		return 0, ctx.Err()
	}

	// Handle query errors separately
	if err != nil {
		log.Printf("Error while inserting user: %v", err)
		return 0, err
	}

	return user.Id, nil
}

func (s UserSvcImpl) GetUsers(ctx context.Context) ([]model.User, error) {

	rows, err := s.db.Query("SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		fmt.Printf("Error while getting users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Printf("Error while scanning users: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s UserSvcImpl) GetUser(ctx context.Context, userId int) (*model.User, error) {

	var user model.User
	row := s.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", userId)
	if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("user not found")
			fmt.Printf("Error while getting user: %v", err)
			return nil, err
		} else {
			fmt.Printf("Error while getting user: %v", err)
			return nil, err
		}
	}

	return &user, nil
}
