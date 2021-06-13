package main

import (
	"fmt"
	"strconv"
	"time"
)

var Timers [19]time.Time
var i int
var Lessons []string = []string{"8h", "8h30m", "8h50m", "9h20m", "9h45m", "10h15m", "10h20m", "10h50m", "11h20m", "11h50m", "12h20m", "12h50m", "12h55m", "13h25m", "13h30m", "14h", "14h10m", "14h40m"}

func status() {

	j := time.Now()
	j = j.Add(time.Hour * time.Duration(time.Now().Hour()) * -1)
	j = j.Add(time.Minute * time.Duration(time.Now().Minute()) * -1)
	j = j.Add(time.Second * time.Duration(time.Now().Second()) * -1)
	fmt.Println("Parsing time...")
	for jd1 := 0; jd1 < 18; jd1++ {
		Timers[jd1] = j
		jd, _ := time.ParseDuration(Lessons[jd1])
		Timers[jd1] = Timers[jd1].Add(jd)
	}
	fmt.Println("Time has been parsed")
	var bul bool = true
	var bul2 bool = true
	Day := int(time.Now().Weekday())

	Today := Monday

	i = Day
	switch i {
	case 0:
		Today = Sunday
	case 1:
		Today = Monday
	case 2:
		Today = Tuesday
	case 3:
		Today = Wednesday
	case 4:
		Today = Thursday
	case 5:
		Today = Friday
	case 6:
		Today = Saturday
	case 7:
		Today = Sunday
	}

	i = 0
	slip := true
	fmt.Println("Status running...")
	for {

		if i >= 18 {
			i = 0

			for jd := 0; jd < 18; jd++ {
				Timers[jd] = Timers[jd].Add(time.Hour * 24)
			}

			Day = int(time.Now().Weekday()) + 1
			switch Day {
			case 0:
				Today = Sunday
			case 1:
				Today = Monday
			case 2:
				Today = Tuesday
			case 3:
				Today = Wednesday
			case 4:
				Today = Thursday
			case 5:
				Today = Friday
			case 6:
				Today = Saturday
			case 7:
				Today = Sunday
			}
		}

		if slip {
			time.Sleep(time.Second * 5)
		}
		do := time.Until(Timers[i])
		if do < 0 {
			i++
			slip = false
			bul = true
			bul2 = true
		} else {
			slip = true
			do = do.Round(time.Second * 5)
			if Today[i/2] == "0" {
				i++
				continue
			}
			if i%2 == 0 {
				if i < 18 && int(time.Now().Weekday()) != 6 {
					dg.UpdateListeningStatus(do.String() + " do " + Today[i/2])
				} else {
					dg.UpdateListeningStatus(do.String() + " do " + Today[8])
				}
			} else {
				dg.UpdateListeningStatus(do.String() + " do Przerwy")
			}

			if do.String() == "1m0s" {
				if i%2 == 0 && bul {
					AtMinute(Today[i/2])
					bul = false
				}

			}
			if i > 0 {
				if Timers[i-1].Round(time.Minute).Add(time.Minute) == time.Now().Round(time.Minute) && i%2 == 0 && bul2 {
					buff, _ := time.ParseDuration(Lessons[i-1])
					buff1 := Timers[i].Add(buff * -1)
					buff2 := strconv.Itoa(Timers[i].Hour()) + ":" + strconv.Itoa(Timers[i].Minute())
					Break(Today[i/2], strconv.Itoa(buff1.Minute()), buff2)
					bul2 = false

				}
			}

		}

	}
}

func AtMinute(lesson string) {

	buff := "Za minutę " + lesson

	if lesson != "0" {
		_, _ = dg.ChannelMessageSendTTS(channelID, buff)
	}
}

func Break(lesson string, LenghtOfBreak string, TimeOfNext string) {
	buff := LenghtOfBreak + " minut przerwy, potem o " + TimeOfNext + " " + lesson
	if lesson != "0" {
		_, _ = dg.ChannelMessageSendTTS(channelID, buff)
	}
}
