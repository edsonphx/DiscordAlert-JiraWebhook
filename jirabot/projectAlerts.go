package main

import (
	"github.com/bwmarrin/discordgo"
)

func (w *WebhookResponse) projectCreatedAlert() *discordgo.MessageSend {
	return internalCreateProjectMessage(w, " Created a Project")
}

func (w *WebhookResponse) projectDeletedAlert() *discordgo.MessageSend {
	return internalCreateProjectMessage(w, " Deleted a Project")
}

func (w *WebhookResponse) projectUpdatedAlert() *discordgo.MessageSend {
	return internalCreateProjectMessage(w, " Updated a Project")
}

func internalCreateProjectMessage(w *WebhookResponse, event string) *discordgo.MessageSend {

	message := discordgo.MessageSend{}

	embed := discordgo.MessageEmbed{}

	user := w.Project.ProjectLead

	author := discordgo.MessageEmbedAuthor{}
	author.Name = user.Name + event
	author.IconURL = user.AvatarURL["24x24"]
	author.URL = user.getProfileURL()

	embed.Author = &author
	embed.Title = w.getProjectName(false)
	embed.URL = w.Project.getProjectURL()
	embed.Color = getColor("")

	message.Embed = &embed

	return &message
}
