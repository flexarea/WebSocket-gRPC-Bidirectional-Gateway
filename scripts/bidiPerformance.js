import {Client, Stream} from 'k6/net/grpc';
import { check } from 'k6';
import { Trend } from 'k6/metrics';

const client = new Client();
client.load([], '../chat/proto/message.proto');

const msgLatency = new Trend('msg_latency', true);

export let options = {
  vus: 50,
  duration: '10s',
};

export default function () {
  client.connect('localhost:50051', { plaintext: true });

  const stream = new Stream(client, 'message.Message/HandleMessage');

  stream.on('data', (msg) => {
    const latency = Date.now() - msg.timestamp;
    msgLatency.add(latency);
    check(latency, { 'latency < 50ms': (l) => l < 50 });

    // ping-pong — send next only after receiving reply
    stream.write({
      content: 'hello',
      src_user_id: 1,
      dest_user_id: 2,
      timestamp: Date.now(),
    });
  });

  // kick off the ping-pong
  stream.write({
    content: 'hello',
    src_user_id: 1,
    dest_user_id: 2,
    timestamp: Date.now(),
  });

  // hold stream open for 10s
  // then end it
  client.close();
}
