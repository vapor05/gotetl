package options


type CommandOptions interface {
    CommandName() string
    Options() map[string]string
}
