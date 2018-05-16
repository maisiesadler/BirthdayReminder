package birthday_reminder

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

type SlackSession struct {
	sendMessage chan string
	users       []string
}

func Go() {
	// persist.Add("testing", "123")
	token := os.Getenv("SLACK_TOKEN")
	bot := CreateBot(token)
	bot.sendMessage <- "Starting Up"

	ch := initNoon()
	for {
		select {
		case _ = <-ch:
			reminders := getBirthdayReminders()
			for _, reminder := range reminders {
				bot.sendMessage <- reminder
			}
		}
	}
}

// CreateBot Creates and Starts bot
func CreateBot(token string) SlackSession {
	// v, _ := persist.Get("test")
	fmt.Printf("ST: %v\n", token)
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	sendToUsers := []string{}
	users, err := rtm.GetUsers()
	if err == nil {
		fmt.Print("usrs")
		for _, usr := range users {
			fmt.Print(usr.Name)
			if usr.Name == "maisie" {
				id, found := getIMChannelForUser(api, usr.ID)
				if found {
					sendToUsers = append(sendToUsers, id)
				}
			}
		}
	}
	st := fmt.Sprintf("%s%d%s", "found", len(sendToUsers), "users")
	fmt.Println(st)
	session := SlackSession{sendMessage: make(chan string), users: sendToUsers}
	go sendMessages(rtm, session)
	go ListenToBot(rtm)
	return session
}

func getIMChannelForUser(api *slack.Client, userID string) (string, bool) {
	channels, _ := api.GetIMChannels()
	for _, channel := range channels {
		if channel.User == userID {
			// rtm.SendMessage(rtm.NewOutgoingMessage("dgf", usr.ID))
			fmt.Println(channel)
			// return channel.conversation.ID
			return channel.ID, true
			// sendToUsers = append(sendToUsers, usr)
		}
	}
	return "", false
}

func sendMessages(rtm *slack.RTM, session SlackSession) {
	for {
		select {
		case message := <-session.sendMessage:
			for _, usrID := range session.users {
				rtm.SendMessage(rtm.NewOutgoingMessage(message, usrID))
			}
		}
	}
}

func ListenToBot(rtm *slack.RTM) {
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				//	info := rtm.GetInfo()
				//prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				// user, _ := rtm.GetUserInfo(ev.User)

				// user := info.GetUserByID(ev.User)
				// output, _ := persist.Get(ev.Text)
				// fmt.Printf("%v ", output)

				// //if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
				rtm.SendMessage(rtm.NewOutgoingMessage("hi", ev.Channel))
				//}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}
