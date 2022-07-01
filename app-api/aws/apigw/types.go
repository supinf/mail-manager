package apigw

// APIStage API ステージ
type APIStage struct {
	APIID *string
	Name  *string
}

// Throttle スロットル（リクエストレート）
type Throttle struct {
	// RateLimit レート上限（1秒あたりのリクエスト数の平均）
	RateLimit *float64
	// BurstLimit バースト上限（バーストとして許容されるリクエスト数）
	BurstLimit *int64
}

// Quota クォーター（特定期間内の最大リクエスト可能数）
type Quota struct {
	// Period 期間
	Period *Period
	// Limit 上限数
	Limit *int64
	// Offset 開始日（ex. Period=WEEK の時, Offset=0 が Sunday, Offset=1 が Monday）
	Offset *int64
}

// PatchOperation 更新操作内容
type PatchOperation struct {
	// From コピー対象となるリソース内の場所
	From *string
	// Op 変更方法
	Op *Operation
	// Path 変更場所
	Path *Path
	// Value 設定値
	Value *string
}

// Period 期間
type Period string

// Operation 変更方法
type Operation string

// Path 変更場所
type Path string

// Enabled 有効/無効
type Enabled string

const (
	// PeriodDay 日
	PeriodDay Period = "DAY"
	// PeriodWeek 週
	PeriodWeek Period = "WEEK"
	// PeriodMonth 月
	PeriodMonth Period = "Month"
)

const (
	// OperationAdd 追加
	OperationAdd Operation = "add"
	// OperationRemove 削除
	OperationRemove Operation = "remove"
	// OperationReplace 変更
	OperationReplace Operation = "replace"
	// OperationCopy コピー
	OperationCopy Operation = "copy"
)

const (
	// PathCustomerID CustomerID
	PathCustomerID Path = "/customerId"
	// PathDescription 説明
	PathDescription Path = "/description"
	// PathEnabled 有効/無効
	PathEnabled Path = "/enabled"
	// PathLabels ラベル
	PathLabels Path = "/labels"
	// PathName 名前
	PathName Path = "/name"
	// PathStages ステージ
	PathStages Path = "/stages"
)

const (
	// Enable 有効
	Enable Enabled = "true"
	// Disable 無効
	Disable Enabled = "false"
)

// String Period を文字列に変換します
func (p Period) String() string {
	return string(p)
}

// ConvertPeriod 文字列を Period に変換します
func ConvertPeriod(value string) Period {
	switch value {
	case PeriodDay.String():
		return PeriodDay
	case PeriodWeek.String():
		return PeriodWeek
	case PeriodMonth.String():
		return PeriodMonth
	default:
		return PeriodDay
	}
}

// String Operation を文字列に変換します
func (op Operation) String() string {
	return string(op)
}

// String Path を文字列に変換します
func (p Path) String() string {
	return string(p)
}

// String Enabled を文字列に変換します
func (e Enabled) String() string {
	return string(e)
}
