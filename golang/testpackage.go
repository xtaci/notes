package main

import (
	"fmt"
	"misc/packet"
)

type user_login_info struct {
	F_login_way          int32
	F_open_udid          string
	F_client_certificate string
	F_client_version     int32
	F_user_lang          string
	F_app_id             string
	F_os_verson          string
	F_device_name        string
	F_device_id          string
	F_device_id_type     int32
}

func (p user_login_info) Pack(w *packet.Packet) {
	w.WriteS32(p.F_login_way)
	w.WriteString(p.F_open_udid)
	w.WriteString(p.F_client_certificate)
	w.WriteS32(p.F_client_version)
	w.WriteString(p.F_user_lang)
	w.WriteString(p.F_app_id)
	w.WriteString(p.F_os_verson)
	w.WriteString(p.F_device_name)
	w.WriteString(p.F_device_id)
	w.WriteS32(p.F_device_id_type)

}

func main() {
	u := user_login_info{}
	fmt.Println(packet.Pack(1, u, nil))
	fmt.Println(packet.Pack(2, &u, nil))
}
