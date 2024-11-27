package service

import (
    "crypto/sha256"
    "encoding/hex"
    "time"
    "github.com/Tabintel/invoice-system/internal/repository"
)

type ShareService struct {
    invoiceRepo *repository.InvoiceRepository
}

func NewShareService(invoiceRepo *repository.InvoiceRepository) *ShareService {
    return &ShareService{invoiceRepo: invoiceRepo}
}

type ShareableLink struct {
    Token     string    `json:"token"`
    ExpiresAt time.Time `json:"expires_at"`
}

func (s *ShareService) GenerateShareableLink(invoiceID int64, expiryHours int) (*ShareableLink, error) {
    // Generate unique token
    hash := sha256.New()
    hash.Write([]byte(fmt.Sprintf("%d-%d", invoiceID, time.Now().UnixNano())))
    token := hex.EncodeToString(hash.Sum(nil))

    link := &ShareableLink{
        Token:     token,
        ExpiresAt: time.Now().Add(time.Duration(expiryHours) * time.Hour),
    }

    // Store in database
    err := s.invoiceRepo.SaveShareableLink(context.Background(), invoiceID, link)
    if err != nil {
        return nil, err
    }

    return link, nil
}
