package ddns

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
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
	go func() {
		for now := range time.Tick(time.Minute) {
			log.Println("Confirming Public IP validity at: " + now.Format("03:04:05 02/01/2006"))
			s.GetPublicIP()
		}
	}()
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

	Ip := strings.TrimSuffix(out.String(), "\n")

	if s.CurrentPublicIP != Ip {
		log.Println("Updating current public IP to: " + Ip)
		s.CurrentPublicIP = Ip
		s.UpdateNoIP()

	}
}
func (s *service) UpdateNoIP() {

	_, err := http.Get(fmt.Sprintf("http://jurgen.schatz@gmail.com:Gr33nfus3@dynupdate.no-ip.com/nic/update?hostname=schatzcorp.ddns.net&myip=%v", s.CurrentPublicIP))
	if err != nil {
		log.Println(err)
	}
	log.Println("No-IP updated to: ", s.CurrentPublicIP)
}
