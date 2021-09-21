package hmail

import "github.com/emersion/go-imap"

// ImapFetch 收邮件
func ImapFetch(auth *Auth, criteria ...*Criteria) (mail []*Mail) {
	if criteria != nil {
	}

	return
}

type Criteria struct {
	Search *imap.SearchCriteria // 邮件查询条件
}
