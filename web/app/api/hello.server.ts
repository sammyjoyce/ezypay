import { createChannel, createClient } from 'nice-grpc';
import { HelloServiceDefinition, HelloServiceClient } from '@generated/protos/hello/service';

// Client Setup
let _client: HelloServiceClient;

function getClient(): HelloServiceClient {
  if (!_client) {
    const channel = createChannel('localhost:50051' );
    _client = createClient(HelloServiceDefinition, channel);
  }
  return _client;
}

export async function sayHello(name: string) {
  try {
    const response = await getClient().sayHello({ name });
    console.log('gRPC response:', response);
    return response;
  } catch (error) {
    console.error('gRPC call failed:', error);
    throw error;
  }
}
