package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
//var bookings = make([]map[string]string,0)

var bookings = make([]UserDataStruct,0)

type UserDataStruct struct {
	firstName string
	lastName string
	email string 
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	
	greetUsers()

	//	var bookings [50]string array declartion
	// var bookings = []string{}
	//bookings:=[]string{}
	// bookings[0] = "Nana"
	// bookings[1] = "rodrig"		
	//bookings[0]=firstName+" "+lastName
	
	for {
		firstName,lastName,email,userTickets := getUserInput()

		isValidEmail,isValidName,isValidTicketNumber:=helper.ValidateUserInput(firstName,lastName,email,userTickets,remainingTickets)

		if  isValidEmail && isValidName&&isValidTicketNumber{

			//book tickets
			bookTicket(userTickets,firstName ,lastName ,email)

			wg.Add(1)
			go sendTicket(userTickets,firstName,lastName ,email)

			//print first names
			fmt.Printf("The first names of bookings are: %v\n", FirstName())

			//var noTicketsRemaining bool = remainingTickets == 0
			//noTicketsRemaining :=remainingTickets==0
			
			if (remainingTickets == 0) {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		}else{

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			  if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			  }
			  if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
wg.Wait()
	}

	// city:="london"
	// switch city{
	// case "new york":
	// 	break
	// case "New York":
	// 	break
	// case "Singapore":
	// 	break
	// case "London":
	// 	break
	// case "Berlin","Mexico City":   // some code here
	// 	break
	// case "Hong Kong":
	// 	break
	// default:
	// 	fmt.Printf("No valid city selected")
	// }

}

func greetUsers(){
	fmt.Printf("Welcome to our %v booking application ",conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, "are still remaining")
	fmt.Println("Get your tickets here to attend")
//	fmt.Printf("conferenceTickets is %T, remainingTickets is %T . conferenceName is %T\n", confTickets, remainingTickets, confName)
}

func FirstName() []string{
	firstNames:= [] string {}
	for _, booking:= range bookings {
	//		var names = strings.Fields(booking) names[0]
// booking["firstName"] for maps

	firstNames = append(firstNames,booking.firstName)
	}
  return firstNames
}

func getUserInput()(string,string,string,uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	fmt.Scan( & firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan( & lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan( & email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan( & userTickets)
	return firstName,lastName,email,userTickets
}

func bookTicket(userTickets uint,firstName string,lastName string,email string){
	// userName="tom"
	// userTickets=2
	// fmt.Printf("The whole array %v\n",bookings)
	// fmt.Printf("The first value array %v\n",bookings[9])
	// fmt.Printf("The array type %T\n",bookings)
	// fmt.Printf("Array length %v\n",len(bookings))

	// fmt.Printf("The whole slice %v\n", bookings)
	// fmt.Printf("The first value slice %v\n", bookings[9])
	// fmt.Printf("The slice type %T\n", bookings)
	// fmt.Printf("slice length %v\n", len(bookings))

	//define slice userSLice := []string{}

	//create map for user
	//	var userData = make(map[string]string)
// userData["firstName"]=firstName
// userData["lastName"]=lastName
// userData["email"]=email
// userData["numberOfTickets"]=strconv.FormatUint(uint64(userTickets),10)

//create userData struct for users
var userData = UserDataStruct{
	firstName : firstName,
	lastName :lastName,
	email :email,
	numberOfTickets: userTickets,
}
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, userData)

	fmt.Printf("List Of Bookings %v\n",bookings)

	fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v number of tickets remaining for %v", remainingTickets,conferenceName)
}

func sendTicket(userTickets uint,firstName string,lastName string,email string){

	time.Sleep(50*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName)	
	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n  %v to email address %v", ticket,email)
	fmt.Println("###########")
	wg.Done()
}