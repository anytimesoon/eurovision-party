package conf

// import (
// 	"strings"

// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/inngest/inngest/pkg/cli"
// )

// var (
// 	// questionName renders a text input to ask for the function name.
// 	questionName = question{
// 		answered: func(m *initModel) bool {
// 			return m.name != ""
// 		},
// 		render: func(m *initModel) string {
// 			if m.name != "" {
// 				return "Function name: " + cli.BoldStyle.Render(m.name) + "\n"
// 			}

// 			b := &strings.Builder{}
// 			b.WriteString(cli.BoldStyle.Render("Function name:") + "\n")
// 			b.WriteString(m.textinput.View())
// 			return b.String()
// 		},
// 		update: func(m *initModel, msg tea.Msg) (tea.Model, tea.Cmd) {
// 			var cmd tea.Cmd
// 			m.textinput.Placeholder = "What should this function be called?"
// 			m.textinput, cmd = m.textinput.Update(msg)
// 			value := m.textinput.Value()

// 			if key, ok := msg.(tea.KeyMsg); ok && key.Type == tea.KeyEnter && value != "" {
// 				m.name = value
// 				m.textinput.Placeholder = eventPlaceholder
// 				m.textinput.SetValue("")
// 			}
// 			return m, cmd
// 		},
// 		next: func(m *initModel) InitQuestion {
// 			if m.template != "" {
// 				return nil
// 			}

// 			return questionTrigger
// 		},
// 	}

// 	// questionTrigger asks the user to select between a scheduled and event trigger.
// 	questionTrigger = question{
// 		answered: func(m *initModel) bool {
// 			return m.triggerType != ""
// 		},
// 		render: func(m *initModel) string {
// 			if m.triggerType != "" {
// 				return "Function trigger: " + cli.BoldStyle.Render(m.triggerType) + "\n"
// 			}

// 			b := &strings.Builder{}
// 			b.WriteString(cli.BoldStyle.Render("How should the function run?") + "\n\n")
// 			b.WriteString(m.triggerList.View())
// 			return b.String()
// 		},
// 		update: updateTrigger,
// 		next: func(m *initModel) InitQuestion {
// 			switch m.triggerType {
// 			case triggerTypeEvent:
// 				return questionEventName
// 			case triggerTypeScheduled:
// 				return questionSchedule
// 			default:
// 				return nil
// 			}
// 		},
// 	}

// 	// questionEventName asks for the event trigger name, rendering an event browser.
// 	questionEventName = question{
// 		answered: func(m *initModel) bool {
// 			return m.event != ""
// 		},
// 		render: renderEvent,
// 		update: updateEvent,
// 		next: func(m *initModel) InitQuestion {
// 			return questionRuntime
// 		},
// 	}

// 	// questionSchedule asks for the event trigger name, rendering an event browser.
// 	questionSchedule = question{
// 		answered: func(m *initModel) bool {
// 			return m.cron != ""
// 		},
// 		render: renderSchedule,
// 		update: updateSchedule,
// 		next: func(m *initModel) InitQuestion {
// 			return questionRuntime
// 		},
// 	}

// 	// questionRuntime asks for the event trigger name, rendering an event browser.
// 	questionRuntime = question{
// 		answered: func(m *initModel) bool {
// 			return m.runtimeType != ""
// 		},
// 		render: renderRuntime,
// 		update: updateRuntime,
// 		next: func(m *initModel) InitQuestion {
// 			switch m.runtimeType {
// 			case runtimeHTTP:
// 				// Ask for URL
// 				return questionURL
// 			case runtimeDocker:
// 				// Ask for Language
// 				return questionLanguage
// 			default:
// 				return nil
// 			}
// 		},
// 	}

// 	// questionLanguage asks for the event trigger name, rendering an event browser.
// 	questionLanguage = question{
// 		name: "language",
// 		answered: func(m *initModel) bool {
// 			return m.language != ""
// 		},
// 		render: renderLanguage,
// 		update: updateLanguage,
// 		next: func(m *initModel) InitQuestion {
// 			return questionScaffold
// 		},
// 	}

// 	// questionScaffold asks for the scaffold to use
// 	questionScaffold = question{
// 		name: "scaffold",
// 		answered: func(m *initModel) bool {
// 			return m.scaffold != nil
// 		},
// 		render: renderScaffold,
// 		update: updateScaffold,
// 		next: func(m *initModel) InitQuestion {
// 			return nil
// 		},
// 	}

// 	// questionURL asks for the event trigger name, rendering an event browser.
// 	questionURL = question{
// 		answered: func(m *initModel) bool {
// 			return m.url != ""
// 		},
// 		render: renderURL,
// 		update: updateURL,
// 		next: func(m *initModel) InitQuestion {
// 			return nil
// 		},
// 	}
// )

// type InitQuestion interface {
// 	// Answered returns whether this question has an answer.  If so,
// 	// the controlling model should skip to the next question in the
// 	// chain via Next()
// 	Answered(m *initModel) bool

// 	// Render renders the question.
// 	Render(m *initModel) string

// 	// Update renders the question.
// 	Update(m *initModel, msg tea.Msg) (tea.Model, tea.Cmd)

// 	// Next retunrs the next question in the chain
// 	Next(m *initModel) InitQuestion
// }

// // question represents an abstract question which is used for single
// // initialization and zero allocation after init.
// type question struct {
// 	// name is a local identifier, used when debugging.
// 	name     string
// 	answered func(m *initModel) bool
// 	render   func(m *initModel) string
// 	update   func(m *initModel, msg tea.Msg) (tea.Model, tea.Cmd)
// 	next     func(m *initModel) InitQuestion
// }

// func (q question) Answered(m *initModel) bool                            { return q.answered(m) }
// func (q question) Render(m *initModel) string                            { return q.render(m) }
// func (q question) Update(m *initModel, msg tea.Msg) (tea.Model, tea.Cmd) { return q.update(m, msg) }
// func (q question) Next(m *initModel) InitQuestion                        { return q.next(m) }

// // https://github.com/inngest/inngest/blob/a92252cd8c1622692c0f30edcee24e87f3532ec2/pkg/cli/initialize/questions.go
