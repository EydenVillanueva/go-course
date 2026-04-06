package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value interface{}
	IsSet bool
}

/*
%v - The default formatting (Minimalista)
%+v - Adds the fields to the values
%#v - Prints the value as you would write it in code using go example: main.Mesa{Numero:5, Capacidad:4, Ocupada:true}
%s - string
%d - integer
%f - float
% .2f - float with 2 decimals
%T - prints the type
%t - boolean
%q - double quote strings
%% - for putting the % character
*/
func (c ConfigItem) String() string {
	return fmt.Sprintf("key: %s, Value: %s, IsSet: %t", c.Key, c.Value, c.IsSet)
}

func main() {

}
