package cycular_queue

type queue struct {
	 length int
	 front int
	 rear int
	 elemets []interface{}
}



func create_queue(size int ) *queue{
	cq := queue{length : size +1 , front: 0,rear: 0}
	cq.elemets = make([]interface{},cq.length)
	return  &cq
}

func (q queue) is_empty() bool {
	return q.rear==q.front
}

func (q queue) is_full() bool  {
	return q.front ==((q.rear+1)&q.length)
}

func (q *queue) push(e interface{})  {
	if q.is_full(){
		panic("sorry . queue is full")
	}
	q.elemets[q.rear]=e
	q.rear=q.rear+1
}

func (q *queue) del() (e interface{} ,ans bool){
	if q.is_empty() {
	return nil, false
	}
	e=q.elemets[q.front]
	q.elemets[q.front]=nil
	q.front=(q.front+1)%(q.length)
	q.length=q.length-1
	return e,true
}

func (q queue) shift() (e interface{}) {
	if q.is_empty() {
		return nil
	}
	e=q.elemets[q.front]
	q.front=(q.front+1)%(q.length)
	return e
}