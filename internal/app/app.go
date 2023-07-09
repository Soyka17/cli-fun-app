package app

import (
	"ExperienceBank/internal/controller/logger"
	"ExperienceBank/internal/entity"
	"ExperienceBank/internal/usecase/db"
	"ExperienceBank/internal/usecase/ui"
	"os"
	"time"
)

type App struct {
	Ui     ui.UiRepository
	Db     db.DatabaseRepository
	Logger logger.LoggerRepository
}

func NewApp(ui ui.UiRepository, db db.DatabaseRepository, l logger.LoggerRepository) *App {
	return &App{Ui: ui, Db: db, Logger: l}
}

func (a *App) Run() {
	a.Logger.Info("App start")
	a.unauthorizedLoop()
}

func (a *App) unauthorizedLoop() {
	err := a.Ui.ShowHomeScreen()
	if err != nil {
		a.Logger.Warn("Can't show home screen: " + err.Error())
	}
	act, _ := a.Ui.GetUserAction("Your choice")
	switch act {
	case "1":
		a.signInProcess()
	case "2":
		a.loginProcess()
	case "9":
		a.exitProcess()

	default:
		a.Ui.ShowErrorScreen("Unexpected button pressed")
		_, _ = a.Ui.GetUserAction("")
		a.unauthorizedLoop()
	}
}

func (a *App) signInProcess() {
	id := a.Db.GetNewId()
	a.Ui.ShowProcessingScreen()
	a.Ui.ShowSignInScreen(id)

	name, _ := a.Ui.GetUserAction("Enter the name")
	pass, _ := a.Ui.GetUserAction("Enter the password")

	newUser := entity.NewUser(id, name, 0, pass)
	a.Db.SaveNewUser(newUser)

	a.Ui.ShowMessage("SignIn succesfull, redirecting to main window", time.Second*2, true)

	a.unauthorizedLoop()
}

func (a *App) loginProcess() {
	errDelay := time.Millisecond * 500
	a.Ui.ShowProcessingScreen()
	a.Ui.ShowLoginScreen()
	var strId, pass string
	var isCancel bool
	var err error
	var id int
	var user *entity.User

	for {
		id, isCancel = a.Ui.ReadIntValue("ID")
		if isCancel {
			a.unauthorizedLoop()
		}

		user, err = a.Db.GetUserById(id)
		if err != nil {
			_ = a.Ui.ShowMessage("Uncorrected ID, try again", errDelay, false)
			continue
		}
		break

	}

	for {
		pass, err = a.Ui.GetUserAction("Enter the password")
		if err != nil {
			_ = a.Ui.ShowMessage("Uncorrected password, try again", errDelay, false)
			continue
		}
		if user.Password != pass {
			_ = a.Ui.ShowMessage("Wrong password for ID:"+strId, errDelay, false)
			continue
		}
		break
	}

	a.Ui.ShowMessage("Login succesfull", time.Second*2, true)
	a.mainLoop(user)
}

func (a *App) mainLoop(user *entity.User) {
	a.Ui.ShowMainScreen(user.Name)

	act, _ := a.Ui.GetUserAction("Your choice")
	switch act {
	case "1":
		a.balanceProcess(user)
	case "2":
		a.depositProcess(user)
	case "3":
		a.withdrawProcess(user)
	case "4":
		a.sendMoneyProcess(user)
	case "9":
		a.unauthorizedLoop()
	default:
		a.Ui.ShowErrorScreen("Unexpected button pressed")
		_, _ = a.Ui.GetUserAction("")
		a.mainLoop(user)
	}
}

func (a *App) balanceProcess(user *entity.User) {
	a.Ui.ShowProcessingScreen()
	a.Ui.ShowBalanceScreen(user.Name, user.Value)

	_, _ = a.Ui.GetUserAction("")

	a.mainLoop(user)
}

func (a *App) depositProcess(user *entity.User) {
	a.Ui.ShowProcessingScreen()
	a.Ui.ShowDepositScreen(user.Name)

	var isCancel bool
	var value int

	defer a.mainLoop(user)

	value, isCancel = a.Ui.ReadIntValue("Deposit amount")
	if isCancel {
		return
	}

	user.Value += value
	a.Ui.ShowSuccessfullyDepositScreen(user.Name, user.Value)
	_, _ = a.Ui.GetUserAction("")
}

func (a *App) withdrawProcess(user *entity.User) {
	a.Ui.ShowProcessingScreen()
	a.Ui.ShowWithdrawScreen(user.Name, user.Value)

	var value int
	var isCancel bool
	errDelay := time.Millisecond * 500
	defer a.mainLoop(user)
	for {
		value, isCancel = a.Ui.ReadIntValue("Withdraw amount")
		if isCancel {
			return
		}
		if value > user.Value {
			_ = a.Ui.ShowMessage("You don't have that much money", errDelay, false)
			continue
		}
		break
	}

	user.Value -= value
	a.Db.UpdateUserInfo(user)
	a.Ui.ShowSuccessfullyWithdrawScreen(user.Name, user.Value)
	_, _ = a.Ui.GetUserAction("")

}

func (a *App) sendMoneyProcess(user *entity.User) {
	a.Ui.ShowProcessingScreen()
	a.Ui.ShowSendScreen(user.Name, user.Value)

	var err error
	var id, value int
	var isCancel bool
	var userReceiver *entity.User
	errDelay := time.Millisecond * 500

	defer a.mainLoop(user)

	for {
		id, isCancel = a.Ui.ReadIntValue("Receiver ID")
		if isCancel {
			return
		}
		if id == user.Id {
			_ = a.Ui.ShowMessage("That's your ID, try again", errDelay, false)
			continue
		}
		userReceiver, err = a.Db.GetUserById(id)
		if err != nil {
			_ = a.Ui.ShowMessage("User not found", errDelay, false)
			continue
		}
		break
	}
	for {
		value, isCancel = a.Ui.ReadIntValue("Send amount")
		if isCancel {
			return
		}
		if value > user.Value {
			_ = a.Ui.ShowMessage("You don't have that much money", errDelay, false)
			continue
		}
		break
	}
	user.Value -= value
	a.Db.UpdateUserInfo(user)
	userReceiver.Value += value
	a.Db.UpdateUserInfo(userReceiver)

	a.Ui.ShowSuccessfullySendScreen(user.Name, user.Value)
	_, _ = a.Ui.GetUserAction("")
}

func (a *App) exitProcess() {
	a.Ui.ShowExitScreen()
	act, _ := a.Ui.GetUserAction("")
	switch act {
	case "y":
		os.Exit(0)
	case "N":
		a.unauthorizedLoop()
	default:
		a.Ui.ShowErrorScreen("Unexpected button pressed")
	}

}
