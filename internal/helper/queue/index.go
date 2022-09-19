package queue

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

var (
	List = []Dequeue{}
)

type Dequeue struct {
	Type       string `json:"type"`
	IncidentId int    `json:"incidentId"`
	CodeName   string `json:"codeName"`
	OfficerId  int    `json:"officerId"`
	Loc        Loc    `json:"loc"`
}

type Loc struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type InC struct {
	ID       int    `json:"id"`
	CodeName string `json:"codeName"`
	Loc      Loc    `json:"loc"`
	OffcerID int    `json:"offcerId"`
}

type Officers struct {
	ID        int    `json:"id"`
	Loc       Loc    `json:"loc"`
	BadgeName string `json:"badgeName"`
}

type Data struct {
	In       []InC      `json:"incidents"`
	Officers []Officers `json:"officers"`
}

func Get(a chan bool) Dequeue {
	return Dequeue{}
}

func adf(byteArray []byte) {
	f, err := os.OpenFile("./data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	n, err := f.Write(byteArray)
	if err != nil {
		fmt.Println(n, err)
	}

	if n, err = f.WriteString("\n"); err != nil {
		fmt.Println(n, err)
	}
}

func MainRa() {
	sdf := "amqp://guest:guest@localhost:5672/"

	amqpServerURL := sdf

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		"events", // queue name
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no local
		false,    // no wait
		nil,      // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)
	// file, err := os.Create("data.txt")
	// f, err := os.OpenFile("text.log",
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }
	go func() {
		for message := range messages {

			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)
			var abc = Dequeue{}
			json.Unmarshal(message.Body, &abc)
			fsd := fmt.Sprintf("%#v", abc)
			fmt.Println("da:", fsd)
			List = append(List, abc)
			adf(message.Body)
		}
	}()

	<-forever

	// rc := RabbitConfig{
	// 	Schema:         "amqp",
	// 	Username:       "",
	// 	Password:       "",
	// 	Host:           "localhost",
	// 	Port:           "5672",
	// 	VHost:          "my_vhost",
	// 	ConnectionName: "events",
	// }
	// rbt := NewRabbit(rc)
	// if err := rbt.Connect(); err != nil {
	// 	log.Fatalln("unable to connect to rabbit", err)
	// }
	// //

	// // Consumer
	// cc := ConsumerConfig{
	// 	ExchangeName:  "user",
	// 	ExchangeType:  "direct",
	// 	RoutingKey:    "create",
	// 	QueueName:     "user_create",
	// 	ConsumerName:  "my_app_name",
	// 	ConsumerCount: 3,
	// 	PrefetchCount: 1,
	// }
	// cc.Reconnect.MaxAttempt = 60
	// cc.Reconnect.Interval = 1 * time.Second
	// csm := NewConsumer(cc, rbt)
	// if err := csm.Start(); err != nil {
	// 	log.Fatalln("unable to start consumer", err)
	// }
	// //

	// select {}
}

///
type RabbitConfig struct {
	Schema         string
	Username       string
	Password       string
	Host           string
	Port           string
	VHost          string
	ConnectionName string
}

type Rabbit struct {
	config     RabbitConfig
	connection *amqp.Connection
}

// NewRabbit returns a RabbitMQ instance.
func NewRabbit(config RabbitConfig) *Rabbit {
	return &Rabbit{
		config: config,
	}
}

// Connect connects to RabbitMQ server.
func (r *Rabbit) Connect() error {
	if r.connection == nil || r.connection.IsClosed() {
		con, err := amqp.DialConfig(fmt.Sprintf(
			"%s://%s:%s@%s:%s/%s",
			r.config.Schema,
			r.config.Username,
			r.config.Password,
			r.config.Host,
			r.config.Port,
			r.config.VHost,
		), amqp.Config{Properties: amqp.Table{"connection_name": r.config.ConnectionName}})
		if err != nil {
			return err
		}
		r.connection = con
	}

	return nil
}

// Connection returns exiting `*amqp.Connection` instance.
func (r *Rabbit) Connection() (*amqp.Connection, error) {
	if r.connection == nil || r.connection.IsClosed() {
		return nil, errors.New("connection is not open")
	}

	return r.connection, nil
}

// Channel returns a new `*amqp.Channel` instance.
func (r *Rabbit) Channel() (*amqp.Channel, error) {
	chn, err := r.connection.Channel()
	if err != nil {
		return nil, err
	}

	return chn, nil
}

// CONSUMER ---------------------------------------------------------------------------------------------

type ConsumerConfig struct {
	ExchangeName  string
	ExchangeType  string
	RoutingKey    string
	QueueName     string
	ConsumerName  string
	ConsumerCount int
	PrefetchCount int
	Reconnect     struct {
		MaxAttempt int
		Interval   time.Duration
	}
}

type Consumer struct {
	config ConsumerConfig
	Rabbit *Rabbit
}

// NewConsumer returns a consumer instance.
func NewConsumer(config ConsumerConfig, rabbit *Rabbit) *Consumer {
	return &Consumer{
		config: config,
		Rabbit: rabbit,
	}
}

// Start declares all the necessary components of the consumer and
// runs the consumers. This is called one at the application start up
// or when consumer needs to reconnects to the server.
func (c *Consumer) Start() error {
	con, err := c.Rabbit.Connection()
	if err != nil {
		return err
	}
	go c.closedConnectionListener(con.NotifyClose(make(chan *amqp.Error)))

	chn, err := con.Channel()
	if err != nil {
		return err
	}

	if err := chn.ExchangeDeclare(
		c.config.ExchangeName,
		c.config.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	if _, err := chn.QueueDeclare(
		c.config.QueueName,
		true,
		false,
		false,
		false,
		amqp.Table{"x-queue-mode": "lazy"},
	); err != nil {
		return err
	}

	if err := chn.QueueBind(
		c.config.QueueName,
		c.config.RoutingKey,
		c.config.ExchangeName,
		false,
		nil,
	); err != nil {
		return err
	}

	if err := chn.Qos(c.config.PrefetchCount, 0, false); err != nil {
		return err
	}

	for i := 1; i <= c.config.ConsumerCount; i++ {
		id := i
		go c.consume(chn, id)
	}

	// Simulate manual connection close
	//_ = con.Close()

	return nil
}

// closedConnectionListener attempts to reconnect to the server and
// reopens the channel for set amount of time if the connection is
// closed unexpectedly. The attempts are spaced at equal intervals.
func (c *Consumer) closedConnectionListener(closed <-chan *amqp.Error) {
	log.Println("INFO: Watching closed connection")

	// If you do not want to reconnect in the case of manual disconnection
	// via RabbitMQ UI or Server restart, handle `amqp.ConnectionForced`
	// error code.
	err := <-closed
	if err != nil {
		log.Println("INFO: Closed connection:", err.Error())

		var i int

		for i = 0; i < c.config.Reconnect.MaxAttempt; i++ {
			log.Println("INFO: Attempting to reconnect")

			if err := c.Rabbit.Connect(); err == nil {
				log.Println("INFO: Reconnected")

				if err := c.Start(); err == nil {
					break
				}
			}

			time.Sleep(c.config.Reconnect.Interval)
		}

		if i == c.config.Reconnect.MaxAttempt {
			log.Println("CRITICAL: Giving up reconnecting")

			return
		}
	} else {
		log.Println("INFO: Connection closed normally, will not reconnect")
		os.Exit(0)
	}
}

// consume creates a new consumer and starts consuming the messages.
// If this is called more than once, there will be multiple consumers
// running. All consumers operate in a round robin fashion to distribute
// message load.
func (c *Consumer) consume(channel *amqp.Channel, id int) {
	msgs, err := channel.Consume(
		c.config.QueueName,
		fmt.Sprintf("%s (%d/%d)", c.config.ConsumerName, id, c.config.ConsumerCount),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(fmt.Sprintf("CRITICAL: Unable to start consumer (%d/%d)", id, c.config.ConsumerCount))

		return
	}

	log.Println("[", id, "] Running ...")
	log.Println("[", id, "] Press CTRL+C to exit ...")

	for msg := range msgs {
		log.Println("[", id, "] Consumed:", string(msg.Body))

		if err := msg.Ack(false); err != nil {
			// TODO: Should DLX the message
			log.Println("unable to acknowledge the message, dropped", err)
		}
	}

	log.Println("[", id, "] Exiting ...")
}
