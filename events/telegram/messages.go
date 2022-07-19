package telegram

const msgHelp = `I can save and keep your links. Also I can offer you them to read.

In order to save the link, just send me it (including the protocol).

In order to get a random link from your list, send me command /random.
Caution! After that, this page will be removed from your list!`

const msgHello = "Hi there! 🤖\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command 🤔"
	msgNoSavedPages   = "You have no saved pages 🤷‍"
	msgSaved          = "Saved! 👌"
	msgAlreadyExists  = "You have already have this page in your list 😉"
)
