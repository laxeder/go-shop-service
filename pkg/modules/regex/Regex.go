package regex

import "regexp"

var HasNumber = regexp.MustCompile(`\d+`)
var HasLetter = regexp.MustCompile(`[A-Za-zÁáÀàÂâÃãÄäÉéÈèÊêËëÍíÌìÏïÓóÔôÕõÖöÚúÙùÛûçÇñÑ]`)
var HasCharSpecials = regexp.MustCompile(`[~!@#$%^&*\(\)=:;"'<>,.?\[\]\/\\\{\}|]`)
var HasCharSpecialsToPhone = regexp.MustCompile(`[~!@#$%^&*=:;"',.?\[\]\/\\]`)
var HasCharSpecialsToCnpj = regexp.MustCompile(`[~!@#$%^&*=:;"',?\[\]\\]`)
var HasUpppercase = regexp.MustCompile(`[A-ZÁÀÂÃÄÉÈÊËÍÌÏÓÔÕÖÚÙÛÇÑ]`)
var HasLowercase = regexp.MustCompile(`[a-záàâãäéèêëíìïóôõöúùûçñ]`)
var IsEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var IsBase64 = regexp.MustCompile(`^data:image\/(?:gif|png|jpeg|bmp|webp|svg\+xml)(?:;charset=utf-8)?;base64,(?:[A-Za-z0-9]|[+/])+={0,2}`)
var MimeType = regexp.MustCompile(`image[/]+[a-z]+`)
var OnlyNumber = regexp.MustCompile(`[^0-9]`)
