package main

import(
    "elevDrivers"
)

func main(){
    elev_init();
    fmt.printf("Press STOP button to stop elevator and exit program.\n");

    elev_set_motor_direction(DIRN_UP);

    for ()
    {
        if (elev_get_floor_sensor_signal() == N_FLOORS -1) {
            elev_set_motor_direction(DIRN_DOWN);
        } else if (elev_get_floor_sensor_signal() == 0) {
            elev_set_motor_direction(DIRN_UP);
        }

        if (elev_get_stop_signal()) {
            elev_set_motor_direction(DIRN_STOP);
            return 0;
            
        }
    }
}



