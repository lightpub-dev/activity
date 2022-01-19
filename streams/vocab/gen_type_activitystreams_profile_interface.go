// Code generated by astool. DO NOT EDIT.

package vocab

// A Profile is a content object that describes another Object, typically used to
// describe Actor Type objects. The describes property is used to reference
// the object being described by the profile.
//
// Example 59 (https://www.w3.org/TR/activitystreams-vocabulary/#ex184a-jsonld):
//   {
//     "describes": {
//       "name": "Sally Smith",
//       "type": "Person"
//     },
//     "summary": "Sally's Profile",
//     "type": "Profile"
//   }
type ActivityStreamsProfile interface {
	// GetActivityStreamsAltitude returns the "altitude" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAltitude() ActivityStreamsAltitudeProperty
	// GetActivityStreamsAttachment returns the "attachment" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAttachment() ActivityStreamsAttachmentProperty
	// GetActivityStreamsAttributedTo returns the "attributedTo" property if
	// it exists, and nil otherwise.
	GetActivityStreamsAttributedTo() ActivityStreamsAttributedToProperty
	// GetActivityStreamsAudience returns the "audience" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAudience() ActivityStreamsAudienceProperty
	// GetActivityStreamsBcc returns the "bcc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBcc() ActivityStreamsBccProperty
	// GetActivityStreamsBto returns the "bto" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBto() ActivityStreamsBtoProperty
	// GetActivityStreamsCc returns the "cc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsCc() ActivityStreamsCcProperty
	// GetActivityStreamsContent returns the "content" property if it exists,
	// and nil otherwise.
	GetActivityStreamsContent() ActivityStreamsContentProperty
	// GetActivityStreamsContext returns the "context" property if it exists,
	// and nil otherwise.
	GetActivityStreamsContext() ActivityStreamsContextProperty
	// GetActivityStreamsDescribes returns the "describes" property if it
	// exists, and nil otherwise.
	GetActivityStreamsDescribes() ActivityStreamsDescribesProperty
	// GetActivityStreamsDuration returns the "duration" property if it
	// exists, and nil otherwise.
	GetActivityStreamsDuration() ActivityStreamsDurationProperty
	// GetActivityStreamsEndTime returns the "endTime" property if it exists,
	// and nil otherwise.
	GetActivityStreamsEndTime() ActivityStreamsEndTimeProperty
	// GetActivityStreamsGenerator returns the "generator" property if it
	// exists, and nil otherwise.
	GetActivityStreamsGenerator() ActivityStreamsGeneratorProperty
	// GetActivityStreamsIcon returns the "icon" property if it exists, and
	// nil otherwise.
	GetActivityStreamsIcon() ActivityStreamsIconProperty
	// GetActivityStreamsImage returns the "image" property if it exists, and
	// nil otherwise.
	GetActivityStreamsImage() ActivityStreamsImageProperty
	// GetActivityStreamsInReplyTo returns the "inReplyTo" property if it
	// exists, and nil otherwise.
	GetActivityStreamsInReplyTo() ActivityStreamsInReplyToProperty
	// GetActivityStreamsLikes returns the "likes" property if it exists, and
	// nil otherwise.
	GetActivityStreamsLikes() ActivityStreamsLikesProperty
	// GetActivityStreamsLocation returns the "location" property if it
	// exists, and nil otherwise.
	GetActivityStreamsLocation() ActivityStreamsLocationProperty
	// GetActivityStreamsMediaType returns the "mediaType" property if it
	// exists, and nil otherwise.
	GetActivityStreamsMediaType() ActivityStreamsMediaTypeProperty
	// GetActivityStreamsName returns the "name" property if it exists, and
	// nil otherwise.
	GetActivityStreamsName() ActivityStreamsNameProperty
	// GetActivityStreamsObject returns the "object" property if it exists,
	// and nil otherwise.
	GetActivityStreamsObject() ActivityStreamsObjectProperty
	// GetActivityStreamsPreview returns the "preview" property if it exists,
	// and nil otherwise.
	GetActivityStreamsPreview() ActivityStreamsPreviewProperty
	// GetActivityStreamsPublished returns the "published" property if it
	// exists, and nil otherwise.
	GetActivityStreamsPublished() ActivityStreamsPublishedProperty
	// GetActivityStreamsReplies returns the "replies" property if it exists,
	// and nil otherwise.
	GetActivityStreamsReplies() ActivityStreamsRepliesProperty
	// GetActivityStreamsSensitive returns the "sensitive" property if it
	// exists, and nil otherwise.
	GetActivityStreamsSensitive() ActivityStreamsSensitiveProperty
	// GetActivityStreamsShares returns the "shares" property if it exists,
	// and nil otherwise.
	GetActivityStreamsShares() ActivityStreamsSharesProperty
	// GetActivityStreamsSource returns the "source" property if it exists,
	// and nil otherwise.
	GetActivityStreamsSource() ActivityStreamsSourceProperty
	// GetActivityStreamsStartTime returns the "startTime" property if it
	// exists, and nil otherwise.
	GetActivityStreamsStartTime() ActivityStreamsStartTimeProperty
	// GetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	GetActivityStreamsSummary() ActivityStreamsSummaryProperty
	// GetActivityStreamsTag returns the "tag" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTag() ActivityStreamsTagProperty
	// GetActivityStreamsTo returns the "to" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTo() ActivityStreamsToProperty
	// GetActivityStreamsUpdated returns the "updated" property if it exists,
	// and nil otherwise.
	GetActivityStreamsUpdated() ActivityStreamsUpdatedProperty
	// GetActivityStreamsUrl returns the "url" property if it exists, and nil
	// otherwise.
	GetActivityStreamsUrl() ActivityStreamsUrlProperty
	// GetForgeFedTeam returns the "team" property if it exists, and nil
	// otherwise.
	GetForgeFedTeam() ForgeFedTeamProperty
	// GetForgeFedTicketsTrackedBy returns the "ticketsTrackedBy" property if
	// it exists, and nil otherwise.
	GetForgeFedTicketsTrackedBy() ForgeFedTicketsTrackedByProperty
	// GetForgeFedTracksTicketsFor returns the "tracksTicketsFor" property if
	// it exists, and nil otherwise.
	GetForgeFedTracksTicketsFor() ForgeFedTracksTicketsForProperty
	// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
	GetJSONLDId() JSONLDIdProperty
	// GetJSONLDType returns the "type" property if it exists, and nil
	// otherwise.
	GetJSONLDType() JSONLDTypeProperty
	// GetTypeName returns the name of this type.
	GetTypeName() string
	// GetUnknownProperties returns the unknown properties for the Profile
	// type. Note that this should not be used by app developers. It is
	// only used to help determine which implementation is LessThan the
	// other. Developers who are creating a different implementation of
	// this type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the Profile type extends from the other
	// type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Profile is lesser, with an arbitrary but
	// stable determination.
	LessThan(o ActivityStreamsProfile) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActivityStreamsAltitude sets the "altitude" property.
	SetActivityStreamsAltitude(i ActivityStreamsAltitudeProperty)
	// SetActivityStreamsAttachment sets the "attachment" property.
	SetActivityStreamsAttachment(i ActivityStreamsAttachmentProperty)
	// SetActivityStreamsAttributedTo sets the "attributedTo" property.
	SetActivityStreamsAttributedTo(i ActivityStreamsAttributedToProperty)
	// SetActivityStreamsAudience sets the "audience" property.
	SetActivityStreamsAudience(i ActivityStreamsAudienceProperty)
	// SetActivityStreamsBcc sets the "bcc" property.
	SetActivityStreamsBcc(i ActivityStreamsBccProperty)
	// SetActivityStreamsBto sets the "bto" property.
	SetActivityStreamsBto(i ActivityStreamsBtoProperty)
	// SetActivityStreamsCc sets the "cc" property.
	SetActivityStreamsCc(i ActivityStreamsCcProperty)
	// SetActivityStreamsContent sets the "content" property.
	SetActivityStreamsContent(i ActivityStreamsContentProperty)
	// SetActivityStreamsContext sets the "context" property.
	SetActivityStreamsContext(i ActivityStreamsContextProperty)
	// SetActivityStreamsDescribes sets the "describes" property.
	SetActivityStreamsDescribes(i ActivityStreamsDescribesProperty)
	// SetActivityStreamsDuration sets the "duration" property.
	SetActivityStreamsDuration(i ActivityStreamsDurationProperty)
	// SetActivityStreamsEndTime sets the "endTime" property.
	SetActivityStreamsEndTime(i ActivityStreamsEndTimeProperty)
	// SetActivityStreamsGenerator sets the "generator" property.
	SetActivityStreamsGenerator(i ActivityStreamsGeneratorProperty)
	// SetActivityStreamsIcon sets the "icon" property.
	SetActivityStreamsIcon(i ActivityStreamsIconProperty)
	// SetActivityStreamsImage sets the "image" property.
	SetActivityStreamsImage(i ActivityStreamsImageProperty)
	// SetActivityStreamsInReplyTo sets the "inReplyTo" property.
	SetActivityStreamsInReplyTo(i ActivityStreamsInReplyToProperty)
	// SetActivityStreamsLikes sets the "likes" property.
	SetActivityStreamsLikes(i ActivityStreamsLikesProperty)
	// SetActivityStreamsLocation sets the "location" property.
	SetActivityStreamsLocation(i ActivityStreamsLocationProperty)
	// SetActivityStreamsMediaType sets the "mediaType" property.
	SetActivityStreamsMediaType(i ActivityStreamsMediaTypeProperty)
	// SetActivityStreamsName sets the "name" property.
	SetActivityStreamsName(i ActivityStreamsNameProperty)
	// SetActivityStreamsObject sets the "object" property.
	SetActivityStreamsObject(i ActivityStreamsObjectProperty)
	// SetActivityStreamsPreview sets the "preview" property.
	SetActivityStreamsPreview(i ActivityStreamsPreviewProperty)
	// SetActivityStreamsPublished sets the "published" property.
	SetActivityStreamsPublished(i ActivityStreamsPublishedProperty)
	// SetActivityStreamsReplies sets the "replies" property.
	SetActivityStreamsReplies(i ActivityStreamsRepliesProperty)
	// SetActivityStreamsSensitive sets the "sensitive" property.
	SetActivityStreamsSensitive(i ActivityStreamsSensitiveProperty)
	// SetActivityStreamsShares sets the "shares" property.
	SetActivityStreamsShares(i ActivityStreamsSharesProperty)
	// SetActivityStreamsSource sets the "source" property.
	SetActivityStreamsSource(i ActivityStreamsSourceProperty)
	// SetActivityStreamsStartTime sets the "startTime" property.
	SetActivityStreamsStartTime(i ActivityStreamsStartTimeProperty)
	// SetActivityStreamsSummary sets the "summary" property.
	SetActivityStreamsSummary(i ActivityStreamsSummaryProperty)
	// SetActivityStreamsTag sets the "tag" property.
	SetActivityStreamsTag(i ActivityStreamsTagProperty)
	// SetActivityStreamsTo sets the "to" property.
	SetActivityStreamsTo(i ActivityStreamsToProperty)
	// SetActivityStreamsUpdated sets the "updated" property.
	SetActivityStreamsUpdated(i ActivityStreamsUpdatedProperty)
	// SetActivityStreamsUrl sets the "url" property.
	SetActivityStreamsUrl(i ActivityStreamsUrlProperty)
	// SetForgeFedTeam sets the "team" property.
	SetForgeFedTeam(i ForgeFedTeamProperty)
	// SetForgeFedTicketsTrackedBy sets the "ticketsTrackedBy" property.
	SetForgeFedTicketsTrackedBy(i ForgeFedTicketsTrackedByProperty)
	// SetForgeFedTracksTicketsFor sets the "tracksTicketsFor" property.
	SetForgeFedTracksTicketsFor(i ForgeFedTracksTicketsForProperty)
	// SetJSONLDId sets the "id" property.
	SetJSONLDId(i JSONLDIdProperty)
	// SetJSONLDType sets the "type" property.
	SetJSONLDType(i JSONLDTypeProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
