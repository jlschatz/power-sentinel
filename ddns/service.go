package ddns

import (
	"bytes"
	"log"
	"net/http"
	"os/exec"
)

type Service interface {
	GetPublicIP()
	UpdateNoIP()
}

type service struct {
	CurrentPublicIP string
}

func NewDDNSService() Service {
	s := &service{}
	s.GetPublicIP()
	return s
}

func (s *service) GetPublicIP() {

	cmd := exec.Command("dig", "+short", "myip.opendns.com", "@resolver1.opendns.com")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Println(err)
	}

	if s.CurrentPublicIP != out.String() {
		log.Println("Updating current public IP to: " + out.String())
		s.CurrentPublicIP = (out.String())
		s.UpdateNoIP()

	}
}
func (s *service) UpdateNoIP() {
	_, err := http.Get("http://jurgen.schatz@gmail.com:Gr33nfus3@dynupdate.no-ip.com/nic/update?hostname=schatzcorp.ddns.net&myip=" + s.CurrentPublicIP)
	if err != nil {
		log.Println(err)
	}
	log.Println("No-IP updated to: ", s.CurrentPublicIP)
}
