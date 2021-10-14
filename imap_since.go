package hmail

import (
	"errors"
	"github.com/emersion/go-imap"
	"log"
	"time"
)

// ImapFetchSince 收取指定时间之后的邮件
func ImapFetchSince(auth *Auth, timeSince time.Time, timeBefore time.Time) (list []*Mail, err error) {

	c, err := ImapClient(auth)
	if err != nil {
		return
	}
	defer c.Logout()

	mbox, err := c.Select("INBOX", false) // Select INBOX
	if err != nil {
		return
	}

	// Get the last message
	if mbox.Messages == 0 {
		err = errors.New("ERR_IMAP:No message in mailbox")
		return
	}

	// 查询指定时间之后的邮件
	criteria := imap.NewSearchCriteria()
	if !timeSince.IsZero() {
		criteria.Since = timeSince
	}
	if !timeBefore.IsZero() {
		criteria.Before = timeBefore
	}
	ids, err := c.Search(criteria)
	if err != nil {
		log.Println("Search:", err)
		return
	}
	log.Println("IDs found:", ids)
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(ids...)

	// Get the whole message body
	var section imap.BodySectionName
	messages := make(chan *imap.Message, len(ids))
	go func() {
		if err = c.Fetch(seqSet, []imap.FetchItem{section.FetchItem()}, messages); err != nil {
			log.Println(err)
		}
	}()

	for msg := range messages {
		list = append(list, messageToMail(msg, section))
	}
	return
}
