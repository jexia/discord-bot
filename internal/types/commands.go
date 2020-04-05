package types

// TODO: Add comments
type Command struct {
	Name        string
	Usage       string
	Description string
	Category    string
	NeedArgs    bool
	Args        map[string]bool
	OwnerOnly   bool
	Enabled     bool
	Run         func(m Message, parameters []string) (DiscordAPIPayload, error)
}

var (
	// TODO: Add comments
	Commands map[string]Command
)

// TODO: Add comments
func (c Command) Register() {
	if Commands == nil {
		Commands = make(map[string]Command)
	}

	if c.Enabled {
		Commands[c.Name] = c
	}
}