package database

import (
	"context"
	"falcon/domain/entity"
	"falcon/domain/model"
)

func (r *repository) GetUserByID(ctx context.Context, userID int64) (*entity.User, model.AppError) {
	var (
		user     *entity.User
		appError model.AppError
		tx       = r.db.WithContext(ctx)
	)

	if err := tx.Where("user_id = ?", userID).Find(&user).Error; err != nil {
		r.logger.Error(err)
		appError.Message = err.Error()
		appError.StatusCode = model.Database
		return nil, appError
	}

	return user, appError
}

func (r *repository) GetUserByName(ctx context.Context, username string) (*entity.User, model.AppError) {
	var (
		user     *entity.User
		appError model.AppError
		tx       = r.db.WithContext(ctx)
	)

	if err := tx.Where("username = ?", username).Find(&user).Error; err != nil {
		r.logger.Error(err)
		appError.Message = err.Error()
		appError.StatusCode = model.Database
		return nil, appError
	}

	return user, appError
}

func (r *repository) GetAllUsers(ctx context.Context) ([]entity.User, model.AppError) {
	var (
		users    []entity.User
		appError model.AppError
		tx       = r.db.WithContext(ctx)
	)

	if err := tx.Find(&users).Error; err != nil {
		r.logger.Error(err)
		appError.Message = err.Error()
		appError.StatusCode = model.Database
		return nil, appError
	}

	return users, appError
}

func (r *repository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, model.AppError) {
	var (
		appError model.AppError
		tx       = r.db.WithContext(ctx)
	)

	if err := tx.Create(&user).Error; err != nil {
		r.logger.Error(err)
		appError.Message = err.Error()
		appError.StatusCode = model.Database
		return nil, appError
	}

	return user, appError
}

func (r *repository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, model.AppError) {
	var (
		appError model.AppError
		tx       = r.db.WithContext(ctx)
	)

	if err := tx.Updates(&user).Error; err != nil {
		r.logger.Error(err)
		appError.Message = err.Error()
		appError.StatusCode = model.Database
		return nil, appError
	}

	return user, appError
}
