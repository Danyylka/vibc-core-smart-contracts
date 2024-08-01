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
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedListener,
  TypedContractMethod,
} from "./common";

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

export interface IUniversalChannelHandlerInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "MW_ID"
      | "closeChannel"
      | "connectedChannels"
      | "dispatcher"
      | "onAcknowledgementPacket"
      | "onChanCloseConfirm"
      | "onChanOpenAck"
      | "onChanOpenConfirm"
      | "onChanOpenTry"
      | "onRecvPacket"
      | "onTimeoutPacket"
      | "openChannel"
      | "registerMwStack"
      | "sendUniversalPacket"
      | "sendUniversalPacketWithFee"
      | "setDispatcher"
  ): FunctionFragment;

  encodeFunctionData(functionFragment: "MW_ID", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "closeChannel",
    values: [BytesLike]
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
    functionFragment: "onAcknowledgementPacket",
    values: [IbcPacketStruct, AckPacketStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "onChanCloseConfirm",
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
  encodeFunctionData(
    functionFragment: "openChannel",
    values: [string, BigNumberish, boolean, string[], string]
  ): string;
  encodeFunctionData(
    functionFragment: "registerMwStack",
    values: [BigNumberish, AddressLike[]]
  ): string;
  encodeFunctionData(
    functionFragment: "sendUniversalPacket",
    values: [BytesLike, BytesLike, BytesLike, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "sendUniversalPacketWithFee",
    values: [
      BytesLike,
      BytesLike,
      BytesLike,
      BigNumberish,
      [BigNumberish, BigNumberish],
      [BigNumberish, BigNumberish]
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "setDispatcher",
    values: [AddressLike]
  ): string;

  decodeFunctionResult(functionFragment: "MW_ID", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "closeChannel",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "connectedChannels",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "dispatcher", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "onAcknowledgementPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onChanCloseConfirm",
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
  decodeFunctionResult(
    functionFragment: "openChannel",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "registerMwStack",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "sendUniversalPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "sendUniversalPacketWithFee",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setDispatcher",
    data: BytesLike
  ): Result;
}

export interface IUniversalChannelHandler extends BaseContract {
  connect(runner?: ContractRunner | null): IUniversalChannelHandler;
  waitForDeployment(): Promise<this>;

  interface: IUniversalChannelHandlerInterface;

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

  MW_ID: TypedContractMethod<[], [bigint], "view">;

  closeChannel: TypedContractMethod<
    [channelId: BytesLike],
    [void],
    "nonpayable"
  >;

  connectedChannels: TypedContractMethod<
    [arg0: BigNumberish],
    [string],
    "view"
  >;

  dispatcher: TypedContractMethod<[], [string], "nonpayable">;

  onAcknowledgementPacket: TypedContractMethod<
    [packet: IbcPacketStruct, ack: AckPacketStruct],
    [void],
    "nonpayable"
  >;

  onChanCloseConfirm: TypedContractMethod<
    [
      channelId: BytesLike,
      counterpartyPortId: string,
      counterpartyChannelId: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  onChanOpenAck: TypedContractMethod<
    [
      channelId: BytesLike,
      counterpartychannelId: BytesLike,
      counterpartyVersion: string
    ],
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
      order: BigNumberish,
      connectionHops: string[],
      channelId: BytesLike,
      counterpartyPortIdentifier: string,
      counterpartychannelId: BytesLike,
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

  openChannel: TypedContractMethod<
    [
      version: string,
      ordering: BigNumberish,
      feeEnabled: boolean,
      connectionHops: string[],
      counterpartyPortIdentifier: string
    ],
    [void],
    "nonpayable"
  >;

  registerMwStack: TypedContractMethod<
    [mwBitmap: BigNumberish, mwAddrs: AddressLike[]],
    [void],
    "nonpayable"
  >;

  sendUniversalPacket: TypedContractMethod<
    [
      channelId: BytesLike,
      destPortAddr: BytesLike,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish
    ],
    [bigint],
    "nonpayable"
  >;

  sendUniversalPacketWithFee: TypedContractMethod<
    [
      channelId: BytesLike,
      destPortAddr: BytesLike,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish,
      gasLimits: [BigNumberish, BigNumberish],
      gasPrices: [BigNumberish, BigNumberish]
    ],
    [bigint],
    "payable"
  >;

  setDispatcher: TypedContractMethod<
    [dispatcher: AddressLike],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "MW_ID"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "closeChannel"
  ): TypedContractMethod<[channelId: BytesLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "connectedChannels"
  ): TypedContractMethod<[arg0: BigNumberish], [string], "view">;
  getFunction(
    nameOrSignature: "dispatcher"
  ): TypedContractMethod<[], [string], "nonpayable">;
  getFunction(
    nameOrSignature: "onAcknowledgementPacket"
  ): TypedContractMethod<
    [packet: IbcPacketStruct, ack: AckPacketStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onChanCloseConfirm"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      counterpartyPortId: string,
      counterpartyChannelId: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onChanOpenAck"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      counterpartychannelId: BytesLike,
      counterpartyVersion: string
    ],
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
      order: BigNumberish,
      connectionHops: string[],
      channelId: BytesLike,
      counterpartyPortIdentifier: string,
      counterpartychannelId: BytesLike,
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
    nameOrSignature: "openChannel"
  ): TypedContractMethod<
    [
      version: string,
      ordering: BigNumberish,
      feeEnabled: boolean,
      connectionHops: string[],
      counterpartyPortIdentifier: string
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "registerMwStack"
  ): TypedContractMethod<
    [mwBitmap: BigNumberish, mwAddrs: AddressLike[]],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "sendUniversalPacket"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      destPortAddr: BytesLike,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish
    ],
    [bigint],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "sendUniversalPacketWithFee"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      destPortAddr: BytesLike,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish,
      gasLimits: [BigNumberish, BigNumberish],
      gasPrices: [BigNumberish, BigNumberish]
    ],
    [bigint],
    "payable"
  >;
  getFunction(
    nameOrSignature: "setDispatcher"
  ): TypedContractMethod<[dispatcher: AddressLike], [void], "nonpayable">;

  filters: {};
}
