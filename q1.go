package main

import "fmt"

// queues are elements that are been processed in order or based on a priority.Elements are added to end and removed from the start.

// Queue array of people at MVA.

type Queue1 [] *MVAWaitList


type MVAWaitList struct {
	priority int 
	quantity int 
	product string
	customerName string
}


// the New method just initialezes the mvawaitlist properties.

func (mva *MVAWaitList) New(priority, quantity int, product,customerName string){
	mva.priority = priority
	mva.quantity = quantity
	mva.product = product
	mva.customerName = customerName
}


// this method takes the mvawaitlist and adds it to the queue type based on priority.

func (queue *Queue1) Add(mva *MVAWaitList){
	// check if there's no element in queue. then append element to queue.
	if len(*queue) == 0{
		*queue = append(*queue,mva)
	}else{
		// if not empty.
		var appended bool 
		appended = false

		var index int 

		var addedToWaitList *MVAWaitList

		for index, addedToWaitList = range *queue {
			if mva.priority > addedToWaitList.priority {
				*queue = append((*queue)[:index],append(Queue1{mva}, (*queue)[index:]...)...)
				appended = true
				break
			}

		}
		if !appended {
			*queue = append(*queue,mva)
		}
	}

}


func main(){
	var queue Queue1
	queue = make(Queue1, 0)

	var mva1 *MVAWaitList = &MVAWaitList{}
	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "ID Card"
	var customerName1 string = "Larry Hogan"

	mva1.New(priority1,quantity1,product1,customerName1)


	var mva2 *MVAWaitList = &MVAWaitList{}
	var priority2 int = 6
	var quantity2 int = 15
	var product2 string = "Tags"
	var customerName2 string = "Henry Ford"

	mva2.New(priority2,quantity2,product2,customerName2)

	queue.Add(mva1)
	queue.Add(mva2)

	var i int 
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}



}