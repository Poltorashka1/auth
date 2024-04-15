package smtp

import (
	"auth/internal/closer"
	"auth/internal/config"
	apperrors "auth/internal/errors"
	"auth/internal/logger"
	"crypto/tls"
	"html/template"
	"net/smtp"
	"os"
)

// todo logging

type SMTP interface {
	SendEmail(email string, code string, username string) error
	Close() error
}

type appSMTP struct {
	cfg    config.SMTPConfig
	log    logger.Logger
	client *smtp.Client
}

func New(cfg config.SMTPConfig, log logger.Logger) (SMTP, error) {
	smtpClient, err := Connect(cfg)
	if err != nil {
		return nil, err
	}
	err = Auth(smtpClient, cfg)
	if err != nil {
		return nil, err
	}

	return &appSMTP{
		cfg:    cfg,
		client: smtpClient,
		log:    log,
	}, nil
}

func Auth(smtpClient *smtp.Client, cfg config.SMTPConfig) error {
	auth := smtp.PlainAuth("", cfg.Username(), cfg.Password(), cfg.Host())

	err := smtpClient.Auth(auth)
	if err != nil {
		return err
	}
	err = smtpClient.Mail(cfg.From())
	if err != nil {
		return err
	}
	return nil
}
func (s *appSMTP) Close() error {
	err := s.client.Quit()
	if err != nil {
		return err
	}
	return nil
}

func Connect(cfg config.SMTPConfig) (*smtp.Client, error) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         cfg.Host(),
	}

	conn, err := tls.Dial("tcp", cfg.Address(), tlsConfig)
	if err != nil {

		return nil, err
	}

	closer.Add(conn.Close)

	smtpClient, err := smtp.NewClient(conn, cfg.Host())
	if err != nil {
		return nil, err
	}
	return smtpClient, nil
}

func (s *appSMTP) Address() string {
	return s.cfg.Address()
}

func (s *appSMTP) Start() error {
	return nil
}

type message struct {
	Username          string
	VerificationLink  string
	VerificationToken string
}

// SendEmail send email to user with verification code
func (s *appSMTP) SendEmail(email, token, username string) error {
	const op = "smtp.SendEmail"

	// if err := s.client.Rcpt(email); err != nil {
	//	s.log.ErrorOp(op, err)
	//	return apperrors.ErrSmtpSendMessage
	//}

	msg := message{
		Username:          username,
		VerificationLink:  email,
		VerificationToken: token,
	}
	templ, err := os.ReadFile(s.cfg.Template())
	if err != nil {
		s.log.ErrorOp(op, err)
		return apperrors.ErrSMTPSendMessage
	}

	emailTemplate, err := template.New("email").Parse(string(templ))
	if err != nil {
		s.log.ErrorOp(op, err)
		return apperrors.ErrSMTPSendMessage
	}

	// w, err := s.client.Data()
	// if err != nil {
	//	s.log.ErrorOp(op, err)
	//	return apperrors.ErrSmtpSendMessage
	//}

	// var buff bytes.Buffer
	err = emailTemplate.Execute(os.Stdout, &msg)
	if err != nil {
		s.log.ErrorOp(op, err)
		return apperrors.ErrSMTPSendMessage
	}
	// _, err = w.Write([]byte("Subject: Verify your email\nTo: " + email + "\n\n"))
	// if err != nil {
	//	s.log.ErrorOp(op, err)
	//	return apperrors.ErrSmtpSendMessage
	//}
	//
	// err = w.Close()
	// if err != nil {
	//	s.log.ErrorOp(op, err)
	//	return apperrors.ErrSmtpSendMessage
	//}

	return nil
}
