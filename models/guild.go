package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guild struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	GuildID string `json:"guildid,omitempty" validate:"required"`
	GuildName string `json:"guildname,omitempty" validate:"required"`
	GuildOwner string `json:"guildowner,omitempty" validate:"required"`
	GuildOwnerID string `json:"guildownerid,omitempty" validate:"required"`
	Flags string `json:"flags,omitempty"`  // TODO: change this from a string to a proper flagging/seperator to objectify the flags of what the guild has enabled
}