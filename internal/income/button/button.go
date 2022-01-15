package button

import "github.com/sepuka/vkbotserver/api/button"

const (
	StartIdButton    = `start`
	WithdrawIdButton = `withdraw`

	TextButtonType button.Type = `text`

	WithdrawLabel button.Text = `вывести на карту`
)

func ModeChoose() [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: WithdrawLabel,
					Payload: button.Payload{
						Command: WithdrawIdButton,
					}.String(),
				},
			},
		},
	}
}
