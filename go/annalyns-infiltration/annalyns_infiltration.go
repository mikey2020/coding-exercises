package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
    canAttack := true
    if knightIsAwake {
        canAttack = false
    }
    return canAttack;
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    canSpy := false
    if knightIsAwake {
        canSpy = true
    } else if archerIsAwake {
        canSpy = true
    } else if prisonerIsAwake {
        canSpy = true
    }
    return canSpy;
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    canSignal := false
    if prisonerIsAwake && !archerIsAwake {
        canSignal = true
    }
    return canSignal
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
    canFree := false
    if prisonerIsAwake && !archerIsAwake && !knightIsAwake {
        canFree = true
    } else if petDogIsPresent && !archerIsAwake { 
        canFree = true
    }
    return canFree
}
