package ui

import (
	"ExperienceBank/internal/controller/logger"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type CliUiRepositoryImpl struct {
	logger      logger.LoggerRepository
	windowWidth int
	outputDelay time.Duration
}

func NewCliUiRepositoryImpl(win int, d time.Duration, l logger.LoggerRepository) *CliUiRepositoryImpl {
	return &CliUiRepositoryImpl{windowWidth: win, outputDelay: d, logger: l}
}

func (c *CliUiRepositoryImpl) ReadIntValue(act string) (int, bool) {
	var strValue string
	var err error
	var value int
	errDelay := time.Millisecond * 500

	for {
		strValue, err = c.GetUserAction(act)
		if err != nil {
			_ = c.ShowMessage("Uncorrected "+act+", try again", errDelay, false)
			continue
		}
		if strValue == "CANCEL" {
			return 0, true
		}

		value, err = strconv.Atoi(strValue)
		if err != nil {
			_ = c.ShowMessage("Uncorrected "+act+", try again", errDelay, false)
			continue
		}
		if value <= 0 {
			_ = c.ShowMessage("Uncorrected "+act+", try again", errDelay, false)
			continue
		}
		break
	}

	return value, false
}

func (c *CliUiRepositoryImpl) ShowHomeScreen() error {
	c.ShowProcessingScreen()
	c.clearScr()
	c.showHeader(true)
	c.showLine("Press 1 to create new account         ", "|", " ", true)
	c.showLine("Press 2 to login into existing account", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Press 9 to exit                       ", "|", " ", true)
	c.showFooter("Home screen", true)
	return nil
}

func (c *CliUiRepositoryImpl) GetUserAction(msg string) (string, error) {
	fmt.Print("    " + msg + ":")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text, nil
}

func (c *CliUiRepositoryImpl) ShowSignInScreen(id int) error {
	c.clearScr()
	c.showHeader(true)
	c.showLine("Create a Experience Bank account", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("ID is issued once when creating an account            ", "|", " ", true)
	c.showLine("Remember it, you will need it to log into your account", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Your BANK ID:"+strconv.Itoa(id)+strings.Repeat(" ", 54-13-len(strconv.Itoa(id))), "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Required fields to create an account:                 ", "|", " ", true)
	c.showLine("    -Name                                             ", "|", " ", true)
	c.showLine("    -Password                                         ", "|", " ", true)
	c.showFooter("SignIn screen", true)
	return nil
}

func (c *CliUiRepositoryImpl) ShowLoginScreen() error {
	c.clearScr()
	c.showHeader(true)
	c.showLine("Login into Experience Bank account", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Required fields to login into account:", "|", " ", true)
	c.showLine("BANK ID ", "|", " ", true)
	c.showLine("Password", "|", " ", true)
	c.showFooter("SignIn screen", true)
	return nil
}

func (c *CliUiRepositoryImpl) ShowMainScreen(name string) error {
	c.clearScr()
	c.showHeader(true)
	c.showLine("You are in "+name+"'s bank account", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Press 1 to check balance ", "|", " ", true)
	c.showLine("Press 2 to deposit money ", "|", " ", true)
	c.showLine("Press 3 to withdraw money", "|", " ", true)
	c.showLine("Press 4 to send money    ", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Press 9 to exit          ", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showFooter(name+"'s bank account", true)
	return nil
}

func (c *CliUiRepositoryImpl) ShowBalanceScreen(name string, balance int) error {
	c.clearScr()
	c.showHeader(true)
	c.showLine("Balance: "+strconv.Itoa(balance), "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Press any key to return to main menu ", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showFooter(name+"'s bank account", true)
	return nil
}

func (c *CliUiRepositoryImpl) ShowDepositScreen(name string) error {
	c.clearScr()
	c.showHeader(true)
	c.showLine("Enter amount of money to deposit", "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Write CANCEL to back to main menu", "|", " ", true)
	c.showFooter(name+"'s bank account", true)
	return nil
}
func (c *CliUiRepositoryImpl) ShowSuccessfullyDepositScreen(name string, balance int) error {
	c.ShowProcessingScreen()
	c.clearScr()
	c.showHeader(true)
	c.showLine("The money has been successfully credited", "|", " ", true)
	c.showLine("Your balance: "+strconv.Itoa(balance), "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Press any key to return to main menu ", "|", " ", true)
	c.showFooter(name+"'s bank account", true)
	return nil
}

func (c *CliUiRepositoryImpl) ShowWithdrawScreen(name string, balance int) error {
	c.clearScr()
	c.showHeader(true)
	c.showLine("Enter amount of money to withdraw", "|", " ", true)
	c.showLine("Your balance: "+strconv.Itoa(balance), "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Write CANCEL to back to main menu", "|", " ", true)
	c.showFooter(name+"'s bank account", true)
	return nil
}

func (c *CliUiRepositoryImpl) ShowSuccessfullyWithdrawScreen(name string, balance int) error {
	c.ShowProcessingScreen()
	c.clearScr()
	c.showHeader(true)
	c.showLine("The money has been successfully withdrawn", "|", " ", true)
	c.showLine("Your balance: "+strconv.Itoa(balance), "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Press any key to return to main menu ", "|", " ", true)
	c.showFooter(name+"'s bank account", true)
	return nil
}
func (c *CliUiRepositoryImpl) ShowSendScreen(name string, balance int) error {
	c.clearScr()
	c.showHeader(true)
	c.showLine("Enter id of money receiver", "|", " ", true)
	c.showLine("Enter amount of money to send", "|", " ", true)
	c.showLine("Your balance: "+strconv.Itoa(balance), "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Write CANCEL to back to main menu", "|", " ", true)
	c.showFooter(name+"'s bank account", true)
	return nil
}

func (c *CliUiRepositoryImpl) ShowSuccessfullySendScreen(name string, balance int) error {
	c.ShowProcessingScreen()
	c.clearScr()
	c.showHeader(true)
	c.showLine("The money has been successfully sent", "|", " ", true)
	c.showLine("Your balance: "+strconv.Itoa(balance), "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Press any key to return to main menu ", "|", " ", true)
	c.showFooter(name+"'s bank account", true)
	return nil
}

func (c *CliUiRepositoryImpl) ShowExitScreen() error {

	c.clearScr()
	c.showHeader(true)
	c.showLine("Are you sure want to get out? y/N?", "|", " ", true)
	c.showFooter("Exit", true)

	return nil
}

func (c *CliUiRepositoryImpl) ShowErrorScreen(msg string) error {
	c.clearScr()
	fmt.Println()
	c.showHeader(true)
	c.showLine(msg, "|", " ", true)
	c.showLine("", "|", " ", true)
	c.showLine("Press any key to return to main menu ", "|", " ", true)
	c.showFooter("Error", true)
	fmt.Println()
	return nil
}

func (c *CliUiRepositoryImpl) ShowProcessingScreen() error {
	delta := 100 * time.Millisecond
	signs := []string{
		"", ".", "..", "...", "..", ".", "",
	}
	for i := 0; i < 24; i++ {
		dots := signs[i%7]
		c.clearScr()
		c.showHeader(false)
		c.showLine("", "|", " ", false)
		c.showLine("", "|", " ", false)
		c.showLine(dots+" Processing "+dots, "|", " ", false)
		c.showLine("", "|", " ", false)
		c.showFooter("Process screen", false)
		time.Sleep(delta)
	}
	return nil
}

func (c *CliUiRepositoryImpl) ShowMessage(msg string, delay time.Duration, isDots bool) error {
	fmt.Print("    " + msg + " ")
	if isDots {
		for i := 0; i < 3; i++ {
			fmt.Print(".")
			time.Sleep(delay / 3)
		}
	} else {
		time.Sleep(delay)
	}
	fmt.Println()
	return nil
}

func (c *CliUiRepositoryImpl) clearScr() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (c *CliUiRepositoryImpl) showHeader(isDelay bool) {
	fmt.Println()
	c.showLine("| Experience Bank Terminal |", "+", "—", isDelay)
	c.showLine("", "|", " ", isDelay)
}

func (c *CliUiRepositoryImpl) showFooter(page string, isDelay bool) {
	c.showLine("", "|", " ", isDelay)
	c.showLine("| "+page+" |", "+", "—", isDelay)
	fmt.Println()
}

func (c *CliUiRepositoryImpl) showLine(msg string, borders string, line string, isDelay bool) {
	spaceLen := c.windowWidth - len(borders)*2 - len(msg)
	spaceL := spaceLen / 2
	spaceR := spaceL
	if spaceLen%2 == 1 {
		spaceR++
	}
	result := "    " + borders + c.getSpacesStr(spaceL, line, " ") + msg + c.getSpacesStr(spaceR, line, " ") + borders

	fmt.Println(result)
	if isDelay {
		time.Sleep(c.outputDelay)
	}
}

func (c *CliUiRepositoryImpl) getSpacesStr(num int, symb string, space string) string {
	result := ""
	for i := 0; i < num-1; i++ {
		if i%2 == 0 {
			result += space
		} else {
			result += symb
		}
	}
	result += space
	return result

}
