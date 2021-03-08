package main

import "github.com/bwmarrin/discordgo"

func (w *WebhookResponse) commentCreatedAlert() *discordgo.MessageSend {
	return internalCreateCommentMessage(w, " Commented on a ")
}

func (w *WebhookResponse) commentDeletedAlert() *discordgo.MessageSend {
	return internalCreateCommentMessage(w, " Deleted a Comment on a ")
}

func (w *WebhookResponse) commentUpdatedAlert() *discordgo.MessageSend {
	return internalCreateCommentMessage(w, " Updated a Comment on a ")
}

func internalCreateCommentMessage(w *WebhookResponse, event string) *discordgo.MessageSend {

	message := discordgo.MessageSend{}

	embed := discordgo.MessageEmbed{}

	user := w.Comment.Author

	author := discordgo.MessageEmbedAuthor{}
	author.Name = user.Name + event + w.Issue.Fields.IssueType.Name
	author.IconURL = user.AvatarURL["24x24"]
	author.URL = user.getProfileURL()

	embed.Author = &author
	embed.Title = w.Issue.Key + ": " + w.Issue.Fields.Summary
	embed.Description = w.Comment.Body
	embed.URL = w.Issue.getIssueURL() + "?focusedCommentId=" + w.Comment.ID
	embed.Color = getColor(w.Issue.Fields.IssueType.Name)

	message.Embed = &embed

	return &message
}
