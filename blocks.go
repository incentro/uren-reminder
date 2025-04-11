package main

import (
	"github.com/slack-go/slack"
)

func createBlock(msg string, url string) []slack.Block {
	blocks := []slack.Block{
		slack.NewHeaderBlock(
			&slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: ":afas: Uren reminder :afas:",
			},
		),
		slack.NewDividerBlock(),
		slack.NewSectionBlock(
			&slack.TextBlockObject{
				Type: slack.MarkdownType,
				Text: msg,
			},
			nil,
			nil,
		),
		slack.NewActionBlock(
			"",
			slack.NewButtonBlockElement(
				"",
				"",
				&slack.TextBlockObject{
					Type: slack.PlainTextType,
					Text: ":calendar: Ga naar Afas",
				},
			).WithStyle(slack.StylePrimary).WithURL(url),
		),
		slack.NewDividerBlock(),
		slack.NewContextBlock("footer",
			&slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: ":github: Geïnteresseerd in mijn code? Check https://github.com/incentro/uren-reminder. :maple_leaf:",
			}),
	}

	return blocks
}
