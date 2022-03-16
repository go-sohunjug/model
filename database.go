package model

type ActionSymbol struct {
	Exchange string `json:"exchange" bson:"exchange"`
	Model    string `json:"model" bson:"model"`
	Symbol   string `json:"symbol" bson:"symbol"`
}

type SymbolAction struct {
	Name    string `json:"name" bson:"name"`
	IP       string `json:"ip" bson:"ip"`
	Port     uint32 `json:"port" bson:"port"`
	Symbols []ActionSymbol `json:"symbols" bson:"symbols"`
}
