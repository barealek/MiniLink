package models

type Mini struct {
	ID       uint64 `json:"id" bson:"_id,omitempty"`  // ID of the miniLink
	Redirect string `json:"link" bson:"link"`         // Link to redirect to
	MiniLink string `json:"minilink" bson:"minilink"` // The path of the minilink
	Clicks   uint64 `json:"clicks" bson:"clicks"`     // Amount of clicks
	Owner    string `json:"owner" bson:"owner"`       // The owner of the minilink
	Random   bool   `json:"random" bson:"random"`     // If a custom minilink is set or not
}
