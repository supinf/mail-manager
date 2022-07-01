package ses

// validateBulkSendEntryCount 一括送信エントリの宛先(To, CC, BCC)件数が制限を超えていないかどうかを判定します
func validateBulkSendEntryCount(entries []*BulkEntry) bool {
	if entries == nil {
		return true
	}

	count := 0
	for _, entry := range entries {
		if entry == nil {
			continue
		}
		count += len(entry.To)
		count += len(entry.Cc)
		count += len(entry.Bcc)
	}

	return BulkSendEntryMaximumCount >= count
}
