package toledoapi

type Object[T any] struct {
	XMLName string `xml:"#ns" json:"#ns"`
	Type    string `xml:"#type" json:"#type"`
	Value   T      `xml:"#value" json:"#value"`
}

type ObjectList[T any] struct {
	Response
	Objects []Object[T] `xml:"Objects" json:"Objects"`
}

type ResponseObjects[T any] struct {
	Object[ObjectList[T]]
}
