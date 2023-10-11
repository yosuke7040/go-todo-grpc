// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file todo/v1/todo.proto (package todo.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum todo.v1.Status
 */
export enum Status {
  /**
   * @generated from enum value: STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: STATUS_DOING = 1;
   */
  DOING = 1,

  /**
   * @generated from enum value: STATUS_DONE = 2;
   */
  DONE = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Status)
proto3.util.setEnumType(Status, "todo.v1.Status", [
  { no: 0, name: "STATUS_UNSPECIFIED" },
  { no: 1, name: "STATUS_DOING" },
  { no: 2, name: "STATUS_DONE" },
]);

/**
 * NOTE:
 * https://maku.blog/p/sp9q7o5/
 * ＞ データを wire 形式にシリアライズするときに、デフォルト値を保存しない
 * ＞ データの受信側では、フィールドに対応する値が入っていなければ、デフォルト値がセットされているものとして振る舞う
 * つまり、Statusの初期値は0だから、クライアントには表示されないので0でセットされていると判断する
 *
 * @generated from message todo.v1.Todo
 */
export class Todo extends Message<Todo> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string title = 2;
   */
  title = "";

  /**
   * @generated from field: todo.v1.Status status = 3;
   */
  status = Status.UNSPECIFIED;

  constructor(data?: PartialMessage<Todo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.Todo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(Status) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Todo {
    return new Todo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Todo {
    return new Todo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Todo {
    return new Todo().fromJsonString(jsonString, options);
  }

  static equals(a: Todo | PlainMessage<Todo> | undefined, b: Todo | PlainMessage<Todo> | undefined): boolean {
    return proto3.util.equals(Todo, a, b);
  }
}

/**
 * @generated from message todo.v1.CreateRequest
 */
export class CreateRequest extends Message<CreateRequest> {
  /**
   * @generated from field: string title = 1;
   */
  title = "";

  constructor(data?: PartialMessage<CreateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.CreateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRequest {
    return new CreateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRequest {
    return new CreateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRequest {
    return new CreateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRequest | PlainMessage<CreateRequest> | undefined, b: CreateRequest | PlainMessage<CreateRequest> | undefined): boolean {
    return proto3.util.equals(CreateRequest, a, b);
  }
}

/**
 * @generated from message todo.v1.CreateResponse
 */
export class CreateResponse extends Message<CreateResponse> {
  /**
   * @generated from field: todo.v1.Todo todo = 1;
   */
  todo?: Todo;

  constructor(data?: PartialMessage<CreateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.CreateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "todo", kind: "message", T: Todo },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateResponse {
    return new CreateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateResponse {
    return new CreateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateResponse {
    return new CreateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateResponse | PlainMessage<CreateResponse> | undefined, b: CreateResponse | PlainMessage<CreateResponse> | undefined): boolean {
    return proto3.util.equals(CreateResponse, a, b);
  }
}

/**
 * @generated from message todo.v1.ReadRequest
 */
export class ReadRequest extends Message<ReadRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<ReadRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.ReadRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReadRequest {
    return new ReadRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReadRequest {
    return new ReadRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReadRequest {
    return new ReadRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ReadRequest | PlainMessage<ReadRequest> | undefined, b: ReadRequest | PlainMessage<ReadRequest> | undefined): boolean {
    return proto3.util.equals(ReadRequest, a, b);
  }
}

/**
 * @generated from message todo.v1.ReadResponse
 */
export class ReadResponse extends Message<ReadResponse> {
  /**
   * @generated from field: todo.v1.Todo todo = 1;
   */
  todo?: Todo;

  constructor(data?: PartialMessage<ReadResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.ReadResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "todo", kind: "message", T: Todo },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReadResponse {
    return new ReadResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReadResponse {
    return new ReadResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReadResponse {
    return new ReadResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ReadResponse | PlainMessage<ReadResponse> | undefined, b: ReadResponse | PlainMessage<ReadResponse> | undefined): boolean {
    return proto3.util.equals(ReadResponse, a, b);
  }
}

/**
 * @generated from message todo.v1.UpdateRequest
 */
export class UpdateRequest extends Message<UpdateRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string title = 2;
   */
  title = "";

  /**
   * @generated from field: todo.v1.Status status = 3;
   */
  status = Status.UNSPECIFIED;

  constructor(data?: PartialMessage<UpdateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.UpdateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(Status) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRequest {
    return new UpdateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRequest {
    return new UpdateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRequest {
    return new UpdateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateRequest | PlainMessage<UpdateRequest> | undefined, b: UpdateRequest | PlainMessage<UpdateRequest> | undefined): boolean {
    return proto3.util.equals(UpdateRequest, a, b);
  }
}

/**
 * @generated from message todo.v1.UpdateResponse
 */
export class UpdateResponse extends Message<UpdateResponse> {
  /**
   * @generated from field: todo.v1.Todo todo = 1;
   */
  todo?: Todo;

  constructor(data?: PartialMessage<UpdateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.UpdateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "todo", kind: "message", T: Todo },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateResponse {
    return new UpdateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateResponse {
    return new UpdateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateResponse {
    return new UpdateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateResponse | PlainMessage<UpdateResponse> | undefined, b: UpdateResponse | PlainMessage<UpdateResponse> | undefined): boolean {
    return proto3.util.equals(UpdateResponse, a, b);
  }
}

/**
 * @generated from message todo.v1.DeleteRequest
 */
export class DeleteRequest extends Message<DeleteRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<DeleteRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.DeleteRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRequest {
    return new DeleteRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRequest {
    return new DeleteRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRequest {
    return new DeleteRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteRequest | PlainMessage<DeleteRequest> | undefined, b: DeleteRequest | PlainMessage<DeleteRequest> | undefined): boolean {
    return proto3.util.equals(DeleteRequest, a, b);
  }
}

/**
 * @generated from message todo.v1.DeleteResponse
 */
export class DeleteResponse extends Message<DeleteResponse> {
  constructor(data?: PartialMessage<DeleteResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.DeleteResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteResponse {
    return new DeleteResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteResponse {
    return new DeleteResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteResponse {
    return new DeleteResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteResponse | PlainMessage<DeleteResponse> | undefined, b: DeleteResponse | PlainMessage<DeleteResponse> | undefined): boolean {
    return proto3.util.equals(DeleteResponse, a, b);
  }
}

/**
 * @generated from message todo.v1.ListRequest
 */
export class ListRequest extends Message<ListRequest> {
  constructor(data?: PartialMessage<ListRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.ListRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRequest {
    return new ListRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRequest {
    return new ListRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRequest {
    return new ListRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListRequest | PlainMessage<ListRequest> | undefined, b: ListRequest | PlainMessage<ListRequest> | undefined): boolean {
    return proto3.util.equals(ListRequest, a, b);
  }
}

/**
 * @generated from message todo.v1.ListResponse
 */
export class ListResponse extends Message<ListResponse> {
  /**
   * @generated from field: repeated todo.v1.Todo todos = 1;
   */
  todos: Todo[] = [];

  constructor(data?: PartialMessage<ListResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "todo.v1.ListResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "todos", kind: "message", T: Todo, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListResponse {
    return new ListResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListResponse {
    return new ListResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListResponse {
    return new ListResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListResponse | PlainMessage<ListResponse> | undefined, b: ListResponse | PlainMessage<ListResponse> | undefined): boolean {
    return proto3.util.equals(ListResponse, a, b);
  }
}
