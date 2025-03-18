package toledoapi

type Object[T any] struct {
	XMLName string `xml:"#ns" json:"#ns"`
	Type    string `xml:"#type" json:"#type"`
	Value   T      `xml:"#value" json:"#value"`
}

type ObjectList[T any] struct {
	Error      *string         `xml:"Error" json:"Error"`
	LastObject *Object[string] `xml:"LastObject,omitempty" json:"LastObject,omitempty"`
	Objects    []Object[T]     `xml:"Objects" json:"Objects"`
}

type ResponseObjects[T any] struct {
	Object[ObjectList[T]]
}
