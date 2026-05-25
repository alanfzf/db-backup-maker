package main

import "fmt"

func AppRecovery(){
    if r := recover(); r != nil {
        SendToNtfy(fmt.Sprintf("%v", r))
    }
}
