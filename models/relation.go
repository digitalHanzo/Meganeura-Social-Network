package models

//Relation - Model NXM
type Relation struct {
	UserID         string `bson:"userid" json:"userId"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId"`
}
