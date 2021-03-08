package main

import "github.com/bwmarrin/discordgo"

func (w *WebhookResponse) issueCreatedAlert() *discordgo.MessageSend {

	message := discordgo.MessageSend{}

	embed := discordgo.MessageEmbed{}

	user := w.User

	author := discordgo.MessageEmbedAuthor{}
	author.Name = user.Name + " Created a " + w.Issue.Fields.IssueType.Name
	author.IconURL = user.AvatarURL["24x24"]
	author.URL = user.getProfileURL()

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Type",
			Value:  w.Issue.Fields.IssueType.Name,
			Inline: true,
		},
		{
			Name:   "Priority",
			Value:  w.Issue.Fields.Priority.Name,
			Inline: true,
		},
		{
			Name:   "Status",
			Value:  w.Issue.Fields.Status.Name,
			Inline: true,
		},
	}

	embed.Author = &author
	embed.Title = w.Issue.Key + ": " + w.Issue.Fields.Summary
	embed.Description = w.Issue.Fields.Description
	embed.URL = w.Issue.getIssueURL()
	embed.Color = getColor(w.Issue.Fields.IssueType.Name)
	embed.Fields = fields

	message.Embed = &embed

	return &message
}

func (w *WebhookResponse) issueDeletedAlert() *discordgo.MessageSend {

	message := discordgo.MessageSend{}

	embed := discordgo.MessageEmbed{}

	user := w.User

	author := discordgo.MessageEmbedAuthor{}
	author.Name = user.Name + " Deleted a " + w.Issue.Fields.IssueType.Name
	author.IconURL = user.AvatarURL["24x24"]
	author.URL = user.getProfileURL()

	embed.Author = &author
	embed.Title = w.Issue.Key + ": " + w.Issue.Fields.Summary
	embed.URL = w.Issue.getIssueURL()
	embed.Color = getColor(w.Issue.Fields.IssueType.Name)

	message.Embed = &embed

	return &message
}

func (w *WebhookResponse) issueUpdatedAlert() *discordgo.MessageSend {

	message := discordgo.MessageSend{}

	embed := discordgo.MessageEmbed{}

	user := w.User

	author := discordgo.MessageEmbedAuthor{}
	author.Name = user.Name + " Updated the " + w.Changelog.Items[0].Field + " of a " + w.Issue.Fields.IssueType.Name
	author.IconURL = user.AvatarURL["24x24"]
	author.URL = user.getProfileURL()

	if w.Changelog.Items[0].Field == "description" {
		embed.Description = w.Issue.Fields.Description
	} else {
		lastValue := w.Changelog.Items[0].LastValue
		newValue := w.Changelog.Items[0].NewValue

		if lastValue == "" {
			lastValue = "-"
		}

		if newValue == "" {
			newValue = "-"
		}

		fields := []*discordgo.MessageEmbedField{
			{
				Name:   "From",
				Value:  lastValue,
				Inline: true,
			},
			{
				Name:   "To",
				Value:  newValue,
				Inline: true,
			},
		}
		embed.Fields = fields
	}

	embed.Author = &author
	embed.Title = w.Issue.Key + ": " + w.Issue.Fields.Summary

	embed.URL = w.Issue.getIssueURL()
	embed.Color = getColor(w.Issue.Fields.IssueType.Name)

	message.Embed = &embed

	return &message
}
