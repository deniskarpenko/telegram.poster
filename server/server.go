package server

type server struct{
	url string
	port int
}

func NewServer(urlNew string, portNew int) *server {
	return  &server{urlNew, portNew}
}


