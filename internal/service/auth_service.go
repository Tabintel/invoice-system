package service

import (
    "context"
    "errors"
    "os"
    "time"
    
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "github.com/Tabintel/invoice-system/internal/repository"
)

type AuthService struct {
    repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
    return &AuthService{repo: repo}
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (string, error) {
    user, err := s.repo.GetByEmail(ctx, req.Email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    claims := &middleware.Claims{
        UserID: user.ID,
        Role:   user.Role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
            IssuedAt:  time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return "", errors.New("failed to generate token")
    }

    return signedToken, nil
}

type RegisterRequest struct {
    Name        string `json:"name"`
    Email       string `json:"email"`
    Password    string `json:"password"`
    CompanyName string `json:"company_name"`
    Phone       string `json:"phone"`
    Location    string `json:"location"`
}

func (s *AuthService) Register(ctx context.Context, req *RegisterRequest) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &repository.User{
        Name:        req.Name,
        Email:       req.Email,
        Password:    string(hashedPassword),
        CompanyName: req.CompanyName,
        Phone:       req.Phone,
        Location:    req.Location,
        Role:        "user",
    }

    return s.repo.Create(ctx, user)
}
