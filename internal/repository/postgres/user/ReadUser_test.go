package userpostgresql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
)

func TestRepository_ReadUser(t *testing.T) {
	testUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	testTime := time.Now()

	tests := []struct {
		name          string
		userID        uuid.UUID
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedUser  *types.User
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful user retrieval",
			userID: testUUID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"user_id", "login", "full_name", "gender", "age", "phone", "email", "avatar", "registration_date", "is_active",
				}).AddRow(
					testUUID,
					"testuser",
					"Test User",
					"male",
					25,
					"+1234567890",
					"test@example.com",
					"https://example.com/avatar.jpg",
					testTime,
					true,
				)
				mock.ExpectQuery(`
					SELECT user_id, login, full_name, gender, age, phone, email, avatar, registration_date, is_active 
					FROM users 
					WHERE user_id = \$1
				`).WithArgs(testUUID).WillReturnRows(rows)
			},
			expectedUser: &types.User{
				UserID:           testUUID,
				Login:            "testuser",
				FullName:         "Test User",
				Gender:           "male",
				Age:              25,
				Phone:            "+1234567890",
				Email:            "test@example.com",
				Avatar:           "https://example.com/avatar.jpg",
				RegistrationDate: testTime,
				IsActive:         true,
			},
			expectError: false,
		},
		{
			name:   "user not found",
			userID: testUUID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`
					SELECT user_id, login, full_name, gender, age, phone, email, avatar, registration_date, is_active 
					FROM users 
					WHERE user_id = \$1
				`).WithArgs(testUUID).WillReturnError(sql.ErrNoRows)
			},
			expectedUser:  nil,
			expectError:   true,
			errorContains: "user not found",
		},
		{
			name:   "database error",
			userID: testUUID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`
					SELECT user_id, login, full_name, gender, age, phone, email, avatar, registration_date, is_active 
					FROM users 
					WHERE user_id = \$1
				`).WithArgs(testUUID).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedUser:  nil,
			expectError:   true,
			errorContains: "failed to read user",
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
			user, err := repo.ReadUser(ctx, tt.userID)

			// Проверка результатов
			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.expectedUser, user)
			}

			// Проверка, что все ожидания выполнены
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_ReadUser_ScanError(t *testing.T) {
	// Тест на ошибку сканирования
	testUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Неправильное количество колонок для вызова ошибки сканирования
	rows := sqlmock.NewRows([]string{"user_id", "login"}).AddRow(testUUID, "testuser")
	mock.ExpectQuery(`
		SELECT user_id, login, full_name, gender, age, phone, email, avatar, registration_date, is_active 
		FROM users 
		WHERE user_id = \$1
	`).WithArgs(testUUID).WillReturnRows(rows)

	repo := NewRepository(db)

	ctx := context.Background()
	user, err := repo.ReadUser(ctx, testUUID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read user")
	assert.Nil(t, user)

	assert.NoError(t, mock.ExpectationsWereMet())
}
