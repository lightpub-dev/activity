package propertynext

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// NextProperty is the functional property "next". It is permitted to be one of
// multiple value types. At most, one type of value can be present, or none at
// all. Setting a value will clear the other types of values so that only one
// of the 'Is' methods will return true. It is possible to clear all values,
// so that this property is empty.
type NextProperty struct {
	CollectionPageMember        vocab.CollectionPageInterface
	LinkMember                  vocab.LinkInterface
	MentionMember               vocab.MentionInterface
	OrderedCollectionPageMember vocab.OrderedCollectionPageInterface
	unknown                     []byte
	iri                         *url.URL
	alias                       string
}

// DeserializeNextProperty creates a "next" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeNextProperty(m map[string]interface{}, aliasMap map[string]string) (*NextProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "next"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "next")
	}
	if i, ok := m[propName]; ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &NextProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &NextProperty{
					CollectionPageMember: v,
					alias:                alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err == nil {
				this := &NextProperty{
					LinkMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap); err == nil {
				this := &NextProperty{
					MentionMember: v,
					alias:         alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &NextProperty{
					OrderedCollectionPageMember: v,
					alias:                       alias,
				}
				return this, nil
			}
		} else if str, ok := i.(string); ok {
			this := &NextProperty{
				alias:   alias,
				unknown: []byte(str),
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "next")
		}
	}
	return nil, nil
}

// NewNextProperty creates a new next property.
func NewNextProperty() *NextProperty {
	return &NextProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *NextProperty) Clear() {
	this.CollectionPageMember = nil
	this.LinkMember = nil
	this.MentionMember = nil
	this.OrderedCollectionPageMember = nil
	this.unknown = nil
	this.iri = nil
}

// GetCollectionPage returns the value of this property. When IsCollectionPage
// returns false, GetCollectionPage will return an arbitrary value.
func (this NextProperty) GetCollectionPage() vocab.CollectionPageInterface {
	return this.CollectionPageMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this NextProperty) GetIRI() *url.URL {
	return this.iri
}

// GetLink returns the value of this property. When IsLink returns false, GetLink
// will return an arbitrary value.
func (this NextProperty) GetLink() vocab.LinkInterface {
	return this.LinkMember
}

// GetMention returns the value of this property. When IsMention returns false,
// GetMention will return an arbitrary value.
func (this NextProperty) GetMention() vocab.MentionInterface {
	return this.MentionMember
}

// GetOrderedCollectionPage returns the value of this property. When
// IsOrderedCollectionPage returns false, GetOrderedCollectionPage will return
// an arbitrary value.
func (this NextProperty) GetOrderedCollectionPage() vocab.OrderedCollectionPageInterface {
	return this.OrderedCollectionPageMember
}

// HasAny returns true if any of the different values is set.
func (this NextProperty) HasAny() bool {
	return this.IsCollectionPage() ||
		this.IsLink() ||
		this.IsMention() ||
		this.IsOrderedCollectionPage() ||
		this.iri != nil
}

// IsCollectionPage returns true if this property has a type of "CollectionPage".
// When true, use the GetCollectionPage and SetCollectionPage methods to
// access and set this property.
func (this NextProperty) IsCollectionPage() bool {
	return this.CollectionPageMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this NextProperty) IsIRI() bool {
	return this.iri != nil
}

// IsLink returns true if this property has a type of "Link". When true, use the
// GetLink and SetLink methods to access and set this property.
func (this NextProperty) IsLink() bool {
	return this.LinkMember != nil
}

// IsMention returns true if this property has a type of "Mention". When true, use
// the GetMention and SetMention methods to access and set this property.
func (this NextProperty) IsMention() bool {
	return this.MentionMember != nil
}

// IsOrderedCollectionPage returns true if this property has a type of
// "OrderedCollectionPage". When true, use the GetOrderedCollectionPage and
// SetOrderedCollectionPage methods to access and set this property.
func (this NextProperty) IsOrderedCollectionPage() bool {
	return this.OrderedCollectionPageMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this NextProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsCollectionPage() {
		child = this.GetCollectionPage().JSONLDContext()
	} else if this.IsLink() {
		child = this.GetLink().JSONLDContext()
	} else if this.IsMention() {
		child = this.GetMention().JSONLDContext()
	} else if this.IsOrderedCollectionPage() {
		child = this.GetOrderedCollectionPage().JSONLDContext()
	}
	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this NextProperty) KindIndex() int {
	if this.IsCollectionPage() {
		return 0
	}
	if this.IsLink() {
		return 1
	}
	if this.IsMention() {
		return 2
	}
	if this.IsOrderedCollectionPage() {
		return 3
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this NextProperty) LessThan(o vocab.NextPropertyInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsCollectionPage() {
		return this.GetCollectionPage().LessThan(o.GetCollectionPage())
	} else if this.IsLink() {
		return this.GetLink().LessThan(o.GetLink())
	} else if this.IsMention() {
		return this.GetMention().LessThan(o.GetMention())
	} else if this.IsOrderedCollectionPage() {
		return this.GetOrderedCollectionPage().LessThan(o.GetOrderedCollectionPage())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "next".
func (this NextProperty) Name() string {
	return "next"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this NextProperty) Serialize() (interface{}, error) {
	if this.IsCollectionPage() {
		return this.GetCollectionPage().Serialize()
	} else if this.IsLink() {
		return this.GetLink().Serialize()
	} else if this.IsMention() {
		return this.GetMention().Serialize()
	} else if this.IsOrderedCollectionPage() {
		return this.GetOrderedCollectionPage().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return string(this.unknown), nil
}

// SetCollectionPage sets the value of this property. Calling IsCollectionPage
// afterwards returns true.
func (this *NextProperty) SetCollectionPage(v vocab.CollectionPageInterface) {
	this.Clear()
	this.CollectionPageMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *NextProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetLink sets the value of this property. Calling IsLink afterwards returns true.
func (this *NextProperty) SetLink(v vocab.LinkInterface) {
	this.Clear()
	this.LinkMember = v
}

// SetMention sets the value of this property. Calling IsMention afterwards
// returns true.
func (this *NextProperty) SetMention(v vocab.MentionInterface) {
	this.Clear()
	this.MentionMember = v
}

// SetOrderedCollectionPage sets the value of this property. Calling
// IsOrderedCollectionPage afterwards returns true.
func (this *NextProperty) SetOrderedCollectionPage(v vocab.OrderedCollectionPageInterface) {
	this.Clear()
	this.OrderedCollectionPageMember = v
}