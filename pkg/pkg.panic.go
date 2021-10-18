package pkg

import "log"

func PanicIfError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
