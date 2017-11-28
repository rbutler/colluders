package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/rbutler/colluders/models"
)

var token = os.Getenv("TOKEN")
var groupID = os.Getenv("GROUP_ID")

var fetchLimit = uint64(100)

func main() {
	messages, err := getAllMessages()
	if err != nil {
		fmt.Println("failed to fetch messages, exiting")
		return
	}

	userMap := make(models.Users)
	for _, message := range messages {
		if userMap[message.UserID] == nil {
			userMap[message.UserID] = &models.User{
				ID:           message.UserID,
				Name:         message.UserName,
				Hearts:       0,
				HeartsGiven:  0,
				MessageCount: 1,
			}
		} else {
			userMap[message.UserID].MessageCount += 1
			userMap[message.UserID].Name = message.UserName
		}

		for _, favBy := range message.FavoritedBy {
			userMap[message.UserID].Hearts += 1

			if userMap[favBy] == nil {
				userMap[favBy] = &models.User{
					ID:          favBy,
					HeartsGiven: 1,
				}
			} else {
				userMap[favBy].HeartsGiven += 1
			}
		}
	}

	users := make([]models.User, 0, len(userMap))
	for _, value := range userMap {
		users = append(users, *value)
	}
	sort.Sort(models.ByHearts(users))

	println("")
	printUsers(users)
	println("")
}

func getAllMessages() ([]models.Message, error) {
	messages := []models.Message{}

	messageResponse, err := getMessages("0", fetchLimit)
	if err != nil {
		fmt.Println(err)
		return messages, err
	}

	messages = messageResponse.Response.Messages

	count := messageResponse.Response.Count
	if count <= fetchLimit {
		return messages, err
	}
	lastID := messages[len(messages)-1].ID

	for i := uint64(fetchLimit); i <= count; i += fetchLimit {
		messageResponse, err = getMessages(lastID, fetchLimit)
		if err != nil {
			fmt.Println(err)
			return messages, err
		}
		respMessages := messageResponse.Response.Messages
		messages = append(messages, respMessages...)
		lastID = respMessages[len(respMessages)-1].ID

	}

	return messages, err
}

func getMessages(lastID string, limit uint64) (models.MessageResponse, error) {
	messageResponse := &models.MessageResponse{}
	url := fmt.Sprintf("https://api.groupme.com/v3/groups/%v/messages?token=%v&limit=%v", groupID, token, limit)
	if lastID != "0" {
		url = fmt.Sprintf("%v&before_id=%v", url, lastID)
	}
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return *messageResponse, err
	}

	err = json.NewDecoder(r.Body).Decode(messageResponse)
	if err != nil {
		fmt.Println(err)
		println("ugh")
		return *messageResponse, nil
	}

	return *messageResponse, nil
}

func printUsers(users []models.User) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 12, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Name\tID\tMessageCount\tHearts Received\tHeartsGiven\t")
	fmt.Fprintln(w, "----\t--\t------------\t---------------\t-----------\t")

	for _, u := range users {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t", u.Name, u.ID, u.MessageCount, u.Hearts, u.HeartsGiven))
	}

	w.Flush()
}
