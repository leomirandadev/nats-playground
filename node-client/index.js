const NATS = require("nats")
const nc = NATS.connect({ json: true })


const generateClientName = () => {
  const num = new Date().getTime()
  return "node-client-" + num
}

// Simple Publisher
nc.publish("cannel-test", {
  client: generateClientName(),
  message: "The Mom is on"
})

// Simple Subscriber
nc.subscribe("cannel-test", function (msg) {
  console.log("Received a message: ", msg)
})