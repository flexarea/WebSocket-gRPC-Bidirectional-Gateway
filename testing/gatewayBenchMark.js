import ws from 'k6/ws';
import { check } from 'k6';
import { Trend } from 'k6/metrics';

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
      const now = Date.now()*1e6

      const latency = (now - msg.timestamp)/1e6;

      const grpcMs = msg.grpcTime/1e6

      msgLatency.add(latency);
      grpcLatency.add(grpcMs);

      check(latency, { 'latency < 50ms': (l) => l < 50 });

      // send next immediately after receiving reply
      socket.send(JSON.stringify({
        content: 'hello',
        src_user_id: 1,
        dest_user_id: 2,
        timestamp: Date.now()*1e6,
        grpcTime: Date.now(),

      }));
    });

    socket.on('open', () => {

      // ping-pong
      socket.send(JSON.stringify({
        content: 'hello',
        src_user_id: 1,
        dest_user_id: 2,
        timestamp: Date.now()*1e6,
        grpcTime: Date.now(),
      }));
    });

    // close connection timeout
    socket.setTimeout(() => {
      socket.close();
    }, 10000);

  });
}
