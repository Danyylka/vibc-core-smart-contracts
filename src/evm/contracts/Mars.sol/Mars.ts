/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  EventFragment,
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedLogDescription,
  TypedListener,
  TypedContractMethod,
} from "../common";

export type IbcEndpointStruct = { portId: string; channelId: BytesLike };

export type IbcEndpointStructOutput = [portId: string, channelId: string] & {
  portId: string;
  channelId: string;
};

export type HeightStruct = {
  revision_number: BigNumberish;
  revision_height: BigNumberish;
};

export type HeightStructOutput = [
  revision_number: bigint,
  revision_height: bigint
] & { revision_number: bigint; revision_height: bigint };

export type IbcPacketStruct = {
  src: IbcEndpointStruct;
  dest: IbcEndpointStruct;
  sequence: BigNumberish;
  data: BytesLike;
  timeoutHeight: HeightStruct;
  timeoutTimestamp: BigNumberish;
};

export type IbcPacketStructOutput = [
  src: IbcEndpointStructOutput,
  dest: IbcEndpointStructOutput,
  sequence: bigint,
  data: string,
  timeoutHeight: HeightStructOutput,
  timeoutTimestamp: bigint
] & {
  src: IbcEndpointStructOutput;
  dest: IbcEndpointStructOutput;
  sequence: bigint;
  data: string;
  timeoutHeight: HeightStructOutput;
  timeoutTimestamp: bigint;
};

export type AckPacketStruct = { success: boolean; data: BytesLike };

export type AckPacketStructOutput = [success: boolean, data: string] & {
  success: boolean;
  data: string;
};

export interface MarsInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "ackPackets"
      | "connectedChannels"
      | "dispatcher"
      | "greet"
      | "greetWithFee"
      | "onAcknowledgementPacket"
      | "onChanCloseConfirm"
      | "onChanCloseInit"
      | "onChanOpenAck"
      | "onChanOpenConfirm"
      | "onChanOpenTry"
      | "onRecvPacket"
      | "onTimeoutPacket"
      | "owner"
      | "recvedPackets"
      | "renounceOwnership"
      | "supportedVersions"
      | "timeoutPackets"
      | "transferOwnership"
      | "triggerChannelClose"
      | "triggerChannelInit"
      | "triggerChannelInitWithFee"
  ): FunctionFragment;

  getEvent(nameOrSignatureOrTopic: "OwnershipTransferred"): EventFragment;

  encodeFunctionData(
    functionFragment: "ackPackets",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "connectedChannels",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "dispatcher",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "greet",
    values: [string, BytesLike, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "greetWithFee",
    values: [
      string,
      BytesLike,
      BigNumberish,
      [BigNumberish, BigNumberish],
      [BigNumberish, BigNumberish]
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onAcknowledgementPacket",
    values: [IbcPacketStruct, AckPacketStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "onChanCloseConfirm",
    values: [BytesLike, string, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "onChanCloseInit",
    values: [BytesLike, string, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "onChanOpenAck",
    values: [BytesLike, BytesLike, string]
  ): string;
  encodeFunctionData(
    functionFragment: "onChanOpenConfirm",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "onChanOpenTry",
    values: [BigNumberish, string[], BytesLike, string, BytesLike, string]
  ): string;
  encodeFunctionData(
    functionFragment: "onRecvPacket",
    values: [IbcPacketStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "onTimeoutPacket",
    values: [IbcPacketStruct]
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "recvedPackets",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "renounceOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "supportedVersions",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "timeoutPackets",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "triggerChannelClose",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "triggerChannelInit",
    values: [string, BigNumberish, boolean, string[], string]
  ): string;
  encodeFunctionData(
    functionFragment: "triggerChannelInitWithFee",
    values: [string, BigNumberish, boolean, string[], string]
  ): string;

  decodeFunctionResult(functionFragment: "ackPackets", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "connectedChannels",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "dispatcher", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "greet", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "greetWithFee",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onAcknowledgementPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onChanCloseConfirm",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onChanCloseInit",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onChanOpenAck",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onChanOpenConfirm",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onChanOpenTry",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onRecvPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onTimeoutPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "recvedPackets",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "supportedVersions",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "timeoutPackets",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "triggerChannelClose",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "triggerChannelInit",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "triggerChannelInitWithFee",
    data: BytesLike
  ): Result;
}

export namespace OwnershipTransferredEvent {
  export type InputTuple = [previousOwner: AddressLike, newOwner: AddressLike];
  export type OutputTuple = [previousOwner: string, newOwner: string];
  export interface OutputObject {
    previousOwner: string;
    newOwner: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface Mars extends BaseContract {
  connect(runner?: ContractRunner | null): Mars;
  waitForDeployment(): Promise<this>;

  interface: MarsInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  ackPackets: TypedContractMethod<
    [arg0: BigNumberish],
    [[boolean, string] & { success: boolean; data: string }],
    "view"
  >;

  connectedChannels: TypedContractMethod<
    [arg0: BigNumberish],
    [string],
    "view"
  >;

  dispatcher: TypedContractMethod<[], [string], "view">;

  greet: TypedContractMethod<
    [message: string, channelId: BytesLike, timeoutTimestamp: BigNumberish],
    [bigint],
    "nonpayable"
  >;

  greetWithFee: TypedContractMethod<
    [
      message: string,
      channelId: BytesLike,
      timeoutTimestamp: BigNumberish,
      gasLimits: [BigNumberish, BigNumberish],
      gasPrices: [BigNumberish, BigNumberish]
    ],
    [bigint],
    "payable"
  >;

  onAcknowledgementPacket: TypedContractMethod<
    [arg0: IbcPacketStruct, ack: AckPacketStruct],
    [void],
    "nonpayable"
  >;

  onChanCloseConfirm: TypedContractMethod<
    [channelId: BytesLike, arg1: string, arg2: BytesLike],
    [void],
    "nonpayable"
  >;

  onChanCloseInit: TypedContractMethod<
    [channelId: BytesLike, arg1: string, arg2: BytesLike],
    [void],
    "nonpayable"
  >;

  onChanOpenAck: TypedContractMethod<
    [channelId: BytesLike, arg1: BytesLike, counterpartyVersion: string],
    [void],
    "nonpayable"
  >;

  onChanOpenConfirm: TypedContractMethod<
    [channelId: BytesLike],
    [void],
    "nonpayable"
  >;

  onChanOpenTry: TypedContractMethod<
    [
      arg0: BigNumberish,
      arg1: string[],
      channelId: BytesLike,
      arg3: string,
      arg4: BytesLike,
      counterpartyVersion: string
    ],
    [string],
    "nonpayable"
  >;

  onRecvPacket: TypedContractMethod<
    [packet: IbcPacketStruct],
    [AckPacketStructOutput],
    "nonpayable"
  >;

  onTimeoutPacket: TypedContractMethod<
    [packet: IbcPacketStruct],
    [void],
    "nonpayable"
  >;

  owner: TypedContractMethod<[], [string], "view">;

  recvedPackets: TypedContractMethod<
    [arg0: BigNumberish],
    [
      [
        IbcEndpointStructOutput,
        IbcEndpointStructOutput,
        bigint,
        string,
        HeightStructOutput,
        bigint
      ] & {
        src: IbcEndpointStructOutput;
        dest: IbcEndpointStructOutput;
        sequence: bigint;
        data: string;
        timeoutHeight: HeightStructOutput;
        timeoutTimestamp: bigint;
      }
    ],
    "view"
  >;

  renounceOwnership: TypedContractMethod<[], [void], "nonpayable">;

  supportedVersions: TypedContractMethod<
    [arg0: BigNumberish],
    [string],
    "view"
  >;

  timeoutPackets: TypedContractMethod<
    [arg0: BigNumberish],
    [
      [
        IbcEndpointStructOutput,
        IbcEndpointStructOutput,
        bigint,
        string,
        HeightStructOutput,
        bigint
      ] & {
        src: IbcEndpointStructOutput;
        dest: IbcEndpointStructOutput;
        sequence: bigint;
        data: string;
        timeoutHeight: HeightStructOutput;
        timeoutTimestamp: bigint;
      }
    ],
    "view"
  >;

  transferOwnership: TypedContractMethod<
    [newOwner: AddressLike],
    [void],
    "nonpayable"
  >;

  triggerChannelClose: TypedContractMethod<
    [channelId: BytesLike],
    [void],
    "nonpayable"
  >;

  triggerChannelInit: TypedContractMethod<
    [
      version: string,
      ordering: BigNumberish,
      feeEnabled: boolean,
      connectionHops: string[],
      counterpartyPortId: string
    ],
    [void],
    "nonpayable"
  >;

  triggerChannelInitWithFee: TypedContractMethod<
    [
      version: string,
      ordering: BigNumberish,
      feeEnabled: boolean,
      connectionHops: string[],
      counterpartyPortId: string
    ],
    [void],
    "payable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "ackPackets"
  ): TypedContractMethod<
    [arg0: BigNumberish],
    [[boolean, string] & { success: boolean; data: string }],
    "view"
  >;
  getFunction(
    nameOrSignature: "connectedChannels"
  ): TypedContractMethod<[arg0: BigNumberish], [string], "view">;
  getFunction(
    nameOrSignature: "dispatcher"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "greet"
  ): TypedContractMethod<
    [message: string, channelId: BytesLike, timeoutTimestamp: BigNumberish],
    [bigint],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "greetWithFee"
  ): TypedContractMethod<
    [
      message: string,
      channelId: BytesLike,
      timeoutTimestamp: BigNumberish,
      gasLimits: [BigNumberish, BigNumberish],
      gasPrices: [BigNumberish, BigNumberish]
    ],
    [bigint],
    "payable"
  >;
  getFunction(
    nameOrSignature: "onAcknowledgementPacket"
  ): TypedContractMethod<
    [arg0: IbcPacketStruct, ack: AckPacketStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onChanCloseConfirm"
  ): TypedContractMethod<
    [channelId: BytesLike, arg1: string, arg2: BytesLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onChanCloseInit"
  ): TypedContractMethod<
    [channelId: BytesLike, arg1: string, arg2: BytesLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onChanOpenAck"
  ): TypedContractMethod<
    [channelId: BytesLike, arg1: BytesLike, counterpartyVersion: string],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onChanOpenConfirm"
  ): TypedContractMethod<[channelId: BytesLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "onChanOpenTry"
  ): TypedContractMethod<
    [
      arg0: BigNumberish,
      arg1: string[],
      channelId: BytesLike,
      arg3: string,
      arg4: BytesLike,
      counterpartyVersion: string
    ],
    [string],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onRecvPacket"
  ): TypedContractMethod<
    [packet: IbcPacketStruct],
    [AckPacketStructOutput],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onTimeoutPacket"
  ): TypedContractMethod<[packet: IbcPacketStruct], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "owner"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "recvedPackets"
  ): TypedContractMethod<
    [arg0: BigNumberish],
    [
      [
        IbcEndpointStructOutput,
        IbcEndpointStructOutput,
        bigint,
        string,
        HeightStructOutput,
        bigint
      ] & {
        src: IbcEndpointStructOutput;
        dest: IbcEndpointStructOutput;
        sequence: bigint;
        data: string;
        timeoutHeight: HeightStructOutput;
        timeoutTimestamp: bigint;
      }
    ],
    "view"
  >;
  getFunction(
    nameOrSignature: "renounceOwnership"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "supportedVersions"
  ): TypedContractMethod<[arg0: BigNumberish], [string], "view">;
  getFunction(
    nameOrSignature: "timeoutPackets"
  ): TypedContractMethod<
    [arg0: BigNumberish],
    [
      [
        IbcEndpointStructOutput,
        IbcEndpointStructOutput,
        bigint,
        string,
        HeightStructOutput,
        bigint
      ] & {
        src: IbcEndpointStructOutput;
        dest: IbcEndpointStructOutput;
        sequence: bigint;
        data: string;
        timeoutHeight: HeightStructOutput;
        timeoutTimestamp: bigint;
      }
    ],
    "view"
  >;
  getFunction(
    nameOrSignature: "transferOwnership"
  ): TypedContractMethod<[newOwner: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "triggerChannelClose"
  ): TypedContractMethod<[channelId: BytesLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "triggerChannelInit"
  ): TypedContractMethod<
    [
      version: string,
      ordering: BigNumberish,
      feeEnabled: boolean,
      connectionHops: string[],
      counterpartyPortId: string
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "triggerChannelInitWithFee"
  ): TypedContractMethod<
    [
      version: string,
      ordering: BigNumberish,
      feeEnabled: boolean,
      connectionHops: string[],
      counterpartyPortId: string
    ],
    [void],
    "payable"
  >;

  getEvent(
    key: "OwnershipTransferred"
  ): TypedContractEvent<
    OwnershipTransferredEvent.InputTuple,
    OwnershipTransferredEvent.OutputTuple,
    OwnershipTransferredEvent.OutputObject
  >;

  filters: {
    "OwnershipTransferred(address,address)": TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;
    OwnershipTransferred: TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;
  };
}
