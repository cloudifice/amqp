package amqp

import ()


type opsInstance struct {
  queueName string      //name of the queue
  qi *queueInstance
  channel *amqp.Channel
}


type queueInstance struct {}

func (q *opsInstance) Publish() {}
func (q *opsInstance) Subscribe() {}
func (q *opsInstance) Unsubscribe() {}

///will return the queue instance
func (qi *queueInstance) Get(q string) (PubSub, error) {
  return &opsInstance{queueName: q, qi: qi}, nil
}

func ()
