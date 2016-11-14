package internal

import (
    "os"
    "strings"
)

const EMPTY_STRING = ""
const DEBUG_TRUE = "TRUE"
const DEBUG_FALSE = "FALSE"

func IsDebug() {
    env := os.Getenv("DEBUG")
    // No environment value Debug
    if env == "" {
        return false
    }

    env = strings.ToUpper(env)
    if env == DEBUG_TRUE {
        return true
    }
    return false

}