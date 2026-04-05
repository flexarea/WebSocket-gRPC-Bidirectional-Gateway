ghz --insecure \
    --proto ./proto/message.proto \
    --call message.Message/HandleMessage \
    --data '{"content": "hello", "src_user_id": 1, "dest_user_id": 2}' \
    --concurrency 10 \
    --duration 10s \
    localhost:50051
