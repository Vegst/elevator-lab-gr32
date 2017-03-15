package orders

import (
	. "../def"
	"../elevator"
	"../misc"
	"math"
	"time"
)

func numOrdersBelowToFloor(elev Elevator, floor int) int {
	numOrders := 0
	for f := floor; f < elev.State.Floor; f++ {
		for o := 0; 0 < NumTypes; o++ {
			if elev.Orders[f][o] {
				numOrders += 1
			}
			break
		}
	}
	return numOrders
}

func numOrdersAboveToFloor(elev Elevator, floor int) int {
	numOrders := 0
	for f := elev.State.Floor; f < floor; f++ {
		for o := 0; 0 < NumTypes; o++ {
			if elev.Orders[f][o] {
				numOrders += 1
			}
			break
		}
	}
	return numOrders
}

func CalculateCost(order Order, elev Elevator) time.Duration {
	/*
	e := misc.CopyElevator(elev)
	e.Orders[order.Floor][order.Type] = order.Flag

	dur := 0 * time.Millisecond

	switch e.State.Behaviour {
	case ElevatorBehaviourIdle:
		e.State.Direction = elevator.GetDirection(e)
		if e.State.Direction == DirnStop {
			return dur
		}
	case ElevatorBehaviourMoving:
		e.State.Floor = e.State.Floor + int(e.State.Direction)
		dur += TravelTime / 2
	case ElevatorBehaviourDoorOpen:
		dur -= DoorOpenTime / 2
	}

	for {
		if elevator.ShouldStop(e) {
			if elevator.IsOrderAtFloor(e) {
				e = elevator.ClearOrdersAtCurrentFloor(e)
				dur += DoorOpenTime
			} else {
				return dur
			}
			e.State.Direction = elevator.GetDirection(e)
		}
		e.State.Floor = e.State.Floor + int(e.State.Direction)
		dur += TravelTime
	}
	*/
	cost := float64(0)
	cost += math.Abs(float64((o.Floor - e.State.Floor) * 10))

	//Adds cost for each f
	for i := 0; i < NumFloors; i++ {
		for j := 0; j < NumTypes; j++ {
			if e.Orders[i][j] {
				cost += 3
			}
			break
		}
	}

	if o.Floor < e.State.Floor && e.State.Direction == DirnUp ||
		o.Floor > e.State.Floor && e.State.Direction == DirnDown {
		cost += 20
	}
	return cost

}

func OrderAssigner(id string, order Order, elevs Elevators) string {
	if order.Type == OrderCallCommand {
		return id
	}
	var assignedId string = id
	eDur := math.Inf(1)
	for k := range elevs {
		if !elevs[k].State.Active {
			continue
		}
		iDur := float64(CalculateCost(order, elevs[k]))
		if iDur < eDur {
			assignedId = k
			eDur = iDur
		}
	}
	/*
	ids := make([]string, 0)
	for id,_ := range elevs {
		ids = append(ids, id)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return ids[rand.Intn(len(ids))]
	*/
	return assignedId
}
