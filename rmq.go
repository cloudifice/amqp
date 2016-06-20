package amqp

import (

  "sync"
  "log"
  "github.com/tsuru/config"
  "github.com/streadway/amqp"

)


type opsInstance struct {
  queueName string      //name of the queue
  qi *amqpInstance
  channel *amqp.Channel
}


type amqpInstance struct { sync.Mutex}

func (q *opsInstance) Publish(message []byte) error {

  channel, err := q.qi.dial(q.queueName) //TODO: revisit if a queueinspect is required for each call.
  if err != nil {
    return err
  }

   log.Printf("[QS] - Info - Publishing message to queue - (%s)  (%q)", q.queueName, message)

    err = channel.Publish(
     "",
		q.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
   if err != nil {
     return err
   }
   log.Printf("[QS] - Info - Published successfully")

return nil

}
//func (q *opsInstance) Subscribe() {}
//func (q *opsInstance) Unsubscribe() {}

///will return the queue instance
func (qi *amqpInstance) New(name string) (PubSub, error) {
  return &opsInstance{queueName: name, qi: qi}, nil
}

func (qi *amqpInstance) dial(queueName string) (*amqp.Channel , error){

  amqpAddr, err := config.GetString("amqp:url") //setup on cloudifice config - cfs.yml
	if err != nil {
		amqpAddr = "amqp://172.17.0.3:5672/"
	}
	conn, err := amqp.Dial(amqpAddr)
	if err != nil {
		return nil, err
	}

	log.Printf(" [QS] Dialed to (%s)", amqpAddr)

	channel, err := conn.Channel()

	if err != nil {
		return nil, err
	}

//NOTE: This is a passive call.
//TODO: Does everycall require a check?
	q, err := channel.QueueInspect(queueName)
  if err != nil {
		return nil, err
	}

	log.Printf(" [x] Connection successful to  (%s,%s)", amqpAddr, q.Name)
	return channel, err
}
