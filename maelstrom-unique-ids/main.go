package main

import (
    "encoding/json"
    "log"
    "github.com/google/uuid"
//    "fmt"
    maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
    // Create a new Maelstrom node
    n := maelstrom.NewNode()

    // Register handler for "echo" message
    n.Handle("generate", func(msg maelstrom.Message) error {
        // Unmarshal message body
        var body map[string]any
	
	id := uuid.New()
	//fmt.Println(id.String())

        if err := json.Unmarshal(msg.Body, &body); err != nil {
            return err
        }

        // Update message type to "echo_ok"
        body["type"] = "generate_ok"
	body["id"] = id.String()

        // Reply with the updated message
        return n.Reply(msg, body)
    })

    // Run the Maelstrom node
    if err := n.Run(); err != nil {
        log.Fatal(err)
    }
}
