package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const token string = "ODA5MTQ2MzEyMjI2MjQyNjEw.YCQ2Og.mrccTpx1w5M1ua5WQucyajmZ8To"

var channelID string = "615631186345328650"

var BotID string

//Plan
var Monday []string = []string{"0", "Polski", "Matma(zgon)", "Angielski", "Wychowawcza", "Historia", "WOS", "WF", "WF"}
var Tuesday []string = []string{"Chemia", "Polski", "Polski", "Angielski", "WOS", "Matma(zgon)", "Fizyka(zgon)", "0", "0"}
var Wednesday []string = []string{"Hiszpanski", "Hiszpanski", "Biologia", "Matma(zgon)", "Geografia", "EDB", "WF", "WF", "0"}
var Thursday []string = []string{"0", "Religia", "Informatyka", "Polski", "Angielski", "Matma(zgon)", "Fizyka(zgon)", "0", "0"}
var Friday []string = []string{"Chemia", "Polski", "Polski", "Matma(zgon)", "Angielski", "Historia", "Religia", "0", "0"}
var Saturday []string = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
var Sunday []string = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"}

//Days of week in Polish
var DaysOfWeek []string = []string{"Poniedziałek", "Wtorek", "Środa", "Czwartek", "Piątek", "Sobota", "Niedziela"}
var TimeOfLessons []string = []string{"8:00-8:30", "8:50-9:20", "9:45-10:15", "10:20-10:50", "11:20-11:50", "12:20-12:50", "12:55-13:25", "13:30-14:00", "14:10-14:40"}

var dg, err = discordgo.New("Bot " + token)

func main() {
	fmt.Println(time.Now().Date())
	fmt.Println("Starting Bot...")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Bot has been started")
	}

	fmt.Println("Setting User...")
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("User has been set")
	}
	BotID = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()
	fmt.Println("Adding Handler...")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Handler has been added...")
	}

	fmt.Println("Running status.go...")
	go status()
	fmt.Println("Running writin.go...")
	go writin()
	fmt.Println("Running handler...")
	<-make(chan struct{})

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Print(time.Now().Format("01-02-2006 15:04:05") + " - " + m.ChannelID + " - " + m.Author.Username + ": " + m.Content + "\n")
	if m.Author.ID == BotID {
		return
		//komendy
	} else if m.Content == "!czas" {

		buff := ""

		for i := 0; i < 9; i++ {
			buff += strconv.FormatInt(int64(i+1), 10) + ". " + TimeOfLessons[i] + "\n"
		}

		SendEmbed(s, m.ChannelID, "", "Czasy:", buff)
	} else if m.Content == "!plan" {

		Today := Monday

		i := int(time.Now().Weekday())
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

		j := 0
		buff := ""
		for j < 9 {
			buff1 := strconv.FormatInt(int64(j+1), 10)

			buff += "**" + buff1 + ". " + Today[j] + "**" + "\n"
			buff += TimeOfLessons[j] + "\n"
			j++
		}
		buff1 := "Dzisiaj jest " + DaysOfWeek[i-1]

		SendEmbed(s, m.ChannelID, "", buff1, buff)

	} else if m.Content == "!jutro" {

		Today := Monday

		i := (int(time.Now().Weekday()) + 1) % 7
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

		j := 0
		buff := ""
		for j < 9 {
			buff1 := strconv.FormatInt(int64(j+1), 10)

			buff += "**" + buff1 + ". " + Today[j] + "**" + "\n"
			buff += TimeOfLessons[j] + "\n"
			j++
		}
		buff1 := "Jutro jest " + DaysOfWeek[i-1]
		fmt.Println(buff)
		SendEmbed(s, m.ChannelID, "", buff1, buff)
		//kejsy
	} else if strings.Contains(m.Content, "fizyka") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Dzięki Oskar")
	} else if strings.Contains(m.Content, "kto pytał") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Twój tata pytał")
	} else if strings.Contains(m.Content, " pszczół") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pszczoły dobre zwierzęta")
	} else if strings.Contains(m.Content, "Fryzjer Przecina Kłódke") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Nie wolno tego wypowiadac")
	}
}
