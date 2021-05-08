package main

// SMTP server config struct
type SmtpServer struct {
	Host string
	Port string
}

// Email content struct
type EmailContent struct {
	Subject string
	Challenges []string
	Cheats []string
	Quote string
}

// challenges data
var challenges = []string{
	"Programming for at least 2 hours",
	"Study for at least 2 hours",
	"Clean the room",
	"Go to the gym",
	"Ruuuuuuuun",
	"Get up at 4 A.M.",
	"Listening to/reading a book for 2 hours",
	"Talk to one new person, or one that you can learn something from",
	"Make a schedule for the day, stick to it",
}

// cheats data
var cheats = []string{
	"Youtube for 1 hour",
	"Read a good book",
	"Take a nap",
	"One ticket for buying candy",
	"Ticket for watching Netflix",
	"Research a topic that interests you",
	"Hang out with friends",
	"Sit back and relax for 1 hour - do nothing",
	"Look up things that motivate you (materialistic things, non-materialistic)",
	"Monkeytype for 1 hour",
}

// NewSmtpServer initializes our email credentials
func NewSmtpServer(host string, port string) *SmtpServer {
	return &SmtpServer{
		Host: host,
		Port: port,
	}
}

// NewEmailContent uses the subject and quote arguments to return a complete email struct
func NewEmailContent(subject string) *EmailContent {
	challengesSlice := GetRandomPicks(challenges)
	cheatsSlice := GetRandomPicks(cheats)

	return &EmailContent{
		Subject: subject,
		Challenges: challengesSlice,
		Cheats: cheatsSlice,
		Quote: GetRandomQuote(),
	}
}

func (s *SmtpServer) Address() string {
	return s.Host + ":" + s.Port
}