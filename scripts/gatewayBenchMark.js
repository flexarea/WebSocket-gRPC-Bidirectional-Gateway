import ws from 'k6/ws';
import { check } from 'k6';
import { Trend } from 'k6/metrics';
import custom from 'k6/x/custom-time';


const msgLatency = new Trend('msg_latency', true);
const grpcLatency = new Trend('grpc_latency', true);

export let options = {
  vus: 50,
  duration: '10s',
};

export default function () {
  ws.connect('ws://localhost:8080/ws', {}, (socket) => {

    socket.on('message', (data) => {
      const msg = JSON.parse(data);
      const now = custom.getNano()

      console.log(`now: ${now} msg: ${msg.timestamp}`)

      const period = (now - msg.timestamp)
      
      const latency = period/1e6;
      //console.log(`micro-diff: ${period} in ms-diff: ${latency}`)
      //console.log("Latency", latency )

      const grpcMs = msg.grpcTime/1e6
     // console.log("received grpcLatency", grpcMs)

      msgLatency.add(latency);
      grpcLatency.add(grpcMs);

      check(latency, { 'latency < 50ms': (l) => l < 50 });

      // send next immediately after receiving reply
      socket.send(JSON.stringify({
        content: 'hello',
        src_user_id: 1,
        dest_user_id: 2,
        timestamp: custom.getNano(),
        grpcTime: Date.now(),
      }));
    });

    socket.on('open', () => {

      // ping-pong
      socket.send(JSON.stringify({
        content: 'hello',
        src_user_id: 1,
        dest_user_id: 2,
        timestamp: custom.getNano(),
        grpcTime: Date.now(),
      }));
    });

    // close connection timeout
    socket.setTimeout(() => {
      socket.close();
    }, 10000);

  });
}
