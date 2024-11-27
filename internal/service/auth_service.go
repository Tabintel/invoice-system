package service

import (
    "context"
    "errors"
    "os"
    "time"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "github.com/Tabintel/invoice-system/internal/repository"
)

type AuthService struct {
    repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
    return &AuthService{repo: repo}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
    user, err := s.repo.GetByEmail(ctx, email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return "", err
    }

    return tokenString, nil
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
