package mailer

type Mailer struct {
	smtpHost  string
	smtpPort  int
	fromEmail string
	user      string
	password  string
}

func NewMailer(
	smtpHost string,
	smtpPort int,
	fromEmail string,
	user string,
	password string,
) *Mailer {
	return &Mailer{
		smtpHost:  smtpHost,
		smtpPort:  smtpPort,
		fromEmail: fromEmail,
		user:      user,
		password:  password,
	}
}
