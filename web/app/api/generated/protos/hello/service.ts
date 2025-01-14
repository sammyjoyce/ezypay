// @ts-nocheck
// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               v3.19.1
// source: service.proto

/* eslint-disable */
import { type CallContext, type CallOptions } from "nice-grpc-common";
import { HelloRequest, HelloResponse } from "./messages.js";

export const protobufPackage = "hello";

/** The greeting service definition */
export type HelloServiceDefinition = typeof HelloServiceDefinition;
export const HelloServiceDefinition = {
  name: "HelloService",
  fullName: "hello.HelloService",
  methods: {
    /** Sends a greeting */
    sayHello: {
      name: "SayHello",
      requestType: HelloRequest,
      requestStream: false,
      responseType: HelloResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface HelloServiceImplementation<CallContextExt = {}> {
  /** Sends a greeting */
  sayHello(request: HelloRequest, context: CallContext & CallContextExt): Promise<DeepPartial<HelloResponse>>;
}

export interface HelloServiceClient<CallOptionsExt = {}> {
  /** Sends a greeting */
  sayHello(request: DeepPartial<HelloRequest>, options?: CallOptions & CallOptionsExt): Promise<HelloResponse>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
