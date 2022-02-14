package validations



func CheckIfParameterEmptyOrNil(parameter string) bool{
	if parameter != "000000000000000000000000"{
		return true
	}else if parameter != ""{
		return true
	}else{
		return false
	}
}