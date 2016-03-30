package elevDrivers // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.h and driver.go
/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
#include "elev.h"
*/
import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"elevTypes"
)

func elev_init() int {
	return int(C.elev_set_motor_direction(C.elev_motor_direction_t(dirn)))
}

/*void elev_init(void) {
    int init_success = io_init();
    assert(init_success && "Unable to initialize elevator hardware!");

    for (int f = 0; f < N_FLOORS; f++) {
        for (elev_button_type_t b = 0; b < N_BUTTONS; b++){
            elev_set_button_lamp(b, f, 0);
        }
    }

    elev_set_stop_lamp(0);
    elev_set_door_open_lamp(0);
    elev_set_floor_indicator(0);
}
*/

func elev_set_motor_direction(elev_motor_direction_t dirn) int {
	return int(C.elev_set_motor_direction(C.elev_motor_direction_t(dirn)))
}

/*void elev_set_motor_direction(elev_motor_direction_t dirn) {
    if (dirn == 0){
        io_write_analog(MOTOR, 0);
    } else if (dirn > 0) {
        io_clear_bit(MOTORDIR);
        io_write_analog(MOTOR, MOTOR_SPEED);
    } else if (dirn < 0) {
        io_set_bit(MOTORDIR);
        io_write_analog(MOTOR, MOTOR_SPEED);
    }
}
*/

func elev_set_button_lamp(elev_button_type_t button, int floor, int value) int {
	return int(C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value)))
}

/*void elev_set_button_lamp(elev_button_type_t button, int floor, int value) {
    assert(floor >= 0);
    assert(floor < N_FLOORS);
    assert(button >= 0);
    assert(button < N_BUTTONS);

    if (value) {
        io_set_bit(lamp_channel_matrix[floor][button]);
    } else {
        io_clear_bit(lamp_channel_matrix[floor][button]);
    }
}
*/

func elev_set_floor_indicator(int floor) int {
	return int(C.elev_set_floor_indicator(C.int(floor)))
}

/*void elev_set_floor_indicator(int floor) {
    assert(floor >= 0);
    assert(floor < N_FLOORS);

    // Binary encoding. One light must always be on.
    if (floor & 0x02) {
        io_set_bit(LIGHT_FLOOR_IND1);
    } else {
        io_clear_bit(LIGHT_FLOOR_IND1);
    }

    if (floor & 0x01) {
        io_set_bit(LIGHT_FLOOR_IND2);
    } else {
        io_clear_bit(LIGHT_FLOOR_IND2);
    }
}
*/

func elev_set_door_open_lamp(int value) int {
	return int(C.elev_set_door_open_lamp(C.int(value)))
}

/*void elev_set_door_open_lamp(int value) {
    if (value) {
        io_set_bit(LIGHT_DOOR_OPEN);
    } else {
        io_clear_bit(LIGHT_DOOR_OPEN);
    }
}
*/

func elev_set_stop_lamp(int value) int {
	return int(C.elev_set_stop_lamp(C.int(value)))
}

/*void elev_set_stop_lamp(int value) {
    if (value) {
        io_set_bit(LIGHT_STOP);
    } else {
        io_clear_bit(LIGHT_STOP);
    }
}
*/

func elev_get_button_signal(elev_button_type_t button, int floor) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

/*int elev_get_button_signal(elev_button_type_t button, int floor) {
    assert(floor >= 0);
    assert(floor < N_FLOORS);
    assert(button >= 0);
    assert(button < N_BUTTONS);


    if (io_read_bit(button_channel_matrix[floor][button])) {
        return 1;
    } else {
        return 0;
    }
}
*/

func elev_get_floor_sensor_signal() int {
	return int(C.elev_get_floor_sensor_signal())
}

/*
int elev_get_floor_sensor_signal(void) {
    if (io_read_bit(SENSOR_FLOOR1)) {
        return 0;
    } else if (io_read_bit(SENSOR_FLOOR2)) {
        return 1;
    } else if (io_read_bit(SENSOR_FLOOR3)) {
        return 2;
    } else if (io_read_bit(SENSOR_FLOOR4)) {
        return 3;
    } else {
        return -1;
    }
}
*/

func elev_get_stop_singal() int {
	return int(C.elev_get_stop_signal())
}

//int elev_get_stop_signal(void) {
//    return io_read_bit(STOP);
//}

func elev_get_obstruction_signal(void) int {
	return int(C.elev_get_obstruction_signal())
}

//int elev_get_obstruction_signal(void) {
//    return io_read_bit(OBSTRUCTION);
//}
