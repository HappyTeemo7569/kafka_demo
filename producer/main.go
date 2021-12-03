package producer

import "time"

func Put() {

	for i := 0; i < 3; i++ {
		producer := new(Producer)
		producer.InitProducer()
		go func() {
			for {
				producer.SendMessage()
				time.Sleep(1 * time.Second)
			}
		}()

	}

}
