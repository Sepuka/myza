package button

import "github.com/sepuka/vkbotserver/api/button"

const (
	StartIdButton        = `start`
	WithdrawIdButton     = `withdraw`
	GenerateAddrIdButton = `generate_addr`

	TextButtonType button.Type = `text`

	WithdrawLabel button.Text = `вывести на карту`
	GenerateAddr  button.Text = `новый адрес`
)

// Buttons builds base list of buttons to answer
func Buttons() [][]button.Button {
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

// ButtonsWithAddr builds base list of buttons to answer
func ButtonsWithAddr() [][]button.Button {
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
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: GenerateAddr,
					Payload: button.Payload{
						Command: GenerateAddrIdButton,
					}.String(),
				},
			},
		},
	}
}
