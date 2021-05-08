# Challenges and Cheats [EMAIL]
Simple project to send an email with a list of random challenges and cheats for the day. Written in Go.

# Important
In order for it to work, you will need to enable less secure apps to run via your gmail address. You can do this by going to your Google Account (the one you wish to send your email from, more specifically the `EMAIL_FROM_ADDRESS`) and enabling it from there. I'm sure there is a guide out there that explains it better than I do.

# Configuration
In order to use it properly you will need to add an environment file (`.env`) with your secrets:

- `EMAIL_FROM_ADDRESS` - your email address (a gmail address to be extra clear)
- `PASSWORD_FROM_ADDRESS` - the password to your gmail address
- `SMTP_HOST` - not typically a secret, most often used is "smtp.gmail.com"
- `SMTP_PORT` - once again, not really a secret, most often used is port 587
- `RECIPIENTS` - a list of the recipients, the email addresses which you intent to send your email to, separated by a "," [comma] i.e. `someone@gmail.com,someone.else@gmail.com`

# Run the application
You can simply run the entire application (all the files at once) by running `go run .`, given that you are in the actual directory with the files. This is simply due to the fact that they all belong to the same "main" package

# Other
Obviously, this is a very basic and simple way to programmatically send an email. I encourage you to experiment with it and come up with your own improvements.
