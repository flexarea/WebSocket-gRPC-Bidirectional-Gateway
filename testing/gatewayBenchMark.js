import ws from 'k6/ws';
import { check } from 'k6';

export let options = {
  vus: 50,
  duration: '10s',
};

export default function () {
  ws.connect('ws://localhost:8080/ws', {}, (socket) => {

    socket.on('message', (data) => {
      check(data, { 'response received': (d) => d.length > 0 });
    });

    socket.on('open', () => {
      for (let i = 0; i < 10000; i++) {
        socket.send(JSON.stringify({
          content: 'hello',
          src_user_id: 1,
          dest_user_id: 2,
        }));
      }
    });

    socket.setTimeout(() => {
      socket.close();
    }, 10000);

  });
}
