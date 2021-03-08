package main

import (
	"InteliBot/constants"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

//Init configs
const (
	Token              = ""
	GuildID            = ""
	DefaultChannelName = "default-channel"

	JiraSiteName = "example"
)

var _numberOfAlerts int = 0

var _channelsDict map[string]*discordgo.Channel = make(map[string]*discordgo.Channel)
var _botID string
var _botSession *discordgo.Session

func main() {

	botSession, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_botSession = botSession

	user, err := _botSession.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	_botID = user.ID

	err = _botSession.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	getAllChannels()

	fmt.Println("Bot is running!")

	http.HandleFunc("/", requestCallBack)
	http.ListenAndServe(":6002", nil)

	return
}

func getAllChannels() {

	channels, err := _botSession.GuildChannels(GuildID)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, channel := range channels {
		_channelsDict[channel.Name] = channel
	}
}

func requestCallBack(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Fprint(w, strconv.Itoa(_numberOfAlerts)+" Alerts are sended")
		return
	}
	if r.Method == http.MethodPost {
		webhookResponse := WebhookResponse{}

		err := json.NewDecoder(r.Body).Decode(&webhookResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = sendMessage(&webhookResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		_numberOfAlerts++

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func sendMessage(w *WebhookResponse) error {

	channel := searchChannel(w.getProjectName(true))

	message := buildMessage(w)
	message.Embed.Timestamp = time.Now().Format(time.RFC3339)

	_, err := _botSession.ChannelMessageSendComplex(channel.ID, message)

	return err
}

func buildMessage(w *WebhookResponse) *discordgo.MessageSend {

	switch w.WebhookEvent {

	case constants.ProjectCreated:
		return w.projectCreatedAlert()

	case constants.ProjectDeleted:
		return w.projectDeletedAlert()

	case constants.ProjectUpdated:
		return w.projectUpdatedAlert()

	case constants.IssueCreated:
		return w.issueCreatedAlert()

	case constants.IssueDeleted:
		return w.issueDeletedAlert()

	case constants.IssueUpdated:
		return w.issueUpdatedAlert()

	case constants.CommentCreated:
		return w.commentCreatedAlert()

	case constants.CommentDeleted:
		return w.commentDeletedAlert()

	case constants.CommentUpdated:
		return w.commentUpdatedAlert()

	default:
		return w.eventNotIdentifiedAlert()

	}
}

func (w *WebhookResponse) eventNotIdentifiedAlert() *discordgo.MessageSend {

	message := discordgo.MessageSend{}

	message.Content = "Event not identified: " + w.WebhookEvent

	return &message
}

func searchChannel(channelName string) *discordgo.Channel {

	channel := _channelsDict[channelName]
	if channel == nil {
		getAllChannels()

		channel = _channelsDict[channelName]
		if channel == nil {
			channel = _channelsDict[DefaultChannelName]
		}
	}

	return channel
}

func getColor(issueType string) int {

	if issueType == constants.Bug {
		return constants.Red
	}
	if issueType == constants.Task || issueType == constants.TaskPT {
		return constants.Blue
	}
	if issueType == constants.Story || issueType == constants.StoryPT {
		return constants.Green
	}
	if issueType == constants.Epic {
		return constants.Purple
	}

	return constants.Grey
}
