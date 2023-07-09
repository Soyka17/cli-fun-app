# CLI-FUN-APP

Mini Project, written in one sleepless night on the train.
I gave a similar task to my students on the python start course, and it seemed wildly interesting to me.
On another train ride with no internet, I wanted to write my own version without using any packages

The task was originally posed as
"Write your own banking application client, with the ability to deposit and withdraw money"
As an additional task, it was proposed to add the function of transferring money between different accounts, but no one got to it in the classroom (+1 to my desire to write it myself)


When starting the application, we are greeted with a main unauthorized window:

    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                    Press 1 to create new account                             |
    |                    Press 2 to login into existing account                    |
    |                                                                              |
    |                    Press 9 to exit                                           |
    |                                                                              |
    + — — — — — — — — — — — — — — — | Home screen | — — — — — — — — — — — — — — —  +

Let's create an account <1>

    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                       Create a Experience Bank account                       |
    |                                                                              |
    |            ID is issued once when creating an account                        |
    |            Remember it, you will need it to log into your account            |
    |                                                                              |
    |            Your BANK ID:1                                                    |
    |                                                                              |
    |            Required fields to create an account:                             |
    |                -Name                                                         |
    |                -Password                                                     |
    |                                                                              |
    + — — — — — — — — — — — — — —  | SignIn screen | — — — — — — — — — — — — — — — +
    
    Enter the name:SEREGA
    Enter the password:123

After successful registration, we are returned to the previous window. Login to the created account <2>

    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                      Login into Experience Bank account                      |
    |                                                                              |
    |                    Required fields to login into account:                    |
    |                                   BANK ID                                    |
    |                                   Password                                   |
    |                                                                              |
    + — — — — — — — — — — — — — —  | SignIn screen | — — — — — — — — — — — — — — — +

After logging in, we get to the main window
    
    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                       You are in SEREGA's bank account                       |
    |                                                                              |
    |                          Press 1 to check balance                            |
    |                          Press 2 to deposit money                            |
    |                          Press 3 to withdraw money                           |
    |                          Press 4 to send money                               |
    |                                                                              |
    |                          Press 9 to exit                                     |
    |                                                                              |
    |                                                                              |
    + — — — — — — — — — — — —  | SEREGA's bank account | — — — — — — — — — — — — — +

Let's check the balance <1>

    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                                Balance: 0                                    |
    |                                                                              |
    |                    Press any key to return to main menu                      |
    |                                                                              |
    |                                                                              |
    + — — — — — — — — — — — —  | SEREGA's bank account | — — — — — — — — — — — — — +

Let's try to deposit money <2>

    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                       Enter amount of money to deposit                       |
    |                                                                              |
    |                      Write CANCEL to back to main menu                       |
    |                                                                              |
    + — — — — — — — — — — — —  | SEREGA's bank account | — — — — — — — — — — — — — +

    Deposit amount:2337

 < Enter >

    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                   The money has been successfully credited                   |
    |                              Your balance: 2337                              |
    |                                                                              |
    |                    Press any key to return to main menu                      |
    |                                                                              |
    + — — — — — — — — — — — —  | SEREGA's bank account | — — — — — — — — — — — — — +

Let's try to withdraw money <3>
    
    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                      Enter amount of money to withdraw                       |
    |                              Your balance: 2337                              |
    |                                                                              |
    |                      Write CANCEL to back to main menu                       |
    |                                                                              |
    + — — — — — — — — — — — —  | SEREGA's bank account | — — — — — — — — — — — — — +

    Withdraw amount:1000

< Enter >
  
    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                  The money has been successfully withdrawn                   |
    |                              Your balance: 1337                              |
    |                                                                              |
    |                    Press any key to return to main menu                      |
    |                                                                              |
    + — — — — — — — — — — — —  | SEREGA's bank account | — — — — — — — — — — — — — +
    

Let's try to transfer money <3>
(Before that, I created an account with ID=2)

    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                          Enter id of money receiver                          |
    |                        Enter amount of money to send                         |
    |                              Your balance: 1337                              |
    |                                                                              |
    |                      Write CANCEL to back to main menu                       |
    |                                                                              |
    + — — — — — — — — — — — —  | SEREGA's bank account | — — — — — — — — — — — — — +

    Receiver ID:2
    Send amount:1337
    

< Enter >


    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                     The money has been successfully sent                     |
    |                               Your balance: 0                                |
    |                                                                              |
    |                    Press any key to return to main menu                      |
    |                                                                              |
    + — — — — — — — — — — — —  | SEREGA's bank account | — — — — — — — — — — — — — +
    

< Enter >

Now let's go to the second account and see the balance

    + — — — — — — — — — — — — | Experience Bank Terminal | — — — — — — — — — — — — +
    |                                                                              |
    |                                Balance: 1337                                 |
    |                                                                              |
    |                    Press any key to return to main menu                      |
    |                                                                              |
    |                                                                              |
    + — — — — — — — — — — —  | NOT SEREGA's bank account | — — — — — — — — — — — — +

Also, loading animation implemented between all windows
