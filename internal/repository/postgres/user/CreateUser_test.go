package userpostgresql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
)

func TestRepository_CreateUser(t *testing.T) {
	tests := []struct {
		name          string
		user          *types.User
		mockSetup     func(mock sqlmock.Sqlmock)
		expectError   bool
		errorContains string
	}{
		{
			name: "successful user creation",
			user: &types.User{
				Login:    "testuser",
				FullName: "Test User",
				Gender:   "male",
				Age:      25,
				Phone:    "+1234567890",
				Email:    "test@example.com",
				Avatar:   "https://example.com/avatar.jpg",
				IsActive: true,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					INSERT INTO users \(login, full_name, gender, age, phone, email, avatar, is_active\)
					VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)
				`).WithArgs(
					"testuser",
					"Test User",
					"male",
					25,
					"+1234567890",
					"test@example.com",
					"https://example.com/avatar.jpg",
					true,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectError: false,
		},
		{
			name: "database error",
			user: &types.User{
				Login:    "erroruser",
				FullName: "Error User",
				Gender:   "female",
				Age:      30,
				Phone:    "+0987654321",
				Email:    "error@example.com",
				Avatar:   "https://example.com/error.jpg",
				IsActive: false,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					INSERT INTO users \(login, full_name, gender, age, phone, email, avatar, is_active\)
					VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)
				`).WithArgs(
					"erroruser",
					"Error User",
					"female",
					30,
					"+0987654321",
					"error@example.com",
					"https://example.com/error.jpg",
					false,
				).WillReturnError(sqlmock.ErrCancelled)
			},
			expectError:   true,
			errorContains: "failed to create user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создание мока базы данных
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			// Настройка ожиданий
			tt.mockSetup(mock)

			// Создание репозитория
			repo := NewRepository(db)

			// Выполнение теста
			ctx := context.Background()
			err = repo.CreateUser(ctx, tt.user)

			// Проверка результатов
			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			} else {
				assert.NoError(t, err)
			}

			// Проверка, что все ожидания выполнены
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_CreateUser_WithNilUser(t *testing.T) {
	// Тест с nil пользователем (если нужно)
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.Background()
	err = repo.CreateUser(ctx, nil)

	// В зависимости от реализации, может быть panic или ошибка
	// Обычно лучше добавить валидацию в метод
	if err != nil {
		assert.Contains(t, err.Error(), "user cannot be nil")
	}

	// Проверка, что мок не ожидает никаких вызовов
	assert.NoError(t, mock.ExpectationsWereMet())
}
