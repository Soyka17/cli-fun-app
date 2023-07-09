package ui

import "time"

type UiRepository interface {
	ReadIntValue(act string) (int, bool)
	ShowHomeScreen() error
	ShowProcessingScreen() error
	GetUserAction(string) (string, error)
	ShowSignInScreen(int) error
	ShowLoginScreen() error
	ShowMainScreen(string) error
	ShowBalanceScreen(string, int) error
	ShowDepositScreen(string) error
	ShowSuccessfullyDepositScreen(string, int) error
	ShowWithdrawScreen(string, int) error
	ShowSuccessfullyWithdrawScreen(string, int) error
	ShowSendScreen(string, int) error
	ShowSuccessfullySendScreen(string, int) error
	ShowExitScreen() error
	ShowErrorScreen(string) error
	ShowMessage(string, time.Duration, bool) error
}
