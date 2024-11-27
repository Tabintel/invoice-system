package repository

import (
    "context"
    "database/sql"
    "time"
)

type User struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name"`
    Email       string    `json:"email"`
    Phone       string    `json:"phone"`
    CompanyName string    `json:"company_name"`
    CompanyLogo string    `json:"company_logo"`
    Location    string    `json:"location"`
    Role        string    `json:"role"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *User) error {
    query := `
        INSERT INTO users (name, email, phone, company_name, company_logo, location, role)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, created_at, updated_at`

    return r.db.QueryRowContext(ctx, query,
        user.Name, user.Email, user.Phone, user.CompanyName,
        user.CompanyLogo, user.Location, user.Role,
    ).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
    user := &User{}
    query := `SELECT * FROM users WHERE email = $1`

    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.ID, &user.Name, &user.Email, &user.Phone,
        &user.CompanyName, &user.CompanyLogo, &user.Location,
        &user.Role, &user.CreatedAt, &user.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }
    return user, nil
}
