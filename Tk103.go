package main

import (
	"fmt"
)

func main() {
	TestPacket := "(027043145986AP05150630V5207.4772N00505.9965E000.5115026138.2610000000L000000ffT324)"
	fmt.Println(TestPacket)
	Head := TestPacket[0:1]
	fmt.Println("Head = " + Head)
	SerialNo := TestPacket[1:13]
	fmt.Println("SerialNo = " + SerialNo)
	Command := TestPacket[13:17]
	fmt.Println("Command = " + Command)
	MessageBody := TestPacket[17 : len(TestPacket)-1]
	fmt.Println("MessageBody = " + MessageBody)
	Trail := TestPacket[len(TestPacket)-1 : len(TestPacket)]
	fmt.Println("Trail = " + Trail)
	CommandType(Command)
}
func CommandType(inputcommand string) {
	Commandchar1 := inputcommand[0:1]
	Commandchar2 := inputcommand[1:2]
	Commandserial := inputcommand[2:4]

	if Commandchar1 == "A" {
		fmt.Println("Down Message")
		if Commandchar2 == "P" {
			fmt.Println("Device Parameter Message")
			switch Commandserial {
			case "00":
				fmt.Println("One time Calling Message")
			case "03":
				fmt.Println("Read device parameter configuring message")
			case "04":
				fmt.Println("Read device operated status message")
			case "05":
				fmt.Println("Device login response message")
			case "07":
				fmt.Println("Center No. configuring message")
			case "11":
				fmt.Println("Cell phone NO. configuring message")
			case "12":
				fmt.Println("Setting vehicle high and low limit speed")
			case "15":
				fmt.Println("Monitor Command")
			case "17":
				fmt.Println("Read device cell phone configuring")
			default:
				fmt.Println("Error in Device Parameter Message")
			}
		} else if Commandchar2 == "Q" {
			fmt.Println("General Communication Message")
			if Commandserial == "00" {
				fmt.Println("Common Message")
			} else if Commandserial == "01" {
				fmt.Println("Attemper Message")
			} else if Commandserial == "02" {
				fmt.Println("Answer of calling message(Taxi)")
			} else if Commandserial == "03" {
				fmt.Println("Calling Message (Taxi)")
			} else if Commandserial == "04" {
				fmt.Println("Navigation Message")
			} else {
				fmt.Println("Error in General Communication Message")
			}
		} else if Commandchar2 == "R" {
			fmt.Println("Vehicle positioning Message Answer message")
			if Commandserial == "00" {
				fmt.Println("Isochronous for continues feedback configuring")
			} else if Commandserial == "01" {
				fmt.Println("Isometry for continues feedback configuring")
			} else if Commandserial == "05" {
				fmt.Println("Set ACC open sending data transmiting intervals")
			} else if Commandserial == "06" {
				fmt.Println("Set ACC open sending data transmiting intervals")
			} else {
				fmt.Println("Error in Vehicle positioning Message Answer message")
			}
		} else if Commandchar2 == "S" {
			fmt.Println("Answer Signal")
			if Commandserial == "01" {
				fmt.Println("Answer Alarm Message")
			} else if Commandserial == "07" {
				fmt.Println("Answer Message for getting customer successfully (Taxi)")
			} else {
				fmt.Println("Error in Answer Signal")
			}
		} else if Commandchar2 == "T" {
			fmt.Println("Control Signal Restarted")
			if Commandserial == "00" {
				fmt.Println("Control the restarted message of the device")
			} else {
				fmt.Println("Error in Control Signal Restarted Message")
			}
		} else if Commandchar2 == "V" {
			fmt.Println("Control Signal")
			if Commandserial == "00" {
				fmt.Println("Circuit control signal")
			} else if Commandserial == "01" {
				fmt.Println("Oil control signal")
			} else if Commandserial == "02" {
				fmt.Println("One key configuring command")
			} else if Commandserial == "03" {
				fmt.Println("Read one key configuring")
			} else {
				fmt.Println("Error in Control Signal")
			}
		} else if Commandchar2 == "X" {
			fmt.Println("Expanding Message")
			if Commandserial == "00" {
				fmt.Println("Answer currency up explaining result message")
			} else if Commandserial == "01" {
				fmt.Println("Alarm configuring message")
			} else if Commandserial == "02" {
				fmt.Println("Device Function configuring command")
			} else if Commandserial == "03" {
				fmt.Println("Device mode configured command")
			} else if Commandserial == "04" {
				fmt.Println("Intialized device command")
			} else if Commandserial == "05" {
				fmt.Println("Setting Geo-fence Message")
			} else {
				fmt.Println("Error in Expanding Message")
			}
		} else {
			fmt.Println("Error in Command Character 2")
		}
	} else if Commandchar1 == "B" {
		fmt.Println("Up Message")
		if Commandchar2 == "O" {
			fmt.Println("Alarm Message")
			if Commandserial == "01" {
				fmt.Println("Alarm Message")
			} else if Commandserial == "02" {
				fmt.Println("Answer device parameter configured message")
			} else if Commandserial == "03" {
				fmt.Println("Answer device operated status message")
			} else if Commandserial == "04" {
				fmt.Println("Answer Calling Message")
			} else if Commandserial == "05" {
				fmt.Println("Answer device login response message")
			} else if Commandserial == "12" {
				fmt.Println("Answer vehicle high and low speed limit")
			} else if Commandserial == "07" {
				fmt.Println("Message for getting customer successfully (Taxi)")
			} else {
				fmt.Println("Error in Alarm Message")
			}
		} else if Commandchar2 == "R" {
			fmt.Println("Vehicle Positioning Message")
			if Commandserial == "00" {
				fmt.Println("Isochronous for continues feedback message")
			} else if Commandserial == "01" {
				fmt.Println("Isometry continous feedback message")
			} else if Commandserial == "02" {
				fmt.Println("Continues feedback ending messsage")
			} else if Commandserial == "05" {
				fmt.Println("Answer the Setting ACC open sending data transmiting intervals")
			} else if Commandserial == "06" {
				fmt.Println("Answer the Setting  ACC open sending data transmiting intervals")
			} else {
				fmt.Println("Error in Vehicle Positioning Message")
			}
		} else if Commandchar2 == "S" {
			fmt.Println("Answer Message")
			if Commandserial == "04" {
				fmt.Println("Answer attempered Message")
			} else if Commandserial == "05" {
				fmt.Println("Answer reading called configuring number")
			} else if Commandserial == "06" {
				fmt.Println("Answer caller configuring number")
			} else if Commandserial == "08" {
				fmt.Println("Answer setting isochronous feedback message")
			} else if Commandserial == "09" {
				fmt.Println("Answer setting Isometry feedback message")
			} else if Commandserial == "20" {
				fmt.Println("Answer response calling message (Taxi)")
			} else if Commandserial == "21" {
				fmt.Println("Answer calling message(Taxi)")
			} else if Commandserial == "23" {
				fmt.Println("Answer navigation message")
			} else {
				fmt.Println("Error in Answer Message")
			}
		} else if Commandchar2 == "T" {
			if Commandserial == "00" {
				fmt.Println("Answer the restarted message of the device")
			} else {
				fmt.Println("Error in Answering the restarted message of the device")
			}
		} else if Commandchar2 == "U" {
			if Commandserial == "00" {
				fmt.Println("Answer the Setting Geo-fence Message")
			} else {
				fmt.Println("Error in Answering the Setting Geo-fence Message")
			}
		} else if Commandchar2 == "V" {
			fmt.Println("Answer Control Sign")
			if Commandserial == "00" {
				fmt.Println("Answer circuit control")
			} else if Commandserial == "01" {
				fmt.Println("Answer oil control")
			} else if Commandserial == "02" {
				fmt.Println("Answer enquiring of one key setting")
			} else {
				fmt.Println("Error in Answer Control Sign")
			}
		} else {
			fmt.Println("Error in Up Message")
		}
	} else {
		fmt.Println("Incorrect Command Type")
	}
}
