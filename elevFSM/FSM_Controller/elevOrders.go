package FSM_Controller

import (
	"fmt"
)

/*
	The matrix(10x2) for the orders of a elevator are on the form

			2	
			
	[FLOOR][BUTTON_TYPE]
	[FLOOR][BUTTON_TYPE]
	[FLOOR][BUTTON_TYPE]
	.						}10
	.
	.
	[FLOOR][BUTTON_TYPE]

	Which is a priority list starting at the top. COMMAND button types have higher priorities than
	other button types and are automatically moved infront.

	There is four floors: 0, 1, 2, and 3.
	There is three button types: up, down and command.
*/

const (
	ROWS = 10
)

const (
	b_UP = 0
	b_DOWN = 1
	b_COMMAND = 2
)

type Button struct{
	Button_type int
	Floor int
}

var orders[ROWS] Button

func Next_order() Button{
	return orders[0]
}

func Add_order(button Button){
	order_exists := Check_if_order_exists(button)
	if order_exists==0{
		for i := 0; i < ROWS; i++{
			if orders[i].Floor == -1{
				orders[i] = button
				/*
				if orders[i].Button_type == b_COMMAND{
					move_order_infront(i)
				}
				*/
				return
			}
		}
	}
	Print_all_orders()
}

func Check_if_order_exists(button Button) int{
	exists := 0
	for i := 0; i < ROWS; i++{
		if orders[i].Floor == button.Floor && orders[i].Button_type == button.Button_type{
			exists = 1
		}
	}
	return exists
}

func Remove_order(current_floor int){
	for i := 0; i < ROWS; i++{
		if orders[i].Floor == current_floor{
			orders[i].Floor = -1
			orders[i].Button_type = -1
			left_shift_orders(i)
		}
	}
}

func left_shift_orders(index int){
	for i := index; i < ROWS-1; i++{
		orders[i].Floor = orders[i+1].Floor
		orders[i].Button_type = orders[i+1].Button_type
	}
	orders[ROWS-1].Floor = -1
	orders[ROWS-1].Button_type = -1
}

func right_shift_orders(index int){
	for i := index; i > 0; i--{
		orders[i].Floor = orders[i-1].Floor
		orders[i].Button_type = orders[i-1].Button_type
	}
}

/*
func move_order_infront(index int){
	temp_floor := orders[index].Floor
	temp_button_type := orders[index].Button_type
	orders[index].Floor = -1
	orders[index].Button_type = -1
	right_shift_orders(index)
	orders[0].Floor = temp_floor
	orders[0].Button_type = temp_button_type
}
*/	

func Order_handler(Button_Press_Chan chan Button){
	for{
	Add_order(<-Button_Press_Chan)
	}
}

func Print_all_orders(){
	for i := 0; i < ROWS;i++{
		fmt.Printf("%d,", orders[i])
	}
	fmt.Printf("\n\n")
}

func Orders_init() {
	for i := 0; i < ROWS;i++{
		orders[i].Floor = -1
		orders[i].Button_type = -1
	}
}
