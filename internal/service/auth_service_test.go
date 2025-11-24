package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"elestial/config"
	"elestial/internal/apperror"
	"elestial/internal/service"
	"elestial/model"
)

// mockUserRepo
type mockUserRepo struct {
	GetUserByNameFn func(ctx context.Context, name string) (model.User, error)
	CreateUserFn    func(ctx context.Context, req model.RegisterRequest) error
}

func (m *mockUserRepo) GetUserByName(ctx context.Context, name string) (model.User, error) {
	return m.GetUserByNameFn(ctx, name)
}

func (m *mockUserRepo) CreateUser(ctx context.Context, req model.RegisterRequest) error {
	return m.CreateUserFn(ctx, req)
}

func (r *mockUserRepo) UpdateUser(ctx context.Context, user model.User) error {
	return nil
}

func (r *mockUserRepo) DeleteUser(ctx context.Context, user model.User) error {

	return nil
}

func (r *mockUserRepo) GetUserById(ctx context.Context, userID int) (model.User, error) {
	var user model.User
	return user, nil
}

// mockAuthRepo
type mockAuthRepo struct {
	SaveRefreshTokenFn func(ctx context.Context, userID int, token string, expires time.Time) error
}

func (m *mockAuthRepo) SaveRefreshToken(ctx context.Context, userID int, token string, expires time.Time) error {
	return m.SaveRefreshTokenFn(ctx, userID, token, expires)
}

func (r *mockAuthRepo) RevokeRefreshToken(ctx context.Context, tokenID int) error {
	return nil
}

func (r *mockAuthRepo) GetRefreshToken(ctx context.Context, token string) (model.RefreshToken, error) {
	var rt model.RefreshToken
	return rt, nil
}

//tests

func TestAuthService_Register(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		userRepo   *mockUserRepo
		input      model.RegisterRequest
		wantErr    bool
		errMessage error
	}{
		{
			name: "OK - new user",
			userRepo: &mockUserRepo{
				GetUserByNameFn: func(_ context.Context, _ string) (model.User, error) {
					return model.User{}, errors.New("not found")
				},
				CreateUserFn: func(_ context.Context, _ model.RegisterRequest) error {
					return nil
				},
			},
			input:   model.RegisterRequest{Name: "Davidson", Email: "elnur@rr.dd", Password: "123qwe!@#QWE", RepeatPassword: "123qwe!@#QWE"},
			wantErr: false,
		},
		{
			name: "Username exists",
			userRepo: &mockUserRepo{
				GetUserByNameFn: func(_ context.Context, _ string) (model.User, error) {
					return model.User{ID: 1}, nil
				},
				CreateUserFn: func(_ context.Context, _ model.RegisterRequest) error {
					return nil
				},
			},
			input:      model.RegisterRequest{Name: "Davidson", Email: "elnur@rr.dd", Password: "123qwe!@#QWE", RepeatPassword: "123qwe!@#QWE"},
			wantErr:    true,
			errMessage: apperror.ErrUserNameExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authRepo := &mockAuthRepo{}
			cfg := &config.Config{}
			svc := service.NewAuth(authRepo, tt.userRepo, cfg)

			err := svc.Register(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if tt.wantErr && err != nil && !errors.Is(err, tt.errMessage) {
				t.Fatalf("expected error message %q, got %v", tt.errMessage, err)
			}
		})
	}
}
