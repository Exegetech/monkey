package util

func IsLetter(ch byte) bool {
  isLowerCase := 'a' <= ch && ch <= 'z'
  isUpperCase := 'A' <= ch && ch <= 'Z'
  isUnderscore := ch == '_'

  return isLowerCase || isUpperCase || isUnderscore
}

func IsDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func IsWhitespace(ch byte) bool {
  isSpace := ch == ' '
  isTab := ch == '\t' 
  isNewLine := ch == '\n' 
  isCarriageReturn := ch == '\r' 

  return isSpace || isTab || isNewLine || isCarriageReturn
}

