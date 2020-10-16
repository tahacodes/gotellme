package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Simple website lookup tool"
	app.Usage = "IP, CNAME, NameServers, MX records"
	hostname := string(os.Args[2])
	appFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: hostname,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "all",
			Usage: "Looks up the Name Servers, IP addresses, CNAME and mx for a particular host",
			Flags : appFlags,
			Action : func(c *cli.Context) error {
				 ns, err := net.LookupNS(c.String("host"))
                                 fmt.Println("ns :")
				 if err != nil {
					 fmt.Println(err)
                                 } else {
                                     for i :=0; i < len(ns); i++ {
                                         fmt.Println(ns[i].Host)
                                     }
				 }
                                 ip, err := net.LookupIP(c.String("host"))
				 fmt.Println("ip :")
                                 if err != nil {
					 fmt.Println(err)
                                 } else {
                                     for i :=0; i < len(ip); i++ {
                                         fmt.Println(ip[i])
                                     }
			         }
                                 cname, err := net.LookupCNAME(c.String("host"))
                                 fmt.Println("cname :")
				 if err != nil {
					 fmt.Println(err)
				 } else {
					 fmt.Println(cname)
			         }
                                 mx, err := net.LookupMX(c.String("host"))
                                 fmt.Println("mx :")
				 if err != nil {
					 fmt.Println(err)
                                 } else {
                                     for i :=0; i < len(mx); i++ {
                                         fmt.Println(mx[i].Host, mx[i].Pref)
                                     }
				 }
                                 return nil
			},
		},
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a particular host",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
