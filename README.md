# Handler Helper
Package that helps with writing http response, especially in handling error, through various processes. The features should be ONLY called by a handler or its helper. Due to its nature, all of its functions requres httpwriter as one of the parameter, and whenever error occured it writes error directly from it.

## Types Format
Utilizes types defined by package `dodol` (`https://github.com/karincake/dodol`)

## Validation
Utilizes validations defined by package `serabi` (`https://github.com/karincake/serabi`)

## Language
Utilizes dictionary features defined by package `lepet` (`https://github.com/karincake/lepet`)

## Disclaimer
This is a VERY OPINIONATED way in writing http response. Each people or company has their own standard in writing http response.