package propertysummary

import (
	"fmt"
	langstring "github.com/go-fed/activity/streams/values/langString"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// SummaryPropertyIterator is an iterator for a property. It is permitted to be
// one of multiple value types. At most, one type of value can be present, or
// none at all. Setting a value will clear the other types of values so that
// only one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type SummaryPropertyIterator struct {
	stringMember     string
	hasStringMember  bool
	langStringMember map[string]string
	unknown          []byte
	iri              *url.URL
	alias            string
	langMap          map[string]string
	myIdx            int
	parent           vocab.SummaryPropertyInterface
}

// NewSummaryPropertyIterator creates a new summary property.
func NewSummaryPropertyIterator() *SummaryPropertyIterator {
	return &SummaryPropertyIterator{alias: ""}
}

// deserializeSummaryPropertyIterator creates an iterator from an element that has
// been unmarshalled from a text or binary format.
func deserializeSummaryPropertyIterator(i interface{}, aliasMap map[string]string) (*SummaryPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &SummaryPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if v, err := string1.DeserializeString(i); err == nil {
		this := &SummaryPropertyIterator{
			alias:           alias,
			hasStringMember: true,
			stringMember:    v,
		}
		return this, nil
	} else if v, err := langstring.DeserializeLangString(i); err == nil {
		this := &SummaryPropertyIterator{
			alias:            alias,
			langStringMember: v,
		}
		return this, nil
	} else if str, ok := i.(string); ok {
		this := &SummaryPropertyIterator{
			alias:   alias,
			unknown: []byte(str),
		}
		return this, nil
	}
	return nil, fmt.Errorf("could not deserialize %q property", "summary")
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this SummaryPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetLangString returns the value of this property. When IsLangString returns
// false, GetLangString will return an arbitrary value.
func (this SummaryPropertyIterator) GetLangString() map[string]string {
	return this.langStringMember
}

// GetLanguage returns the value for the specified BCP47 language code, or an
// empty string if it is either not a language map or no value is present.
func (this SummaryPropertyIterator) GetLanguage(bcp47 string) string {
	if this.langMap == nil {
		return ""
	} else if v, ok := this.langMap[bcp47]; ok {
		return v
	} else {
		return ""
	}
}

// GetString returns the value of this property. When IsString returns false,
// GetString will return an arbitrary value.
func (this SummaryPropertyIterator) GetString() string {
	return this.stringMember
}

// HasAny returns true if any of the values are set, except for the natural
// language map. When true, the specific has, getter, and setter methods may
// be used to determine what kind of value there is to access and set this
// property. To determine if the property was set as a natural language map,
// use the IsLanguageMap method instead.
func (this SummaryPropertyIterator) HasAny() bool {
	return this.IsString() ||
		this.IsLangString() ||
		this.iri != nil
}

// HasLanguage returns true if the natural language map has an entry for the
// specified BCP47 language code.
func (this SummaryPropertyIterator) HasLanguage(bcp47 string) bool {
	if this.langMap == nil {
		return false
	} else {
		_, ok := this.langMap[bcp47]
		return ok
	}
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this SummaryPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsLangString returns true if this property has a type of "langString". When
// true, use the GetLangString and SetLangString methods to access and set
// this property.. To determine if the property was set as a natural language
// map, use the IsLanguageMap method instead.
func (this SummaryPropertyIterator) IsLangString() bool {
	return this.langStringMember != nil
}

// IsLanguageMap determines if this property is represented by a natural language
// map. When true, use HasLanguage, GetLanguage, and SetLanguage methods to
// access and mutate the natural language map. The clear method can be used to
// clear the natural language map. Note that this method is only used for
// natural language representations, and does not determine the presence nor
// absence of other values for this property.
func (this SummaryPropertyIterator) IsLanguageMap() bool {
	return this.langMap != nil
}

// IsString returns true if this property has a type of "string". When true, use
// the GetString and SetString methods to access and set this property.. To
// determine if the property was set as a natural language map, use the
// IsLanguageMap method instead.
func (this SummaryPropertyIterator) IsString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this SummaryPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string

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
func (this SummaryPropertyIterator) KindIndex() int {
	if this.IsString() {
		return 0
	}
	if this.IsLangString() {
		return 1
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
func (this SummaryPropertyIterator) LessThan(o vocab.SummaryPropertyIteratorInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsString() {
		return string1.LessString(this.GetString(), o.GetString())
	} else if this.IsLangString() {
		return langstring.LessLangString(this.GetLangString(), o.GetLangString())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "summary".
func (this SummaryPropertyIterator) Name() string {
	return "summary"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this SummaryPropertyIterator) Next() vocab.SummaryPropertyIteratorInterface {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this SummaryPropertyIterator) Prev() vocab.SummaryPropertyIteratorInterface {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *SummaryPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetLangString sets the value of this property and clears the natural language
// map. Calling IsLangString afterwards will return true. Calling
// IsLanguageMap afterwards returns false.
func (this *SummaryPropertyIterator) SetLangString(v map[string]string) {
	this.clear()
	this.langStringMember = v
}

// SetLanguage sets the value for the specified BCP47 language code.
func (this *SummaryPropertyIterator) SetLanguage(bcp47, value string) {
	this.hasStringMember = false
	this.langStringMember = nil
	this.unknown = nil
	this.iri = nil
	if this.langMap == nil {
		this.langMap = make(map[string]string)
	}
	this.langMap[bcp47] = value
}

// SetString sets the value of this property and clears the natural language map.
// Calling IsString afterwards will return true. Calling IsLanguageMap
// afterwards returns false.
func (this *SummaryPropertyIterator) SetString(v string) {
	this.clear()
	this.stringMember = v
	this.hasStringMember = true
}

// clear ensures no value and no language map for this property is set. Calling
// HasAny or any of the 'Is' methods afterwards will return false.
func (this *SummaryPropertyIterator) clear() {
	this.hasStringMember = false
	this.langStringMember = nil
	this.unknown = nil
	this.iri = nil
	this.langMap = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this SummaryPropertyIterator) serialize() (interface{}, error) {
	if this.IsString() {
		return string1.SerializeString(this.GetString())
	} else if this.IsLangString() {
		return langstring.SerializeLangString(this.GetLangString())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return string(this.unknown), nil
}

// SummaryProperty is the non-functional property "summary". It is permitted to
// have one or more values, and of different value types.
type SummaryProperty struct {
	properties []*SummaryPropertyIterator
	alias      string
}

// DeserializeSummaryProperty creates a "summary" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeSummaryProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.SummaryPropertyInterface, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "summary"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "summary")
	}
	if i, ok := m[propName]; ok {
		this := &SummaryProperty{
			alias:      alias,
			properties: []*SummaryPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeSummaryPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeSummaryPropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewSummaryProperty creates a new summary property.
func NewSummaryProperty() *SummaryProperty {
	return &SummaryProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property "summary"
func (this *SummaryProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &SummaryPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendLangString appends a langString value to the back of a list of the
// property "summary". Invalidates iterators that are traversing using Prev.
func (this *SummaryProperty) AppendLangString(v map[string]string) {
	this.properties = append(this.properties, &SummaryPropertyIterator{
		alias:            this.alias,
		langStringMember: v,
		myIdx:            this.Len(),
		parent:           this,
	})
}

// AppendString appends a string value to the back of a list of the property
// "summary". Invalidates iterators that are traversing using Prev.
func (this *SummaryProperty) AppendString(v string) {
	this.properties = append(this.properties, &SummaryPropertyIterator{
		alias:           this.alias,
		hasStringMember: true,
		myIdx:           this.Len(),
		parent:          this,
		stringMember:    v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this SummaryProperty) At(index int) vocab.SummaryPropertyIteratorInterface {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this SummaryProperty) Begin() vocab.SummaryPropertyIteratorInterface {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this SummaryProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this SummaryProperty) End() vocab.SummaryPropertyIteratorInterface {
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this SummaryProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this SummaryProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "summary" property.
func (this SummaryProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this SummaryProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetString()
			rhs := this.properties[j].GetString()
			return string1.LessString(lhs, rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetLangString()
			rhs := this.properties[j].GetLangString()
			return langstring.LessLangString(lhs, rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this SummaryProperty) LessThan(o vocab.SummaryPropertyInterface) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property: "summary".
func (this SummaryProperty) Name() string {
	return "summary"
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "summary".
func (this *SummaryProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*SummaryPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependLangString prepends a langString value to the front of a list of the
// property "summary". Invalidates all iterators.
func (this *SummaryProperty) PrependLangString(v map[string]string) {
	this.properties = append([]*SummaryPropertyIterator{{
		alias:            this.alias,
		langStringMember: v,
		myIdx:            0,
		parent:           this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependString prepends a string value to the front of a list of the property
// "summary". Invalidates all iterators.
func (this *SummaryProperty) PrependString(v string) {
	this.properties = append([]*SummaryPropertyIterator{{
		alias:           this.alias,
		hasStringMember: true,
		myIdx:           0,
		parent:          this,
		stringMember:    v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "summary", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *SummaryProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &SummaryPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this SummaryProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// SetIRI sets an IRI value to be at the specified index for the property
// "summary". Panics if the index is out of bounds.
func (this *SummaryProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &SummaryPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetLangString sets a langString value to be at the specified index for the
// property "summary". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *SummaryProperty) SetLangString(idx int, v map[string]string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &SummaryPropertyIterator{
		alias:            this.alias,
		langStringMember: v,
		myIdx:            idx,
		parent:           this,
	}
}

// SetString sets a string value to be at the specified index for the property
// "summary". Panics if the index is out of bounds. Invalidates all iterators.
func (this *SummaryProperty) SetString(idx int, v string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &SummaryPropertyIterator{
		alias:           this.alias,
		hasStringMember: true,
		myIdx:           idx,
		parent:          this,
		stringMember:    v,
	}
}

// Swap swaps the location of values at two indices for the "summary" property.
func (this SummaryProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}