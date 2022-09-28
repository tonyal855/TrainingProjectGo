package views

import (
	"sesi8-latihan/server/params"
)

type GetAllPeopleSwagger struct {
	Response
	Payload []params.Person
}
