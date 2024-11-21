package main

import (
	"reflect"
	"testing"

	"github.com/slack-go/slack"
)

func TestCreateBlock(t *testing.T) {
	msg := "This is a test message."
	url := "https://example.com"

	expectedBlocks := []slack.Block{
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
	}

	actualBlocks := createBlock(msg, url)

	if !reflect.DeepEqual(expectedBlocks, actualBlocks) {
		t.Errorf("createBlock(%q, %q) = %v, want %v", msg, url, actualBlocks, expectedBlocks)
	}
}
