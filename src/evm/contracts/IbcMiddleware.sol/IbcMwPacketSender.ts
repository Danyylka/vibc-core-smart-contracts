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
} from "../common";

export interface IbcMwPacketSenderInterface extends Interface {
  getFunction(nameOrSignature: "sendMWPacket"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "sendMWPacket",
    values: [
      BytesLike,
      BytesLike,
      BytesLike,
      BigNumberish,
      BytesLike,
      BigNumberish
    ]
  ): string;

  decodeFunctionResult(
    functionFragment: "sendMWPacket",
    data: BytesLike
  ): Result;
}

export interface IbcMwPacketSender extends BaseContract {
  connect(runner?: ContractRunner | null): IbcMwPacketSender;
  waitForDeployment(): Promise<this>;

  interface: IbcMwPacketSenderInterface;

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

  sendMWPacket: TypedContractMethod<
    [
      channelId: BytesLike,
      srcPortAddr: BytesLike,
      destPortAddr: BytesLike,
      srcMwIds: BigNumberish,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish
    ],
    [bigint],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "sendMWPacket"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      srcPortAddr: BytesLike,
      destPortAddr: BytesLike,
      srcMwIds: BigNumberish,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish
    ],
    [bigint],
    "nonpayable"
  >;

  filters: {};
}